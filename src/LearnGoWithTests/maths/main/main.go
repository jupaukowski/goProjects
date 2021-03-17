// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"os"
	"test/maths/svg"
	"time"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
