package main

import (
	"encoding/json"
	"fmt"
	"github.com/shawnwyckoff/go-utils/dsa/jsons"
)

// 基础类型（比如int）及其type出来的新类型，支持omitempty
// 结构体及其type出来的新类型，不支持omitempty
// 结构体S实现omitempty的唯一方法就是，在使用它的结构体P中显式的使用它的指针*S作为成员

type InInfo struct {
	Address string
}

type Info InInfo

func (info Info) MarshalJSON() ([]byte, error) {
	/*
		// 和Score一样的实现，但是搁结构体就是不支持空[]byte的omitempty
		if info.Address == "" {
			return []byte{}, nil
		}
		if info.Address == "" {
			return []byte(""), nil
		}*/
	return json.Marshal(InInfo(info))
}

/*
// build error: invalid receiver type Show (Show is a pointer type)

type Score *int

func (s Score) MarshalJSON() ([]byte, error) {
	if int(s) == 0 {
		return []byte{}, nil
	}
	return []byte(fmt.Sprintf("%d", int(s))), nil
}*/

type Score int

func (s Score) MarshalJSON() ([]byte, error) {
	if int(s) == 0 {
		return []byte{}, nil
	}
	return []byte(fmt.Sprintf("%d", int(s))), nil
}

func main() {
	type S struct {
		Name   string
		Score  Score `json:"Score,omitempty"`
		Info   Info  `json:"Info,omitempty"`
		InfoEx *Info `json:"InfoEx,omitempty"`
	}

	d := S{Name: "Bob", Score: 0}
	fmt.Println(jsons.MarshalStringDefault(d, false))
}
