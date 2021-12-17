package main

import (
	"fmt"
	"hello/tasks/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Task Management")
	r := router.Router()
	fmt.Println("Server is getting started..Start your test  now")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
	fmt.Println("-------------------- ...")
}
