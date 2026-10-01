package collatzconjecture
import "errors"

func CollatzConjecture(n int) (int, error) {
    if n < 1 {
        return 0, errors.New("n must be an integer greater than 0")
    }
	var counter int
    currentNum := n

    for currentNum != 1 {
        if currentNum % 2 == 0 {
            currentNum /= 2
        } else {
            currentNum = (currentNum * 3) + 1
        }
        counter++
    }
    return counter, nil
}
