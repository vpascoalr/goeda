package bar

import (
	"context"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
)

func Start() {
	subscriber, err := nats.NewSubscriber(
		nats.SubscriberConfig{
			URL:              "nats://localhost:4222",
			CloseTimeout:     30 * time.Second,
			AckWaitTimeout:   30 * time.Second,
			QueueGroupPrefix: "bar",
			Unmarshaler:      nats.GobMarshaler{},
			JetStream: nats.JetStreamConfig{
				AutoProvision: true,
			},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "example-topic")
	if err != nil {
		panic(err)
	}

	go process(messages)
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("service=bar msg=received message message_id=%s payload=%s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
