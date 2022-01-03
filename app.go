package main

import (
	"fmt"
	"math/rand"
	"time"
)

func qsort(s []float64) []float64 {
    if len(s) < 2 {
        return s
    }

    left, right := 0, len(s)-1

    pivot := 0

    s[pivot], s[right] = s[right], s[pivot]

    for i := range s {
        if s[i] < s[right] {
            s[left], s[i] = s[i], s[left]
            left++
        }
    }

    s[left], s[right] = s[right], s[left]

    qsort(s[:left])
    qsort(s[left+1:])

    return s
}

func main() {
    fmt.Println("Starting")
    const x = 100000000
    const y = x * 10
    var s [y]float64
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    start1 := time.Now()
    for i := 0; i < y; i++ {
        s[i] = r1.NormFloat64()
    }
    end1 := time.Since(start1)
    ss := s[:]
    start2 := time.Now()
    ss = qsort(ss)
    end2 := time.Since(start2)
    fmt.Println(end1)
    fmt.Println(end2)
    fmt.Println("Number: ", ss[x])
}