package main

import (
	"farmerCalendar/handler/http"
	"sync"
)

func main() {

	go http.StartServer()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
