package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

var (
	numBurn        int
	updateInterval int
)

func cpuBurn() {
	for {
		for i := 0; i < 2147483647; i++ {
		}

		// Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
		// Gosched让当前goroutine让出处理器，从而使得其它goroutine可以运行。它不会挂起/暂停当前的goroutine，它会自动恢复执行。
		runtime.Gosched()
	}
}

func init() {
	flag.IntVar(&numBurn, "n", 0, "number of cores to burn (0 = all)")
	flag.IntVar(&updateInterval, "u", 10, "seconds between updates (0 = don't update)")
	flag.Parse()
	if numBurn <= 0 {
		numBurn = runtime.NumCPU()
	}
}

func main() {
	runtime.GOMAXPROCS(numBurn)
	fmt.Printf("Burning %d CPUs/cores\n", numBurn)
	for i := 0; i < numBurn; i++ {
		go cpuBurn()
	}
	if updateInterval > 0 {
		t := time.Tick(time.Duration(updateInterval) * time.Second)
		for secs := updateInterval; ; secs += updateInterval {
			<-t
			fmt.Printf("%d seconds\n", secs)
		}
	} else {
		select {} // wait forever
	}
}
