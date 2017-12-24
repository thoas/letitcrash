package call911

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/maruel/panicparse/stack"
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

// PrintError handles error and format it to stdout.
func PrintError(thr interface{}) error {
	var out io.Writer = os.Stdout

	s := stack.AnyPointer

	buffer := new(bytes.Buffer)
	buffer.WriteString(fmt.Sprintf("runtime error: %s%s%s\n", errorColor, thr, resetFG))
	buffer.Write(debug.Stack())

	return process(buffer, out, &defaultPalette, s, false, true)
}
