package main

import (
	"io"
	"os"
	"strconv"
	"time"
)

type Counter struct {
	Writer io.Writer
}

func (f *Counter) Count(n uint64) uint64 {
	if n == 0 {
		f.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}
	cur := n
	f.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	time.Sleep(1 * time.Second)
	return f.Count(n - 1)
}

func main() {
	pipeReader, pipeWriter := io.Pipe()
	defer pipeWriter.Close()
	defer pipeReader.Close()

	counter := Counter{
		Writer: pipeWriter,
	}

	file, _ := os.Create("counter.txt")
	tee := io.TeeReader(pipeReader, file)

	go func() {
		io.Copy(os.Stdout, tee)
	}()
	counter.Count(5)
	time.Sleep(1 * time.Second)
}
