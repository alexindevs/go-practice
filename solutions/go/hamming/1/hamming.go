package hamming
import "errors"

func Distance(a, b string) (int, error) {
    hammingCount := 0
    if len(a) != len(b) {
		return 0, errors.New("a and b must be the same length")
	}
	for i := range a {
	    if a[i] != b[i] {
            hammingCount++
        }
	}

    return hammingCount, nil
}
