package isbnverifier
import "strings"

func IsValidISBN(isbn string) bool {
    parts := strings.Split(isbn, "-")
    joined := strings.Join(parts, "")
    arrayOfChars := strings.Split(joined, "")

    if len(arrayOfChars) != 10 {
        return false
    }

    var sum int
    multiplier := 10

    for i := 0; i < len(arrayOfChars); i++ {
        var val int
        if arrayOfChars[i] == "X" {
            if i != 9 { 
                return false
            } else {
            	val = 10
            }
        } else {
            c := arrayOfChars[i][0]
            if c < '0' || c > '9' {
                return false
            }
            val = int(c - '0')
        }
		sum += val * multiplier
        multiplier--
    }

    return sum % 11 == 0
}

    
