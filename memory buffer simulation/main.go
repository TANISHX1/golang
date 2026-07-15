// A conceptual simulation of a memory buffer tracker built to demonstrate custom types,
// pointer receivers, and method indirection in Go.
package main

import (
	"fmt"
)

// 1. Custom non-struct type
type Bytes int64

// ToKB() method for Bytes (Value Receiver)
func (b Bytes) ToKB() float64 {
	return float64(b) / 1024.0
}

// 2. The Buffer Struct
type Buffer struct {
	Name  string
	Total Bytes
	Used  Bytes
}

// IsFull() method for Buffer (Value Receiver)
func (buf Buffer) IsFull() bool {
	if buf.Used >= buf.Total {
		return true
	}
	return false
}

// Write() method for Buffer (Pointer Receiver)
func (buf *Buffer) Write(size Bytes) bool {
	if (size + buf.Used) > buf.Total {
		fmt.Println("Buffer Overflow")
		return false
	} else {
		buf.Used += size
	}
	return true

}

func main() {
	// We create a standard VALUE variable (not a pointer)
	myBuf := Buffer{
		Name:  "L1_Cache",
		Total: 1024, // 1024 Bytes = 1 KB
		Used:  0,
	}
	myBuf.show()
	// TEST 1: Automatic Address-Of (&) Indirection
	// myBuf is a VALUE, but Write() expects a POINTER receiver (*Buffer).
	// Call myBuf.Write(512) directly without using '&'. Go should auto-convert it!
	if myBuf.Write(512) {
		fmt.Println("Succeeded")
	}
	myBuf.show()
	// TEST 2: Automatic Dereference (*) Indirection
	// We create a POINTER to myBuf
	ptrBuf := &myBuf

	// ptrBuf is a POINTER, but IsFull() expects a VALUE receiver (Buffer).
	// Call ptrBuf.IsFull() directly without using '*'. Go should auto-convert it!
	if ptrBuf.IsFull() {
		fmt.Println("buffer is Full")
	}
	fmt.Println("TO KB : ", ptrBuf.Used.ToKB())
}

func (myBuf Buffer) show() {
	fmt.Printf("Name:%s\nTotal :%d [%.3f] \nUsed: %d [%.3f]\n",
		myBuf.Name,
		myBuf.Total,
		myBuf.Total.ToKB(),
		myBuf.Used,
		myBuf.Used.ToKB())
}
