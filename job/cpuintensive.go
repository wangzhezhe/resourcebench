package job

import (
	"log"
	"time"
)

//you could adjust the interval value according to your cpu core num to get different cpu growth rate
func CpuBenchmark(interval int) {

	log.Println("start cpu intensive testing")

	ticker := time.NewTicker(time.Millisecond * time.Duration(interval))

	for t := range ticker.C {
		log.Println("Tick at", t)
		//create the basic load
		after := time.After(time.Millisecond * 1000)
		go AddOneLoad(after)
	}

}

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
