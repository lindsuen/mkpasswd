// mkpasswd - main.go
// Copyright (C) 2023 LindSuen <lindsuen@foxmail.com>
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"

	s "github.com/lindsuen/mkpasswd/src"
)

func main() {
	password := new(s.Password)
	s.SetParameter(password)
	s.HandleInputError(password.Length, password.Number, password.Character)

	var counter uint64
	for counter = 0; counter < password.Quantity; counter++ {
		fmt.Println(s.GenerateRandomString(password.Length, password.Number, password.Character))
	}
}
