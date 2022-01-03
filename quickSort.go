package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
    fmt.Println("Starting")
    const x = 10000000
    const y = x * 10
    var s [y]float64
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    start1 := time.Now()
    for i := 0; i < y; i++ {
        s[i] = r1.Float64()
    }
    end1 := time.Since(start1)
    ss := s[:]
    start2 := time.Now()
    sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
    end2 := time.Since(start2)
    fmt.Println(end1)
    fmt.Println(end2)
    fmt.Println("Number: ", ss[x])
}
