package main

import (
	"fmt"
	"sync"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://localhost"
	url := "ws://localhost:23718/hello"
	ch := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go dialServer(i, origin, url, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}

func dialServer(id int, o, url string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	ws, err := websocket.Dial(url, "", o)
	if err != nil {
		fmt.Println("Error dialing socket: ", err)
	}
	msg := make([]byte, 1024)
	if _, err := ws.Read(msg); err != nil {
		fmt.Println("Error reading response from socket: ", err)
	}
	ch <- fmt.Sprintf("Routine id: %v, msg: %v", id, string(msg))
}
