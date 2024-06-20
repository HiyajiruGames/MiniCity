package util

import (
	"fmt"
	"runtime"

	"github.com/HiyajiruGames/MiniCity/hgpkg/log"
)

// PrintMemUsage outputs the current, total and OS memory being used. As well as the
// number of garage collection cycles completed. For info on each,
// see: https://golang.org/pkg/runtime/#MemStats
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Debug("******************************")
	log.Debug(fmt.Sprintf("Alloc = %v bytes\nTotalAlloc = %v bytes\nSys = %v bytes\nNumGC = %v", m.Alloc, m.TotalAlloc, m.Sys, m.NumGC))
	log.Debug("******************************")
}
