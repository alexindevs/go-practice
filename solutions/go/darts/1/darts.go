package darts
import "math"

func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)
    var score int
    if distance <= 1 {
        score = 10
    } else if distance <= 5 {
        score = 5
    } else if distance <= 10 {
        score = 1
    }
    return score
}
