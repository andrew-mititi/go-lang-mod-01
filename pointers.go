package main

import (
	"fmt"
	"io"
)

type BufferedReader struct {
	reader io.Reader
	start int
	cursor int
}


func NewBufferedReader(reader io.Reader) *BufferedReader {
  return &BufferedReader{reader: reader, start: 0, cursor: 0}
}

func (nb *BufferedReader)encodeContents(contents string) error{
	contents_byte_array := []byte(contents)
	nb.reader.Read(contents_byte_array)
	fmt.Println(contents_byte_array)
	return nil
}




