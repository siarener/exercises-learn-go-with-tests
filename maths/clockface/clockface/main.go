package main

import (
	"os"
	"time"

	"github.com/siarener/exercises-learn-go-with-tests/math/clockface" // REPLACE THIS!
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
