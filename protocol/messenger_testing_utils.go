package protocol

import (
	"context"
	"errors"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/status-im/status-go/protocol/common"
	"github.com/status-im/status-go/protocol/protobuf"
	"github.com/status-im/status-go/protocol/tt"
)

// WaitOnMessengerResponse Wait until the condition is true or the timeout is reached.
func WaitOnMessengerResponse(m *Messenger, condition func(*MessengerResponse) bool, errorMessage string) (*MessengerResponse, error) {
	response := &MessengerResponse{}
	err := tt.RetryWithBackOff(func() error {
		var err error
		r, err := m.RetrieveAll()
		if err != nil {
			panic(err)
		}

		if err := response.Merge(r); err != nil {
			panic(err)
		}

		if err == nil && !condition(response) {
			err = errors.New(errorMessage)
		}
		return err
	})
	return response, err
}

type MessengerSignalsHandlerMock struct {
	MessengerSignalsHandler

	responseChan chan *MessengerResponse
}

func (m *MessengerSignalsHandlerMock) MessengerResponse(response *MessengerResponse) {
	// Non-blocking send
	select {
	case m.responseChan <- response:
	default:
	}
}

func (m *MessengerSignalsHandlerMock) MessageDelivered(chatID string, messageID string) {}

func WaitOnSignaledMessengerResponse(m *Messenger, condition func(*MessengerResponse) bool, errorMessage string) (*MessengerResponse, error) {
	interval := 500 * time.Millisecond
	timeoutChan := time.After(10 * time.Second)

	if m.config.messengerSignalsHandler != nil {
		return nil, errors.New("messengerSignalsHandler already provided/mocked")
	}

	responseChan := make(chan *MessengerResponse, 1)
	m.config.messengerSignalsHandler = &MessengerSignalsHandlerMock{
		responseChan: responseChan,
	}

	defer func() {
		m.config.messengerSignalsHandler = nil
	}()

	for {
		_, err := m.RetrieveAll()
		if err != nil {
			return nil, err
		}

		select {
		case r := <-responseChan:
			if condition(r) {
				return r, nil
			}
			return nil, errors.New(errorMessage)

		case <-timeoutChan:
			return nil, errors.New("timed out: " + errorMessage)

		default: // No immediate response, rest & loop back to retrieve again
			time.Sleep(interval)
		}
	}
}

func FindFirstByContentType(messages []*common.Message, contentType protobuf.ChatMessage_ContentType) *common.Message {
	for _, message := range messages {
		if message.ContentType == contentType {
			return message
		}
	}
	return nil
}

func PairDevices(s *suite.Suite, device1, device2 *Messenger) {
	// Send pairing data
	response, err := device1.SendPairInstallation(context.Background(), nil)
	s.Require().NoError(err)
	s.Require().NotNil(response)
	s.Len(response.Chats(), 1)
	s.False(response.Chats()[0].Active)

	i, ok := device1.allInstallations.Load(device1.installationID)
	s.Require().True(ok)

	// Wait for the message to reach its destination
	response, err = WaitOnMessengerResponse(
		device2,
		func(r *MessengerResponse) bool {
			for _, installation := range r.Installations {
				if installation.ID == device1.installationID {
					return installation.InstallationMetadata != nil &&
						i.InstallationMetadata.Name == installation.InstallationMetadata.Name &&
						i.InstallationMetadata.DeviceType == installation.InstallationMetadata.DeviceType
				}
			}
			return false

		},
		"installation not received",
	)
	s.Require().NoError(err)
	s.Require().NotNil(response)

	// Ensure installation is enabled
	err = device2.EnableInstallation(device1.installationID)
	s.Require().NoError(err)
}
