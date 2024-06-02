package main

import (
	"github.com/alexsuriano/stress-test/tester"
)

func main() {
	cfg := tester.GetParams()

	report := tester.StressTest(cfg)
	report.Presentation()
}
