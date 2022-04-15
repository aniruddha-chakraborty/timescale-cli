package services

import (
	"github.com/gosuri/uilive"
)

type Output struct {
	writer *uilive.Writer
}

func (o *Output) Init() {
	o.writer = uilive.New()
	o.writer.Start()
}

func (o *Output) Writer() *uilive.Writer {
	return o.writer
}