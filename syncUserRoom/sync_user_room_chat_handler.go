package createRoom

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	"flash.chat.com/createRoom/internal"
	"flash.chat.com/createRoom/internal/models"
)

var syncUserChatFirestoreClient *firestore.Client

// runs outside of the function invocation, run only during a cold start.
func init() {
	cxt := context.Background()
	var err error
	projectId := os.Getenv("GCP_PROJECT")
	syncUserChatFirestoreClient, err = firestore.NewClient(cxt, projectId)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}

func SyncUserOnRoomCreated(ctx context.Context, roomCreatedEvent internal.RoomCreatedEvent) error {

	userEmailAddress, ok := roomCreatedEvent.Value.Fields["owner"]["stringValue"]
	if !ok {
		return errors.New("user email address missing, cannot sync room with user")
	}

	userRoom, err := models.NewUserRoomFromRoomCreateFields(roomCreatedEvent.Value.Fields)
	if err != nil {
		return err
	}

	// _, result, err :=

	userCollRef := syncUserChatFirestoreClient.Collection("users")
	userRoomCollRef := userCollRef.Doc(internal.GetDocId(fmt.Sprint(userEmailAddress))).Collection("userRooms")
	result, err := userRoomCollRef.Doc(getCreatedRoomId(roomCreatedEvent.Value)).Create(ctx, userRoom)
	if err != nil {
		return err
	}

	fmt.Printf("Document for user %s updated at %s", fmt.Sprint(userEmailAddress), result.UpdateTime)
	return nil
}

func getCreatedRoomId(currentDocument internal.RoomCreatedValue) string {
	fullDocumentPath := currentDocument.Name
	documentPathTokens := strings.Split(fullDocumentPath, "/")
	documentID := documentPathTokens[len(documentPathTokens)-1:][0]
	return documentID
}
