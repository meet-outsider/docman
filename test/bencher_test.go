package test

import (
	"fmt"
	"testing"
	"time"
)
func TestBencher(t *testing.T) {
	startTime := time.Now()
    var sum float64
    for i := 1; i <= 1e7; i++ {
        sum += float64(i)
    }
    endTime := time.Now()
    fmt.Printf("Result: %v\n", sum)
    fmt.Printf("Time elapsed: %v\n", endTime.Sub(startTime))
}