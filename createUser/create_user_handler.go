package createRoom

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"flash.chat.com/createRoom/internal"
	"flash.chat.com/createRoom/internal/models"
)

var createUserFirestoreClient *firestore.Client

// runs outside of the function invocation, run only during a cold start.
func init() {
	cxt := context.Background()
	var err error
	projectId := os.Getenv("GCP_PROJECT")
	createUserFirestoreClient, err = firestore.NewClient(cxt, projectId)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	requestPayloadDecoder := json.NewDecoder(request.Body)
	defer request.Body.Close()

	var userRequest models.UserWebModel
	err := requestPayloadDecoder.Decode(&userRequest)

	if err != nil {
		requestBodyAsString, _ := ioutil.ReadAll(request.Body)
		fmt.Printf("Error parsing request with body %s, and error %s", requestBodyAsString, err)
		response.WriteHeader(http.StatusBadRequest)
	}

	docID := internal.GetDocId(userRequest.UserName)
	userDbModel := models.NewUserDBFromWebModel(userRequest)
	_, err = createUserFirestoreClient.Collection("users").Doc(docID).Create(request.Context(), userDbModel)
	if err != nil {
		fmt.Printf("Error Creating User, dup user with ID %s\n", docID)
		response.WriteHeader(http.StatusConflict)
		response.Header().Add("Content-Type", "application/json")
		response.Write([]byte(`{"message":"duplicate user"}`))
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte(""))
}
