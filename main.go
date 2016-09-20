package main

import (
	"errors"
	"flag"
	"log"
	"runtime"
	"time"

	"github.com/resourcebench/job"
)

var (
	BenchType      string
	QiniuAccessKey string
	QiniuSecretKey string
	IntervalTime   int
	LastingTime    int
	CPU            = "cpu"
	MEM            = "mem"
	NET            = "net"
	IO             = "io"
	IDLE           = "idle"
)

func main() {

	flag.StringVar(&BenchType, "apptype", "", "the type for benchmark app, support cpu,mem,io,net,idle")
	flag.IntVar(&IntervalTime, "interval", 2000, "the interval to do add a load (this parameters could be used to adjust the rate)")
	flag.IntVar(&LastingTime, "lasting", 20000, "the time for adding the specific preassure")
	flag.StringVar(&QiniuAccessKey, "qiniuaccesskey", "", "qiniu accesskey")
	flag.StringVar(&QiniuSecretKey, "qiniusecretkey", "", "qiniu secrestkey")
	flag.Parse()
	setMaxProcs()

	//TODO check the parameter
	log.Println("the BenchType:", BenchType)
	err := CheckAndStart(BenchType)
	if err != nil {
		log.Println(err)
	}

}

func setMaxProcs() {
	// TODO(vmarmol): Consider limiting if we have a CPU mask in effect.
	// Allow as many threads as we have cores unless the user specified a value.
	var numProcs int

	numProcs = runtime.NumCPU()

	runtime.GOMAXPROCS(numProcs)

	// Check if the setting was successful.
	actualNumProcs := runtime.GOMAXPROCS(0)
	if actualNumProcs != numProcs {
		log.Println("Specified max procs of %v but using %v", numProcs, actualNumProcs)
	}
}

func CheckAndStart(benchType string) error {
	/*
		    CPU          = "cpu"
			MEM          = "mem"
			NET          = "net"
			IO           = "io"
			PACKET       = "packet"
			IDLE         = "idle"
	*/
	switch benchType {

	case CPU:
		job.CpuBenchmark(IntervalTime)
	case MEM:
		if LastingTime <= 0 {
			return errors.New("Lasting time could not be empty")
		}
		job.MemBenchmark(IntervalTime, LastingTime)

	case NET:

		if QiniuAccessKey == "" || QiniuSecretKey == "" {
			return errors.New("qiniu access key and qiniu secret key should be added")
		}

		job.NetBenchmark(IntervalTime, QiniuAccessKey, QiniuSecretKey)

	case IO:
		job.IOBenchmark(IntervalTime)

	case IDLE:
		//just out put some logs
		ticker := time.NewTicker(time.Second * 1)
		for t := range ticker.C {
			log.Println("Tick at", t)
		}
	default:
		return errors.New(benchType + "is not supported")

	}
	return nil

}
