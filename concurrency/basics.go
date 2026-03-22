package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

// There is main go routine that execute program squentially.
// Go provides us construct -> Go Routine is like thread but very different than thread.
// Go routine is managed by go scheduler/go runtime not by operating system.
// Advantages of using concurrency primitives is that it is very light weight in comaprision to threads (in stack)

// What is concurrency ??
/* Managing the execution of different task effieciently so that it does not block system flow.*/
