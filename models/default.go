package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// ResponseJSON is Struct of response JSON
	ResponseJSON struct {
		RequestID bson.ObjectId `json:"requestId"`
		Message   string        `json:"message"`
		Code      int           `json:"code"`
		Timestamp int64        	`json:"timestamp"`
		Data      string        `json:"data"` // data will mounted as json string
	}
)
