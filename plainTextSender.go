package main

import (
	"bytes"
	"math/rand"
	"time"
)

type PlainTextReq struct {
	lines []string
}

func (p PlainTextReq) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	for _, line := range p.lines {
		buf.WriteString(line)
		buf.WriteString("\n")
	}
	return buf.Bytes(), nil
}

func (p PlainTextReq) Size() int {
	return len(p.lines)
}

func NewPlainTextSender(opts LogSenderOpts) ISender {
	l := &GenericSender{
		LogSenderOpts: opts,
		rnd:           rand.New(rand.NewSource(time.Now().UnixNano())),
		timeout:       time.Second,
		path:          "/test-lines",
	}
	if l.Headers == nil {
		l.Headers = make(map[string]string)
	}
	l.Headers["Content-Type"] = "text/plain"
	l.generate = func() IRequest {
		req := PlainTextReq{
			lines: make([]string, opts.LinesPS),
		}
		for i := 0; i < opts.LinesPS; i++ {
			req.lines[i] = l.pickRandom(opts.Lines)
		}
		return req
	}
	return l
}
