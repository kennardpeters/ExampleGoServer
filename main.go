package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type countHandler struct {
	mu sync.Mutex
	n int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {


	http.Handle("/count", new(countHandler))


	msgChannel := make(chan int)
	
	wg := sync.WaitGroup{}

	wg.Add(2)

	publishingMessage := func() {

		for i := 0; i < 13; i++ {
			msgChannel <- i
		}
		close(msgChannel)
		wg.Done()
	}

	receivingMessage := func() {
		//loop through channel
		for number := range msgChannel {
			log.Println(number)
		}
		wg.Done()
	}


	go publishingMessage()

	go receivingMessage()

	wg.Wait()
	
	log.Fatal(http.ListenAndServe(":8080", new(countHandler)))
}
