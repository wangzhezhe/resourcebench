package main

import (
	//"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"sync"
	"time"
)

// it's better to use multiple client send the info at the same time

func main() {
	totalnum := 0
	//var wg sync.WaitGroup
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	url := "http://10.183.47.2:9998/"

	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	//num of packet in one second
	//var packetvar = flag.Int("num", 1000, "the number of packet in a second")
	//flag.Parse()
	//packetnum := *packetvar
	//fmt.Println("the packet num:", packetnum)
	//interval := time.Duration(1000000/packetnum) * time.Microsecond
	//fmt.Printf("interval: %f s ", interval.Seconds())

	ticker := time.NewTicker(time.Millisecond * 50)

	for t := range ticker.C {
		log.Println("tick at: ", t)
		response, err := client.Do(reqest)
		if err != nil {
			fmt.Println(err)
		}

		returnbody, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(returnbody))
		totalnum = totalnum + 1
		fmt.Println("return num ", totalnum)
	}

}
