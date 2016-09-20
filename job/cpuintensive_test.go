package job

import (
	"testing"
)

func TestCpuBentch(t *testing.T) {
	t.Log("start to test cpu intensive app")
	interval := 500
	CpuBenchmark(interval)
}
