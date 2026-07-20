package stats

import "fmt"

type PipelineStats struct {
	BytesRead    int64
	BytesMasked  int64
	BytesWritten int64
}

func (p *PipelineStats) AddRead(n int) {
	p.BytesRead += int64(n)
}

func (p *PipelineStats) AddMasked(n int) {
	p.BytesMasked += int64(n)
}

func (p *PipelineStats) AddWritten(n int) {
	p.BytesWritten += int64(n)
}

func (p PipelineStats) String() string {
	return fmt.Sprintf(`=== STREAM-GUARD METRICS ===
Total Bytes Read    : [%05d]
Total Bytes Masked  : [%05d]
Total Bytes Written : [%05d]
============================`, p.BytesRead, p.BytesMasked, p.BytesWritten)
}
