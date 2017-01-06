package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * time.Duration(10*rand.Intn(10)))
	fmt.Fprintf(w, "echo from server!")
	//s := r.Body
	//fmt.Println("the body of request :", s)
}

func main() {
	http.HandleFunc("/", sayhelloName)

	err := http.ListenAndServe("10.183.47.2:9998", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}

}
