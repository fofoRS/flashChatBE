package internal

import "time"

// FirestoreEvent is the payload of a Firestore event.
type RoomCreatedEvent struct {
	OldValue   RoomCreatedValue `json:"oldValue"`
	Value      RoomCreatedValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type RoomCreatedValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log the interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     RoomCreatedFieldMap `json:"fields"`
	Name       string              `json:"name"`
	UpdateTime time.Time           `json:"updateTime"`
}

type RoomCreatedFieldMap map[string]map[string]interface{}
