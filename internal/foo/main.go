package foo

import (
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
)

func Start() {
	publisher, err := nats.NewPublisher(nats.PublisherConfig{
		URL:       "nats://localhost:4222",
		Marshaler: nats.GobMarshaler{},
	}, watermill.NewStdLogger(false, false))
	if err != nil {
		panic(err)
	}

	publishMessages(publisher)
}

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))

		if err := publisher.Publish("example-topic", msg); err != nil {
			panic(err)
		}
		log.Printf("service=foo msg=published message to topic message_id=%s", msg.UUID)

		time.Sleep(time.Second)
	}
}
