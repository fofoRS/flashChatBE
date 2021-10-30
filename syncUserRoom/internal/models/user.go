package models

import (
	"errors"
	"fmt"
	"time"

	"flash.chat.com/createRoom/internal"
)

type UserWebModel struct {
	UserName string `json:"userName"`
}

type UserDBModel struct {
	UserName  string            `firestore:"username"`
	Timestamp time.Time         `firestore:"timestamp,serverTimestamp"`
	Rooms     []UserRoomDBModel `firestore:"rooms,omitempty"`
}

type UserRoomDBModel struct {
	RoomName string `firestore:"roomName"`
	ImageURL string `firestore:"imageUrl,omitempty"`
}

func NewUserDBFromWebModel(webModel UserWebModel) UserDBModel {
	return UserDBModel{UserName: webModel.UserName}
}

func NewUserRoomFromRoomCreateFields(roomCreatedFields internal.RoomCreatedFieldMap) (*UserRoomDBModel, error) {
	nameFieldMap, ok := roomCreatedFields["name"]
	var userRoom *UserRoomDBModel
	if ok {
		nameValue, ok := nameFieldMap["stringValue"]
		if ok {
			valueAsString := fmt.Sprint(nameValue)
			userRoom = &UserRoomDBModel{RoomName: valueAsString}
		}
	} else {
		return nil, errors.New("Room name is empty")
	}
	return userRoom, nil
}
