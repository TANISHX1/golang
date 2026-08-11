package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	Reset           = "\033[0m"
	Bold            = "\033[1m"
	FgRed           = "\033[31m"
	FgGreen         = "\033[32m"
	FgBrightRed     = "\033[91m"
	FgBrightGreen   = "\033[92m"
	FgBrightYellow  = "\033[93m"
	FgBrightBlue    = "\033[94m"
	FgBrightMagenta = "\033[95m"
	FgBrightCyan    = "\033[96m"
)

type Job struct {
	job_id          int
	Processing_time time.Duration
}
type Metrics struct {
	Totaljobcompleted int
	Totaltimespent    time.Duration
	critical_sec      sync.Mutex
	waiting           sync.WaitGroup
}

func workers(worker_id int, ch1 chan Job, ch2 chan string, sysdata *Metrics) {

	for job := range ch1 {
		time.Sleep(job.Processing_time)
		sysdata.critical_sec.Lock()
		sysdata.Totaljobcompleted++
		sysdata.Totaltimespent += job.Processing_time
		sysdata.critical_sec.Unlock()
		ch2 <- fmt.Sprintf("\033[96m Worker [%d] Finished Job [%d] \033[0m", worker_id, job.job_id)
	}
	sysdata.waiting.Done()
}

func main() {

	buffered_ch := make(chan Job, 50)
	unbuffered_ch := make(chan string)

	metrics := Metrics{}
	for i := 0; i < 50; i++ {
		new := Job{
			job_id:          i,
			Processing_time: time.Duration(rand.Intn(91)+10) * time.Millisecond,
		}
		buffered_ch <- new
	}
	close(buffered_ch)

	go func() {
		metrics.waiting.Add(3)
		for i := 0; i < 3; i++ {
			go workers(i, buffered_ch, unbuffered_ch, &metrics)
		}
		metrics.waiting.Wait()
		close(unbuffered_ch)
	}()

	timeout := time.After(1500 * time.Millisecond)
	start := time.Now()
	for {
		select {
		case result, ok := <-unbuffered_ch:
			if !ok {
				fmt.Println(FgBrightGreen, "All Completed \n",
					"Total Jobs Done : ", metrics.Totaljobcompleted,
					"\nTotal TimeConsumed:", metrics.Totaltimespent,
					"\nActual time elasped:", time.Since(start))
				return
			}
			fmt.Println(result)

		case <-timeout:
			fmt.Println(FgBrightRed, "[TIMEOUT] | gorouteines are failed to complete Processing on time ", Reset)
			metrics.critical_sec.Lock()
			fmt.Println(FgBrightGreen, "All Completed \n",
				"Total Jobs Done : ", metrics.Totaljobcompleted,
				"\nTotal TimeConsumed:", metrics.Totaltimespent)
			metrics.critical_sec.Unlock()
			return
		}
	}
}
