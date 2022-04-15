package services

import (
	"bytes"
	"github.com/apoorvam/goterminal"
)

type Output struct {
	writer *goterminal.Writer
}

func (o *Output) Init() {
	o.writer = goterminal.New(&bytes.Buffer{})
}

func (o *Output) Writer() *goterminal.Writer {
	return o.writer
}