package twofer
import fmt "fmt"

// ShareWith describes who we're sharing with, by name if we have their name, by saying "you" if we don't.
func ShareWith(name string) string {
    if name != "" {
        return fmt.Sprintf("One for %s, one for me.", name)
    } else {
        return "One for you, one for me."
    }
}
