package main

import (
	"fmt"
	"time"

	allInOneBroker "github.com/sainipankaj15/All-In-One-Broker"
)

func main() {
	fmt.Println("Hello World!")
	time.Sleep(5 * time.Second)

	allInOneBroker.ExitAllPosition_Tiqs("ABCD12")

}
