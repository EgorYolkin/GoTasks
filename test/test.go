package main

import (
	"fmt"
	"time"
)

func main() {
    ch := make(chan string, 1)
    channels(ch)
    select {
        case msg := <- ch:
            fmt.Println(msg)
    }
}


func channels(ch chan string) {
    time.Sleep(time.Second * 10)
    ch <- "Test!"
}
