package main

import (
	"flag"
	"log"
	"syscall"
)

const GB = 1 << 30

var allocationInGB = flag.Uint("size", 2, "allocation size in GB")

func main() {
	flag.Parse()
	allocationSize := *allocationInGB * GB

	log.Printf("Allocating 0x%X bytes\n", allocationSize)
	memory := make([]byte, allocationSize)

	// mlock prevents memory from being swapped to disk
	err := syscall.Mlock(memory)
	if err != nil {
		log.Fatalln("Error locking memory.", err)
	}

	var value byte

	for {
		log.Printf("Writing 0x%X to 0x%X bytes.\n", value, allocationSize)
		for i, _ := range memory {
			memory[i] = value
		}
		log.Printf("Verifying %x written to %x bytes.\n", value, allocationSize)
		for i, v := range memory {
			if v != value {
				log.Printf("Found unexpected byte value 0x%X at position %d.\n", v, i)
			}
		}
		value++
	}
}
