package math32

import "math"

// The Go standard library only contains float64 math functions. Implementing
// the most commonly used of these in 32 bit variants gives a slight speed
// improvement, and a huge code readability improvement.

func Minf(a, b float32) float32

func minf(a, b float32) float32 {
    return float32(math.Min(float64(a), float64(b)))
}

func Maxf(a, b float32) float32

func maxf(a, b float32) float32 {
    return float32(math.Max(float64(a), float64(b)))
}

func sqrtf(num float32) float32 {
    return float32(math.Sqrt(float64(num)))
}

func Sqrtf(num float32) float32
