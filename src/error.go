// mkpasswd - error.go
// Copyright (C) 2023 LindSuen <lindsuen@foxmail.com>
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package src

import (
	"fmt"
	"os"
)

// HandleInputError handles the error of input parameters. The values of -n and -c must be less than the value of -l.
func HandleInputError(totalLength uint64, numLength uint64, charLength uint64) {
	if numLength > totalLength || charLength > totalLength {
		fmt.Println("The values of -n and -c must be less than the value of -l.")
		os.Exit(1)
	}
}
