package src

import "math/rand"

const (
	defaultLetters   string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberLetters    string = "1234567890"
	characterLetters string = "!@#$%^&*+?"
)

func randCharacter(letter []rune, l *[]rune) []rune {
	for n := range *l {
		(*l)[n] = letter[rand.Intn(len(letter))]
	}
	return *l
}

func shuffleArray(r *[]rune) []rune {
	for i := len(*r) - 1; i >= 0; i-- {
		num := rand.Intn(len(*r))
		(*r)[i], (*r)[num] = (*r)[num], (*r)[i]
	}
	return *r
}

func RandomString(Len uint64, numLen uint64, charLen uint64) string {
	defl := make([]rune, Len-numLen-charLen)
	numl := make([]rune, numLen)
	charl := make([]rune, charLen)

	defLetters := []rune(defaultLetters)
	numLetters := []rune(numberLetters)
	charLetters := []rune(characterLetters)

	randCharacter(defLetters, &defl)
	randCharacter(numLetters, &numl)
	randCharacter(charLetters, &charl)

	finalLetters := []rune(string(defl) + string(numl) + string(charl))
	return string(shuffleArray(&finalLetters))
}
