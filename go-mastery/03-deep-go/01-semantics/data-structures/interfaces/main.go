package main

import (
	"fmt"
	"unsafe"
)

// DBCloser represents any system component that can safely tear down connections
type DBCloser interface {
	Close()
}

// SQLConnection is our concrete type
type SQLConnection struct {
	ID 	string
}

func (s SQLConnection) Close() {
	fmt.Printf("Closing SQL Connection: %s\n", s.ID)
}

// This struct mimics the internal runtime.iface layout exactly
type ifaceHeader struct {
	tab 	uintptr
	data 	uintptr
}

func main() {
	// Instantiate concrete type
	conn := SQLConnection{ID: "prod-cluster-01"}

	// Assign it to the interface (Creates the Fat Pointer)
	var closer DBCloser = conn

	// Unsafe cast to peek at the underlying 2-word structure
	header := (*ifaceHeader)(unsafe.Pointer(&closer))

	fmt.Printf("Interface Memory Location: %p\n", &closer)
	fmt.Printf("	-> itab Pointer (tab): 0x%x\n", header.tab)
	fmt.Printf("	-> Data Pointer (data): 0x%x\n", header.data)

	// Invoke the method via Dynamic Dispatch
	// runtime looks up Close() function inside header.tab and passes header.data in it
	closer.Close()
}