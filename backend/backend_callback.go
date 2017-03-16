package main

import (
	"fmt"
	"time"
	"trib"
	"triblab"
)

func main() {
	fmt.Println("Hello")
	var binClient = triblab.NewBinClient([]string{"localhost:12461"})
	var succ bool
	var client = binClient.Bin("rkapoor")
	client.Set(&trib.KeyValue{Key: "clock:123", Value: "false"}, &succ)
	var ret string //make(chan string, 1)
	client.Get("clock:123", &ret)

	select {
	case r := <-ret:
		fmt.Println("Got value", r)

	case <-time.After(10 * time.Second):
		fmt.Println("Timed out after 10 Seconds")
	}

	// for {
	// 	fmt.Println("Waiting")
	// 	time.Sleep(3 * time.Second)

	// 	fmt.Println("Got value:", ret)
	// }
}
