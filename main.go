package main

import (
	"lifeisFarm/handler/http"
	"sync"
)

func main() {

	go http.StartServer()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
