package job

import (
	"log"
	"time"
)

//you could adjust the interval value according to your cpu core num to get different cpu growth rate
func Idle(interval int) {

	log.Println("start idle testing")

	ticker := time.NewTicker(time.Millisecond * time.Duration(interval))

	for t := range ticker.C {
		log.Println("Tick at", t)
	}

}
