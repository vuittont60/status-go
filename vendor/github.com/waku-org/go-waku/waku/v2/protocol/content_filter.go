package protocol

import "golang.org/x/exp/maps"

type ContentTopicSet map[string]struct{}

func NewContentTopicSet(contentTopics ...string) ContentTopicSet {
	s := make(ContentTopicSet, len(contentTopics))
	for _, ct := range contentTopics {
		s[ct] = struct{}{}
	}
	return s
}

// ContentFilter is used to specify the filter to be applied for a FilterNode.
// Topic means pubSubTopic - optional in case of using contentTopics that following Auto sharding, mandatory in case of named or static sharding.
// ContentTopics - Specify list of content topics to be filtered under a pubSubTopic (for named and static sharding), or a list of contentTopics (in case ofAuto sharding)
// If pubSub topic is not specified, then content-topics are used to derive the shard and corresponding pubSubTopic using autosharding algorithm
type ContentFilter struct {
	PubsubTopic   string
	ContentTopics ContentTopicSet
}

func (cf ContentFilter) ContentTopicsList() []string {
	return maps.Keys(cf.ContentTopics)
}

func NewContentFilter(pubsubTopic string, contentTopics ...string) ContentFilter {
	return ContentFilter{pubsubTopic, NewContentTopicSet(contentTopics...)}
}
