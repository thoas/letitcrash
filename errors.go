package letitgo

import (
	"bytes"
	"fmt"
	"io"
	"runtime/debug"

	"github.com/thoas/panicparse/stack"
)

// process copies stdin to stdout and processes any "panic: " line found.
func process(in io.Reader, out io.Writer, p *stack.Palette, s stack.Similarity, fullPath, parse bool) error {
	goroutines, err := stack.ParseDump(in, out)

	if err != nil {
		return err
	}
	if parse {
		stack.Augment(goroutines)
	}
	buckets := stack.SortBuckets(stack.Bucketize(goroutines, s))
	srcLen, pkgLen := stack.CalcLengths(buckets, fullPath)

	for _, bucket := range buckets {
		io.WriteString(out, p.BucketHeader(&bucket, fullPath, len(buckets) > 1))
		io.WriteString(out, p.StackLines(&bucket.Signature, srcLen, pkgLen, fullPath))
	}
	return err
}

func bufferFromError(err interface{}) *bytes.Buffer {
	// We handle pkg/errors.withStack struct
	if err, ok := err.(stackTracer); ok {
		buffer := new(bytes.Buffer)
		// Simulate a goroutine fake id to be correctly handled by panicparse
		buffer.WriteString("goroutine 0 [running]:")
		stackTrace := fmt.Sprintf("%+v", err.StackTrace())
		buffer.WriteString(stackTrace)

		return buffer
	}

	buffer := new(bytes.Buffer)
	buffer.WriteString(fmt.Sprintf("runtime error: %s%s%s\n", errorColor, err, resetFG))
	buffer.Write(debug.Stack())

	return buffer
}

// PrintError handles error and format it to stdout.
func PrintError(err interface{}, out io.Writer) error {
	s := stack.AnyPointer

	return process(bufferFromError(err), out, &defaultPalette, s, false, true)
}
