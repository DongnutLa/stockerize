package shared_domain

type MessagingTopics string

const (
	HandleStockTopic    MessagingTopics = "stock.handle"
	HandleOrdersSummary MessagingTopics = "orders.summary.handle"
)

var TopicList = []string{
	string(HandleStockTopic),
	string(HandleOrdersSummary),
}

type MessageEvent struct {
	EventTopic MessagingTopics `json:"eventTopic"`
	Topic      AwsTopics
	Data       map[string]interface{} `json:"data"`
}

type AwsTopics string

const (
	HandleEventsTopic AwsTopics = "handle-events"
)
