-- Query includes duplicates, will return multiple rows for the same transaction if both to and from addresses are in the address list.
--
-- The addresses list will have priority in deciding the source of the duplicate transaction; see filter_addresses temp table
-- TODO: #11980
-- However, if the addresses list is empty, and all addresses should be included, the accounts table will be used
--
-- The switch for tr_type is used to de-conflict the source for the two entries for the same transaction
--
-- UNION ALL is used to avoid the overhead of DISTINCT given that we don't expect to have duplicate entries outside the sender and receiver addresses being in the list which is handled separately
--
-- Only status FailedAS, PendingAS and CompleteAS are returned. FinalizedAS requires correlation with blockchain current state. As an optimization we approximate it by using timestamp information; see startTimestamp and endTimestamp
--
-- Token filtering has two parts
-- 1. Filtering by symbol (multi_transactions and pending_transactions tables) where the chain ID is ignored, basically the filter_networks will account for that
-- 2. Filtering by token identity (chain and address for transfers table) where the symbol is ignored and all the token identities must be provided
--
WITH filter_conditions AS (
	SELECT
		? AS startFilterDisabled,
		? AS startTimestamp,
		? AS endFilterDisabled,
		? AS endTimestamp,
		? AS filterActivityTypeAll,
		? AS filterActivityTypeSend,
		? AS filterActivityTypeReceive,
		? AS fromTrType,
		? AS toTrType,
		? AS filterAllAddresses,
		? AS filterAllToAddresses,
		? AS filterAllActivityStatus,
		? AS filterStatusCompleted,
		? AS filterStatusFailed,
		? AS filterStatusFinalized,
		? AS filterStatusPending,
		? AS statusFailed,
		? AS statusSuccess,
		? AS statusPending,
		? AS includeAllTokenTypeAssets,
		? AS includeAllNetworks,
		? AS pendingStatus
),
filter_addresses(address) AS (
	SELECT
		HEX(address)
	FROM
		%s
	WHERE
		(
			SELECT
				filterAllAddresses
			FROM
				filter_conditions
		) != 0
	UNION
	ALL
	SELECT
		*
	FROM
		(
			VALUES
				%s
		)
	WHERE
		(
			SELECT
				filterAllAddresses
			FROM
				filter_conditions
		) = 0
),
filter_to_addresses(address) AS (
	VALUES
		%s
),
assets_token_codes(token_code) AS (
	VALUES
		%s
),
assets_erc20(chain_id, token_address) AS (
	VALUES
		%s
),
filter_networks(network_id) AS (
	VALUES
		%s
),
tr_status AS (
	SELECT
		multi_transaction_id,
		MIN(status) AS min_status,
		COUNT(*) AS count,
		network_id
	FROM
		transfers
	WHERE
		transfers.loaded == 1
		AND transfers.multi_transaction_id != 0
	GROUP BY
		transfers.multi_transaction_id
),
tr_network_ids AS (
	SELECT
		multi_transaction_id
	FROM
		transfers
	WHERE
		transfers.loaded == 1
		AND transfers.multi_transaction_id != 0
		AND network_id IN filter_networks
	GROUP BY
		transfers.multi_transaction_id
),
pending_status AS (
	SELECT
		multi_transaction_id,
		COUNT(*) AS count,
		network_id
	FROM
		pending_transactions,
		filter_conditions
	WHERE
		pending_transactions.multi_transaction_id != 0
		AND pending_transactions.status = pendingStatus
	GROUP BY
		pending_transactions.multi_transaction_id
),
pending_network_ids AS (
	SELECT
		multi_transaction_id
	FROM
		pending_transactions,
		filter_conditions
	WHERE
		pending_transactions.multi_transaction_id != 0
		AND pending_transactions.status = pendingStatus
		AND pending_transactions.network_id IN filter_networks
	GROUP BY
		pending_transactions.multi_transaction_id
)
SELECT
	transfers.hash AS transfer_hash,
	NULL AS pending_hash,
	transfers.network_id AS network_id,
	0 AS multi_tx_id,
	transfers.timestamp AS timestamp,
	NULL AS mt_type,
	CASE
		WHEN from_join.address IS NOT NULL
		AND to_join.address IS NULL THEN fromTrType
		WHEN to_join.address IS NOT NULL
		AND from_join.address IS NULL THEN toTrType
		WHEN from_join.address IS NOT NULL
		AND to_join.address IS NOT NULL THEN CASE
			WHEN from_join.address < to_join.address THEN fromTrType
			ELSE toTrType
		END
		ELSE NULL
	END as tr_type,
	transfers.tx_from_address AS from_address,
	transfers.tx_to_address AS to_address,
	transfers.address AS owner_address,
	transfers.amount_padded128hex AS tr_amount,
	NULL AS mt_from_amount,
	NULL AS mt_to_amount,
	CASE
		WHEN transfers.status IS 1 THEN statusSuccess
		ELSE statusFailed
	END AS agg_status,
	1 AS agg_count,
	transfers.token_address AS token_address,
	transfers.token_id AS token_id,
	NULL AS token_code,
	NULL AS from_token_code,
	NULL AS to_token_code,
	NULL AS out_network_id,
	NULL AS in_network_id,
	transfers.type AS type,
	transfers.contract_address AS contract_address
FROM
	transfers,
	filter_conditions
	LEFT JOIN filter_addresses from_join ON HEX(transfers.tx_from_address) = from_join.address
	LEFT JOIN filter_addresses to_join ON HEX(transfers.tx_to_address) = to_join.address
WHERE
	transfers.loaded == 1
	AND transfers.multi_transaction_id = 0
	AND (
		(
			startFilterDisabled
			OR transfers.timestamp >= startTimestamp
		)
		AND (
			endFilterDisabled
			OR transfers.timestamp <= endTimestamp
		)
	)
	AND (
		filterActivityTypeAll
		OR (
			filterActivityTypeSend
			AND (
				filterAllAddresses
				OR (
					HEX(transfers.tx_from_address) IN filter_addresses
				)
			)
		)
		OR (
			filterActivityTypeReceive
			AND (
				filterAllAddresses
				OR (HEX(transfers.tx_to_address) IN filter_addresses)
			)
		)
	)
	AND (
		filterAllAddresses
		OR (
			HEX(transfers.tx_from_address) IN filter_addresses
		)
		OR (HEX(transfers.tx_to_address) IN filter_addresses)
	)
	AND (
		filterAllToAddresses
		OR (
			HEX(transfers.tx_to_address) IN filter_to_addresses
		)
	)
	AND (
		includeAllTokenTypeAssets
		OR (
			transfers.type = "eth"
			AND ("ETH" IN assets_token_codes)
		)
		OR (
			transfers.type = "erc20"
			AND (
				(
					transfers.network_id,
					HEX(transfers.token_address)
				) IN assets_erc20
			)
		)
	)
	AND (
		includeAllNetworks
		OR (transfers.network_id IN filter_networks)
	)
	AND (
		filterAllActivityStatus
		OR (
			(
				filterStatusCompleted
				OR filterStatusFinalized
			)
			AND transfers.status = 1
		)
		OR (
			filterStatusFailed
			AND transfers.status = 0
		)
	)
UNION
ALL
SELECT
	NULL AS transfer_hash,
	pending_transactions.hash AS pending_hash,
	pending_transactions.network_id AS network_id,
	0 AS multi_tx_id,
	pending_transactions.timestamp AS timestamp,
	NULL AS mt_type,
	CASE
		WHEN from_join.address IS NOT NULL
		AND to_join.address IS NULL THEN fromTrType
		WHEN to_join.address IS NOT NULL
		AND from_join.address IS NULL THEN toTrType
		WHEN from_join.address IS NOT NULL
		AND to_join.address IS NOT NULL THEN CASE
			WHEN from_join.address < to_join.address THEN fromTrType
			ELSE toTrType
		END
		ELSE NULL
	END as tr_type,
	pending_transactions.from_address AS from_address,
	pending_transactions.to_address AS to_address,
	NULL AS owner_address,
	pending_transactions.value AS tr_amount,
	NULL AS mt_from_amount,
	NULL AS mt_to_amount,
	statusPending AS agg_status,
	1 AS agg_count,
	NULL AS token_address,
	NULL AS token_id,
	pending_transactions.symbol AS token_code,
	NULL AS from_token_code,
	NULL AS to_token_code,
	NULL AS out_network_id,
	NULL AS in_network_id,
	pending_transactions.type AS type,
	NULL as contract_address
FROM
	pending_transactions,
	filter_conditions
	LEFT JOIN filter_addresses from_join ON HEX(pending_transactions.from_address) = from_join.address
	LEFT JOIN filter_addresses to_join ON HEX(pending_transactions.to_address) = to_join.address
WHERE
	pending_transactions.multi_transaction_id = 0
	AND pending_transactions.status = pendingStatus
	AND (
		filterAllActivityStatus
		OR filterStatusPending
	)
	AND (
		(
			startFilterDisabled
			OR timestamp >= startTimestamp
		)
		AND (
			endFilterDisabled
			OR timestamp <= endTimestamp
		)
	)
	AND (
		filterActivityTypeAll
		OR filterActivityTypeSend
	)
	AND (
		filterAllAddresses
		OR (
			HEX(pending_transactions.from_address) IN filter_addresses
		)
		OR (
			HEX(pending_transactions.to_address) IN filter_addresses
		)
	)
	AND (
		filterAllToAddresses
		OR (
			HEX(pending_transactions.to_address) IN filter_to_addresses
		)
	)
	AND (
		includeAllTokenTypeAssets
		OR (
			UPPER(pending_transactions.symbol) IN assets_token_codes
		)
	)
	AND (
		includeAllNetworks
		OR (
			pending_transactions.network_id IN filter_networks
		)
	)
UNION
ALL
SELECT
	NULL AS transfer_hash,
	NULL AS pending_hash,
	NULL AS network_id,
	multi_transactions.ROWID AS multi_tx_id,
	multi_transactions.timestamp AS timestamp,
	multi_transactions.type AS mt_type,
	NULL as tr_type,
	multi_transactions.from_address AS from_address,
	multi_transactions.to_address AS to_address,
	NULL AS owner_address,
	NULL AS tr_amount,
	multi_transactions.from_amount AS mt_from_amount,
	multi_transactions.to_amount AS mt_to_amount,
	CASE
		WHEN tr_status.min_status = 1
		AND COALESCE(pending_status.count, 0) = 0 THEN statusSuccess
		WHEN tr_status.min_status = 0 THEN statusFailed
		ELSE statusPending
	END AS agg_status,
	COALESCE(tr_status.count, 0) + COALESCE(pending_status.count, 0) AS agg_count,
	NULL AS token_address,
	NULL AS token_id,
	NULL AS token_code,
	multi_transactions.from_asset AS from_token_code,
	multi_transactions.to_asset AS to_token_code,
	multi_transactions.from_network_id AS out_network_id,
	multi_transactions.to_network_id AS in_network_id,
	NULL AS type,
	NULL as contract_address
FROM
	multi_transactions,
	filter_conditions
	LEFT JOIN tr_status ON multi_transactions.ROWID = tr_status.multi_transaction_id
	LEFT JOIN pending_status ON multi_transactions.ROWID = pending_status.multi_transaction_id
WHERE
	(
		(
			startFilterDisabled
			OR multi_transactions.timestamp >= startTimestamp
		)
		AND (
			endFilterDisabled
			OR multi_transactions.timestamp <= endTimestamp
		)
	)
	AND (
		filterActivityTypeAll
		OR (multi_transactions.type IN (%s))
	)
	AND (
		filterAllAddresses
		OR (
			HEX(multi_transactions.from_address) IN filter_addresses
		)
		OR (
			HEX(multi_transactions.to_address) IN filter_addresses
		)
	)
	AND (
		filterAllToAddresses
		OR (
			HEX(multi_transactions.to_address) IN filter_to_addresses
		)
	)
	AND (
		includeAllTokenTypeAssets
		OR (
			multi_transactions.from_asset != ''
			AND (
				UPPER(multi_transactions.from_asset) IN assets_token_codes
			)
		)
		OR (
			multi_transactions.to_asset != ''
			AND (
				UPPER(multi_transactions.to_asset) IN assets_token_codes
			)
		)
	)
	AND (
		filterAllActivityStatus
		OR (
			(
				filterStatusCompleted
				OR filterStatusFinalized
			)
			AND agg_status = statusSuccess
		)
		OR (
			filterStatusFailed
			AND agg_status = statusFailed
		)
		OR (
			filterStatusPending
			AND agg_status = statusPending
		)
	)
	AND (
		includeAllNetworks
		OR (
			multi_transactions.from_network_id IN filter_networks
		)
		OR (
			multi_transactions.to_network_id IN filter_networks
		)
		OR (
			COALESCE(multi_transactions.from_network_id, 0) = 0
			AND COALESCE(multi_transactions.to_network_id, 0) = 0
			AND (
				EXISTS (
					SELECT
						1
					FROM
						tr_network_ids
					WHERE
						multi_transactions.ROWID = tr_network_ids.multi_transaction_id
				)
				OR EXISTS (
					SELECT
						1
					FROM
						pending_network_ids
					WHERE
						multi_transactions.ROWID = pending_network_ids.multi_transaction_id
				)
			)
		)
	)
ORDER BY
	timestamp DESC
LIMIT
	? OFFSET ?