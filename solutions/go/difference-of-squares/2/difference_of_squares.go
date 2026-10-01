package differenceofsquares

func SquareOfSum(n int) int {
	// sum := 0
 //    for i := 0; i <= n; i++ {
 //        sum += i
 //    }
 //    return sum * sum

    // someone told me about gauss's trick
    sum := n*(n+1) / 2
    return sum * sum
}

func SumOfSquares(n int) int {
	// square := 0
 //    for i := 0; i <= n; i++ {
 //        square += i * i
 //    }
 //    return square

    square := n * (n + 1) * (2 * n + 1) / 6
    return square
}

func Difference(n int) int {
	// successfully got this from O(N) to O(1)
	sqOfSum := SquareOfSum(n)
    sumOfSq := SumOfSquares(n)

    return sqOfSum - sumOfSq
}
