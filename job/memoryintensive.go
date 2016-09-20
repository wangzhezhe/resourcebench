package job

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

//you could adjust the interval value according to your cpu core num to get different cpu growth rate

var memMap = make(map[string][]int)

//interval(ms) is the time frequency to excute the mem allocation
//memIncreasing(ms) is the time for excuting time allocation
//you could adjust to these parameter accordin to memory on your machine
//during the time of the memory allocating, the cpu value will also be kept in a high position
func MemBenchmark(interval int, memIncreasing int) {

	log.Println("start mem intensive testing")

	log.Println("the lasting time is ", memIncreasing)

	globalAfter := time.After(time.Millisecond * time.Duration(memIncreasing))

	ticker := time.NewTicker(time.Millisecond * time.Duration(interval))
A:
	for t := range ticker.C {
		log.Println("Tick at", t)
		//create the basic load
		select {
		case <-globalAfter:
			break A
		default:
			loadAfter := time.After(time.Millisecond * 1000)
			go MemAllocationLoad(loadAfter)
		}

	}
	for t := range ticker.C {
		log.Println("stop to allocate mem")
		//prevent the memMap to be reclaimed by gc
		_ = memMap
		log.Println("Tick at", t)
	}

}

func MemAllocationLoad(ch <-chan time.Time) {
	randomValue := rand.Intn(10000)
	key := "test" + "_" + strconv.Itoa(randomValue)
	value := make([]int, 1024*1024*800)
	memMap[key] = value
A:
	for {
		select {
		case <-ch:
			log.Println("job finished")
			//delete(memMap, key)
			break A
		default:

		}
	}
}
