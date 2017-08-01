package http

import (
	"net/http"
	"fmt"
)

const(
	port = "7000"
)

func StartServer() {
	http.HandleFunc("/harvest", calHarvest)

	err := http.ListenAndServe("127.0.0.1:" + port, nil)
	if err != nil {
		fmt.Println("failed to start server:", err)
	}
	fmt.Println("server started")
}