package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
    filename := "large_input_1.txt"
    totalLines := 1_000_000      // 想生成多少行
    minValue := 1                     // 最小值（含）
    maxValue := 1_000_000             // 最大值（含）

    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Failed to create file:", err)
        return
    }
    defer file.Close()

    rand.Seed(time.Now().UnixNano())

    for i := 0; i < totalLines; i++ {
        number := rand.Intn(maxValue-minValue+1) + minValue
        _, err := fmt.Fprintf(file, "%d\n", number)
        if err != nil {
            fmt.Println("Failed to write line:", err)
            return
        }
    }

    fmt.Printf("✅ Successfully generated %s with %d random numbers between %d and %d\n", filename, totalLines, minValue, maxValue)
}
