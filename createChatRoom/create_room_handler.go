package createRoom

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"flash.chat.com/createRoom/internal/models"
)

var firestoreClient *firestore.Client

// runs outside of the function invocation, run only during a cold start.
func init() {
	cxt := context.Background()
	var err error
	projectId := os.Getenv("GCP_PROJECT")
	firestoreClient, err = firestore.NewClient(cxt, projectId)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}

func CreateRoomChat(response http.ResponseWriter, request *http.Request) {
	requestBodyDecoder := json.NewDecoder(request.Body)
	var roomRequest models.RoomWebModel
	err := requestBodyDecoder.Decode(&roomRequest)
	if err != nil {
		bodyDataBytes, _ := ioutil.ReadAll(request.Body)
		log.Printf("Error Parsing request body with body %s and error %s", string(bodyDataBytes), err)
		errorPayload := []byte(`{"message": "Invalid Request, cannot be parsed"}`)
		callToResponse(errorPayload, http.StatusBadRequest, response)
		return
	}

	if len(roomRequest.Name) > 100 {
		errorPayload := []byte(`{"message": "Lenght of room name exceeded, please try a name less than 100 characters"}`)
		callToResponse(errorPayload, http.StatusBadRequest, response)
		return
	}
	roomDBModel := models.RoomDBModel{Name: roomRequest.Name, Owner: roomRequest.Owner}
	savedRoom, _, err := firestoreClient.Collection("rooms").Add(request.Context(), roomDBModel)

	if err != nil {
		log.Print(err)
		errorPayload := []byte(`{"message": "Error creating room chat, try again."}`)
		callToResponse(errorPayload, http.StatusInternalServerError, response)
		return
	}

	var roomResponseModel = models.RoomWebModel{Name: roomRequest.Name, Id: savedRoom.ID}
	responseData, _ := json.Marshal(roomResponseModel)
	callToResponse(responseData, http.StatusOK, response)

}

func callToResponse(body []byte, status int, response http.ResponseWriter) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(status)
	if body != nil {
		response.Write(body)
		return
	}
	return

}
