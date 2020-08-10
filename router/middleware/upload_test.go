package middleware

import (
	"github.com/Shopify/sarama"
	"testing"
)

func TestUploadUserInfo(t *testing.T) {
	config := sarama.NewConfig()
	producer, err := sarama.NewAsyncProducer([]string{"47.110.255.149:9092"}, config)
	if err != nil {
		t.Fatal(err)
	}
	message := &sarama.ProducerMessage{Topic: "user_request_upload", Value: sarama.StringEncoder("{\"test\":001}")}
	producer.Input() <- message
}

func TestHandleError(t *testing.T) {
	a := map[string][]string{"man": {"you", "me"}, "woman": {"she", "her"}}

	for i, v := range a {
		t.Log(i)
		t.Log(v)
	}

}
