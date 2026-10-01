package differenceofsquares

func SquareOfSum(n int) int {
	sum := 0
    for i := 0; i <= n; i++ {
        sum += i
    }
    return sum * sum
}

func SumOfSquares(n int) int {
	square := 0
    for i := 0; i <= n; i++ {
        square += i * i
    }
    return square
}

func Difference(n int) int {
	sqOfSum := SquareOfSum(n)
    sumOfSq := SumOfSquares(n)

    return sqOfSum - sumOfSq
}
