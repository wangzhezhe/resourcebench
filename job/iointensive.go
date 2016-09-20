package job

import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"
)

//you could adjust the interval value according to your cpu core num to get different cpu growth rate
func IOBenchmark(interval int) {
	var wg sync.WaitGroup
	log.Println("start io intensive testing")

	buffer := make([]byte, 1024*1024*100)
	after := time.After(time.Millisecond * time.Duration(interval))
	wg.Add(3)
	go ReadLoad(after, buffer)
	go WriteLoad(after, buffer)
	go Ticker()
	wg.Wait()

}

//read the data from /dev/sda1
func ReadLoad(ch <-chan time.Time, buffer []byte) {
	f, err := os.Open("/dev/sda1")
	defer f.Close()
	if err != nil {
		log.Println("failed to open /dev/sda1")
	}
	reader := bufio.NewReader(f)

A:
	for {
		select {
		case <-ch:
			log.Println("job finished")
			break A
		default:
			//log.Println("transfer")
			_, err := reader.Read(buffer)
			time.Sleep(time.Millisecond * 100)
			if err != nil {
				log.Println(err)
			}

		}
	}
}

func WriteLoad(ch <-chan time.Time, buffer []byte) {
	f, err := os.Create("temp.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

A:
	for {
		select {
		case <-ch:
			log.Println("job finished")
			break A
		default:
			writer := bufio.NewWriter(f)
			_, err := writer.Write(buffer)
			if err != nil {
				log.Println(err)
			}
			//os.Chmod("temp.txt",)
			os.Truncate("temp.txt", 0)
			time.Sleep(time.Millisecond * 300)
		}
	}

}

//keep the process runing for a long time
func Ticker() {
	ticker := time.NewTicker(time.Second * 5)
	for t := range ticker.C {
		log.Println("Tick at", t)
	}
}
