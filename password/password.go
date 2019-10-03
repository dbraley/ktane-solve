package password

var DefaultWords = []string{
	"about", "after", "again", "below", "could",
	"every", "first", "found", "great", "house",
	"large", "learn", "never", "other", "place",
	"plant", "point", "right", "small", "sound",
	"spell", "still", "study", "their", "there",
	"these", "thing", "think", "three", "water",
	"where", "which", "world", "would", "write",
}

func Solve(possibleIndexLetters map[int][]rune, possibleWords []string) []string {
	for letterIndex, possibleRunes := range possibleIndexLetters {
		possibleWords = keepIf(possibleWords, func(word string) bool {
			if len(word) <= letterIndex {
				return false
			}
			return contains(possibleRunes, rune(word[letterIndex]))
		})
	}
	return possibleWords
}

func keepIf(s []string, f func(string) bool) []string {
	i := 0 // output index
	for _, x := range s {
		if f(x) {
			// copy and increment index
			s[i] = x
			i++
		}
	}
	return s[:i]
}

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
