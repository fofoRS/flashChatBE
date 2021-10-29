package interfaces

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
)

type firestoreClient struct {
	client firestore.Client
}

func NewFirestoreClient() *firestoreClient {
	return &firestoreClient{}
}

// this method is the entry point to start interacting with the firestore client and its API
func (f *firestoreClient) initFirestore(ctx context.Context, projectId string) {
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}
	f.client = *client
}

func (f *firestoreClient) Client() (*firestore.Client, error) {
	if &f.client == nil {
		return nil, errors.New("connection to firestore service is not stablished yet, please call initFirestore method first")
	}
	return &f.client, nil
}
