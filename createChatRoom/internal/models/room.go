package models

import "time"

type RoomWebModel struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type RoomDBModel struct {
	Name      string    `firestore:"name"`
	Owner     string    `firestore:"owner"`
	Timestamp time.Time `firestore:"timestamp,serverTimestamp"`
	ImageUrl  string    `firestore:"imageUrl,omitempty"`
}
