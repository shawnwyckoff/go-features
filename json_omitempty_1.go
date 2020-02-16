package main

import (
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"log"
	"strconv"
)

/*
correct:
`json:"Bar,omitempty" bson:"Bar,omitempty"`

bad syntax
`json:"Bar,omitempty", bson:"Bar,omitempty"` // , between json and bson shouldn't be there
`json:"Bar,omitempty" bson:"Bar, omitempty"` // space between "Bar," and "omitempty" shouldn't be there
*/

type Bar struct {
	a int64
}

type Foo struct {
	Bar *Bar `json:"Bar,omitempty" bson:"Bar,omitempty"`
}

func (b Bar) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bsonx.String(fmt.Sprintf("%d", b.a)).MarshalBSONValue()
}

func (b *Bar) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	if t != bsontype.String {
		return errors.Errorf("bsontype(%s) not allowed in Bar.UnmarshalBSONValue, only string accept", t.String())
	}
	str, _, ok := bsoncore.ReadString(data)
	if !ok {
		return errors.Errorf("decode string, but string not found")
	}

	dec, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	b = &Bar{a: dec}
	return nil
}

func main() {
	f := Foo{}
	_, err := bson.Marshal(f)
	if err != nil {
		log.Fatalf("Failed to marshal. Error: %v", err)
	}
}
