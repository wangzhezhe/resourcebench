package job

import (
	"testing"
)

func TestIOBenchmark(t *testing.T) {
	t.Log("start to test net intensive app")
	interval := 1000
	NetBenchmark(interval)
}
