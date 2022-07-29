package nats

type INats interface {
	Publish(subject string, message []byte) error
}
