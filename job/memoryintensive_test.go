package job

import (
	"testing"
)

func TestMemBenchmark(t *testing.T) {
	t.Log("start to test cpu intensive app")
	interval := 10000
	lastingTime := 20000
	MemBenchmark(interval, lastingTime)
}
