package main

import (
	 "fmt"
		"bytes"
)

func main() {
		var wc WriterCloser = NewBufferedWriterCloser() // Create the wc that the type is one interface called WriterCloser
		wc.Write([]byte("Hello YouTube listeners, this is a test")) // Acess my method "Write" inside my wc interface 
		wc.Close() // Acess my method "Close" inside my wc interface 
}

type Writer interface { //Create my Writer Interface
	 Write([]byte) (int, error) // Parameters need be type []byte and the method needs to return int, error
}

type Closer interface { //Create my Closer Interface
	 Close() error // The method needs to return one error
}

type WriterCloser interface { // Multiple interfaces inside other interface
	 Writer
		Closer
}

type BufferedWriterCloser struct { // Creating my BufferedWriterCloser struct
	 buffer *bytes.Buffer // It has a buffer variable with the type bytes.Buffer and one pointer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) { // The function Write use the bwc with type BufferedWriterCloser as paramater too
		 myResult, err := bwc.buffer.Write(data) // Try to write the data value on my bwc variable and see if is equal with the type
			if err != nil { // Check if don't have error
				 return 0, err
			}

			v := make([]byte, 8) // make a new byte slice
			for bwc.buffer.Len() > 8 {
					_, err := bwc.buffer.Read(v)
					if err != nil {
						 return 0, err
					}
					_, err = fmt.Println(string(v))
					if err != nil {
						 return 0, err
					}
			}
			return myResult, nil
}

func (bwc *BufferedWriterCloser) Close() error {
		for bwc.buffer.Len() > 0 {
			 data := bwc.buffer.Next(8)
				_, err := fmt.Println(string(data))
				if err != nil {
					 return err
				}
		}
		return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser { //The function needs to return one BufferedWriterCloser struct
	 return &BufferedWriterCloser {
			 buffer: bytes.NewBuffer([]byte{}), 
		}
}
