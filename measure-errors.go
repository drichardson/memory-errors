package main

import (
	"log"
	"syscall"
	"time"
)

const GB = 1 << 30
const allocationSize = 8 * GB

func main() {
	memory := make([]byte, allocationSize)

	err := syscall.Mlock(memory)
	if err != nil {
		log.Fatalln("Error locking memory.", err)
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
