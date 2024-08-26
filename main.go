// mkpasswd - main.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

const (
	defaultString   string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberString    string = "1234567890"
	characterString string = "!@#$%^&*+?"
)

var (
	defString  = []rune(defaultString)
	numString  = []rune(numberString)
	charString = []rune(characterString)
)

// password defines the property of a password.
type password struct {
	quantity  uint64
	length    uint64
	number    uint64
	character uint64
}

// setParameter assigns values from input parameters to the password object.
func setParameter(p *password) {
	flag.Uint64Var(&p.quantity, "N", 1, "The quantity of created passwords.")
	flag.Uint64Var(&p.length, "l", 16, "The length of password.")
	flag.Uint64Var(&p.number, "n", 4, "The number of Arabic numerals (0-9).")
	flag.Uint64Var(&p.character, "c", 4, "The number of special characters (!@#$%^&*+?).")
	flag.Parse()
}

// handleInputError handles the error of input parameters.
// The value of -n and -c must be less than the value of -l. And, the sum of -n and -c must
// be less than the value of -l.
func handleInputError(totalLength uint64, numLength uint64, charLength uint64) {
	if numLength >= totalLength || charLength >= totalLength {
		fmt.Println("The value of -n and -c must be less than the value of -l.")
		os.Exit(1)
	} else if numLength+charLength >= totalLength {
		fmt.Println("The sum of -n and -c must be less than the value of -l.")
		os.Exit(1)
	}
}

// selectCharacters obtains the corresponding number of characters randomly.
func selectCharacters(s []rune, l *[]rune) {
	for n := range *l {
		(*l)[n] = s[rand.Intn(len(s))]
	}
}

// shuffleString shuffles the final combined string.
func shuffleString(r *[]rune) string {
	for i := len(*r) - 1; i >= 0; i-- {
		num := rand.Intn(len(*r))
		(*r)[i], (*r)[num] = (*r)[num], (*r)[i]
	}
	return string(*r)
}

// createRandomString creates a random string that matches the rules.
func createRandomString(totalLength uint64, numberLength uint64, characterLength uint64) string {
	defFinalString := make([]rune, totalLength-numberLength-characterLength)
	numFinalString := make([]rune, numberLength)
	charFinalString := make([]rune, characterLength)

	selectCharacters(defString, &defFinalString)
	selectCharacters(numString, &numFinalString)
	selectCharacters(charString, &charFinalString)

	finalString := append(append(defFinalString, numFinalString...), charFinalString...)
	return shuffleString(&finalString)
}

func main() {
	password := new(password)
	setParameter(password)
	handleInputError(password.length, password.number, password.character)

	var counter uint64
	for counter = 0; counter < password.quantity; counter++ {
		fmt.Println(createRandomString(password.length, password.number, password.character))
	}
}
