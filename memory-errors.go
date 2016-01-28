package main

import (
	"flag"
	"log"
	"syscall"
	"time"
)

const GB = 1 << 30

var allocationInGB = flag.Uint("size", 2, "allocation size in GB")

func main() {
	flag.Parse()
	allocationSize := *allocationInGB * GB
	memory := make([]byte, allocationSize)

	// mlock prevents memory from being swapped to disk
	err := syscall.Mlock(memory)
	if err != nil {
		log.Fatalln("Error locking memory.", err)
	}

	// mprotect PROT_READ makes the memory read only. Writes on memory[i] after the syscall
	// result in a fault that terminates the program.
	err = syscall.Mprotect(memory, syscall.PROT_READ)
	if err != nil {
		log.Fatalln("Error protecting memory.", err)
	}

	for {
		log.Println("Scanning ", allocationSize, " bytes of memory...")
		for i, v := range memory {
			if v != 0 {
				log.Println("Found non-zero byte at position", i)
			}
		}
		log.Println("Scan completed")
		time.Sleep(5 * time.Minute)
	}
}
