package job

import (
	"log"
	"time"
)

//you could adjust the rate value according to your cpu core num to get different cpu growth rate
func CpuBenchmark(rate int) {

	log.Println("start cpu intensive testing")
	if (rate - 1) < 0 {
		log.Fatal("rate should more than 1")
	}

	//ticker := time.NewTicker(time.Millisecond * time.Duration(500+(rate-1)*100))

	for {
		//log.Println("Tick at", t)
		//create the basic load
		after := time.After(time.Millisecond * time.Duration(100))
		go AddOneLoad(after)
		time.Sleep(time.Millisecond * time.Duration(rate))
	}

}

//this is the basic load unit which plus 1 to 100
func AddOneLoad(ch <-chan time.Time) {
	count := 0
A:
	for {
		select {
		case <-ch:
			log.Println("job finished")
			break A
		default:
			count = count + 1
		}
	}
}
