// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"log"
	"net/http"
	"os"

)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	client := &http.Client{}

	request, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("BaseURL"),
		nil,
	)
	if err != nil{
		return err

	}
	request.Header.Set("Accept", "application/json; charset=utf-8")

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	log.Println(string(m.Data))
	return nil
}
