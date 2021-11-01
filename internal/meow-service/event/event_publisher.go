package event

type EventPublisher interface {
	Publish(event Event) error
}
