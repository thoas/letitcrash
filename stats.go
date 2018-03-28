package letitcrash

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

var start = time.Now()

const mb float64 = 1.0 * 1024 * 1024

func getServerInformation() map[string]interface{} {
	mem := &runtime.MemStats{}
	runtime.ReadMemStats(mem)

	return map[string]interface{}{
		"Uptime":                 fmt.Sprintf("%d seconds", getUptime()),
		"Allocated memory":       fmt.Sprintf("%.2f MB", toMegaBytes(mem.Alloc)),
		"Total allocated memory": fmt.Sprintf("%.2f MB", toMegaBytes(mem.TotalAlloc)),
		"Number of Goroutines":   runtime.NumGoroutine(),
		"Number of CPUs":         runtime.NumCPU(),
	}
}

func getUptime() int64 {
	return time.Now().Unix() - start.Unix()
}

func toMegaBytes(bytes uint64) float64 {
	return toFixed(float64(bytes)/mb, 2)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
