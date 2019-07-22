package main

import (
	"fmt"
	"time"
)

func loadDinner(fd fullDinner) {
	// Push to Remote data source
	time.Sleep(1250 * time.Millisecond)

	fmt.Println("Sent:", fd.name)
}
