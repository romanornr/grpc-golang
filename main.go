package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

//type Person struct {
//	Name string
//	Age int32
//}

func main() {
	elliot := &Person {
		Name: "Elliot",
		Age: 24,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Panic()
	}

	fmt.Println(data)

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		panic(err)
	}

	fmt.Println(newElliot.Age)
}
