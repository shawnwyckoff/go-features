package main

import (
	"bytes"
	"github.com/Workiva/go-datastructures/queue"
)

func main() {
	bb := bytes.NewBuffer([]byte{})
	rb := queue.NewRingBuffer(1024)

	count := 10000
	for i := 0; i < count; i++ {
		bb.Write([]byte("good"))
	}
	for i := 0; i < count; i++ {
		bb.Read()
	}
}
