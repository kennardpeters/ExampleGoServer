package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"testing"

	"github.com/kennardpeters/ExampleGoServer/datastore"
	"github.com/kennardpeters/ExampleGoServer/server"
	"golang.org/x/net/websocket"
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

type runHandler struct {
	mu sync.Mutex
	n int
}

func (h *runHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	ra := r.RemoteAddr


	fmt.Fprintf(w, "ra: %s count is %d\n", ra, h.n)

	h.n++
}

func main() {

	wsServer := server.NewServer()
	http.Handle("/ws", websocket.Handler(wsServer.HandleWS))


	http.Handle("/count", new(countHandler))

	http.Handle("/run", new(runHandler))

	email, err := datastore.NewDatastore()
	if err != nil {
		panic(err)
	}

	log.Println(email)


	//msgChannel := make(chan int)
	//
	//wg := sync.WaitGroup{}

	//wg.Add(2)

	//publishingMessage := func() {

	//	for i := 0; i < 13; i++ {
	//		msgChannel <- i
	//	}
	//	close(msgChannel)
	//	wg.Done()
	//}

	//receivingMessage := func() {
	//	//loop through channel
	//	for number := range msgChannel {
	//		log.Println(number)
	//	}
	//	wg.Done()
	//}


	//go publishingMessage()

	//go receivingMessage()

	//wg.Wait()
	
	log.Fatal(http.ListenAndServe(":8333", nil))
}


func ExampleTest(t *testing.T) {

	t.Parallel()


	//Obtain a db connection here

	type testCase struct {
		name string
		data int
	}

	testCases := []testCase{
		{
			name: "test 1",
			data: 1,
		},
		{
			name: "test 1",
			data: 2,
		},
	}

	for i := range testCases {

		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			//wer
			t.Parallel()


		})
	}





	
	
}
