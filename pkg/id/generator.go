package id

//package id implements several methods of generating id.
import (
	"math/rand"
	"time"
)

var (
	LowerLetters       = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Letters            = LowerLetters + UpperLetters
	Numbers            = "0123456789"
	LettersNumbers     = Letters + Numbers
	LowerLetterNumbers = LowerLetters + Numbers
	UpperLetterNumbers = UpperLetters + Numbers
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

//RandomStr returns a fixed-length(length=n) string.
func RandomStr(set string, n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(set) {
			b[i] = set[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

//GenerateRandomID returns a random string that starts with prefix.
func GenerateRandomID(prefix string) string {
	now := time.Now()
	b := make([]byte, 3)
	b[0] = Letters[now.Year()%len(Letters)]
	b[1] = Letters[now.Month()]
	b[2] = Letters[now.Day()]
	return prefix + string(b) + RandomStr(LettersNumbers, 5)
}
