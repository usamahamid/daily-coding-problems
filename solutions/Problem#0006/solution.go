package main

import (
	"unsafe"
)

// Node of doubly linked list
type Node struct {
	Data    string
	Pointer uintptr
}

// Doubly linked list using Xor to store Pointer
type XorList struct {
	Head uintptr
	Last uintptr
}

func (list XorList) insert(data string) {
	newNode := Node{}
	newNode.Data = data
	if list.Head != 0 {
		newNode.Pointer = xor(list.Head, uintptr(unsafe.Pointer(&newNode)))
	} else {

	}

}

func xor(prev uintptr, next uintptr) uintptr {
	return prev ^ next
}

func main() {
	// list := XorList{}

}
