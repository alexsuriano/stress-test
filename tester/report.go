package tester

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type Report struct {
	TotalTime          time.Duration
	TotalRequests      int64
	TotalStatusOk      int64
	TotalAnotherStatus int64
}

func (r *Report) Presentation() {

	color.HiWhite("Elapsed time: %v", r.TotalTime)

	totalRequests := strconv.Itoa(int(r.TotalRequests))
	fmt.Printf("Number of requests: %v\n", color.HiBlueString(totalRequests))

	statusOk := strconv.Itoa(int(r.TotalStatusOk))
	fmt.Printf("Total HTTP status OK: %v\n", color.HiGreenString(statusOk))

	anotherStatus := strconv.Itoa(int(r.TotalAnotherStatus))
	fmt.Printf("Total another HTTP status: %v\n", color.HiRedString(anotherStatus))
}
