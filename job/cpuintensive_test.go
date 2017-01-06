package job

import (
	"testing"
)

func TestCpuBentch(t *testing.T) {
	t.Log("start to test cpu intensive app")
	//2 core cpu rate=500 %cpu=20%
	rate := 100
	CpuBenchmark(rate)
}
