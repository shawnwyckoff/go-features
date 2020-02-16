package main

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

// Safe close chan with sync.Mutex & select
func safeCloseDemo() {
	ch := make(chan struct{})
	var chLock sync.Mutex
	close(ch)

	// This is safe close method
	chLock.Lock()
	select {
	case <-ch: // Closed already
		chLock.Unlock()
		return
	default: // Not closed yet
		close(ch)
		chLock.Unlock()
	}
}

// UNSAFE close chan with sync.Mutex & select
func unsafeCloseDemo() {
	ch := make(chan struct{})
	var chLock sync.Mutex
	close(ch)

	// panic: close of closed channel
	chLock.Lock()
	defer chLock.Lock()
	close(ch)
}

// TODO: check whether safe when multi threads close channel simutaneously?
// Safe close chan with wrap function
func SafeCloseChanStruct(ch chan struct{}) (err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("BrokenPipe")
		}
	}()

	// assume ch != nil here.
	close(ch) // panic if ch is closed
	return nil
}

func main() {
	err := SafeCloseChanStruct(nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Close success")
	}
}
