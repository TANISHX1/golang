package sink

import (
	"fmt"
)

type Sink interface {
	Write(p []byte) (n int, err error)
}

type MultiSink struct {
	Destinations []Sink
}

func (r *MultiSink) Write(p []byte) (n int, err error) {

	for _, f := range r.Destinations {
		_, err = f.Write(p)
		if err != nil {
			return 0, err
		}
	}
	return len(p), nil
}

func FormatMetadata(key string, val any) string {
	switch v := val.(type) {
	case string:
		return fmt.Sprintf("[%s: %s]", key, v)
	case int:
		return fmt.Sprintf("[%s (CODE:%d) ] ", key, v)
	default:
		return fmt.Sprintf("[%s:%v] ", key, v)
	}
}
