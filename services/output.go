package services

import (
	"bytes"
	"github.com/apoorvam/goterminal"
)

type Output struct {

}

func (c *Output) Writer() *goterminal.Writer {
	return goterminal.New(&bytes.Buffer{})
}