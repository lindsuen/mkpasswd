// mkpasswd - character.go
// Copyright (C) 2023 LindSuen <lindsuen@foxmail.com>
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package src

import "math/rand"

const (
	defaultString   string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberString    string = "1234567890"
	characterString string = "!@#$%^&*+?"
)

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

// GenerateRandomString generates a random string that matches the rules.
func GenerateRandomString(totalLength uint64, numberLength uint64, characterLength uint64) string {
	defFinalString := make([]rune, totalLength-numberLength-characterLength)
	numFinalString := make([]rune, numberLength)
	charFinalString := make([]rune, characterLength)

	defString := []rune(defaultString)
	numString := []rune(numberString)
	charString := []rune(characterString)

	selectCharacters(defString, &defFinalString)
	selectCharacters(numString, &numFinalString)
	selectCharacters(charString, &charFinalString)

	finalString := []rune(string(defFinalString) + string(numFinalString) + string(charFinalString))
	return shuffleString(&finalString)
}
