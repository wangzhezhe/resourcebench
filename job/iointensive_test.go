package job

import (
	"testing"
)

func TestIOBenchmark(t *testing.T) {
	t.Log("start to test cpu intensive app")
	interval := 100000
	IOBenchmark(interval)
}
