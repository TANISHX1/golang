package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"stream-guard/sink"
	"stream-guard/stats"
	"stream-guard/stream"
)

func main() {
	// status instance created
	status := stats.PipelineStats{}

	// Storage file
	file, err := os.Create("backup.log")
	if err != nil {
		fmt.Println("[ Error ] Failed to open file ")
		return
	}
	defer file.Close()
	// setup sinks
	sinks := sink.MultiSink{}
	sinks.Destinations = append(sinks.Destinations, file, os.Stdout)

	// setup masker
	masker := stream.MaskingReader{os.Stdin, 0}
	buffer := make([]byte, 1024)
	for {
		n, readerr := masker.Read(buffer)
		if n > 0 {
			status.AddRead(n)
			status.AddMasked(bytes.Count(buffer[:n], []byte("*")))
		}
		written, writerr := sinks.Write(buffer[:n])
		if writerr != nil {
			fmt.Println("[Error] Write Failed : ", writerr)
			break
		}
		status.AddWritten(written)
		if readerr != nil {

			if readerr == io.EOF {
				break
			}
			fmt.Println("[Error] Read Fail: ", readerr)
			break
		}
	}
	fmt.Println(status)
}
