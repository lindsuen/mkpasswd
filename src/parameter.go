// mkpasswd - parameter.go
// Copyright (C) 2023 LindSuen <lindsuen@foxmail.com>
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package src

import "flag"

// SetParameter assigns values from input parameters to the password object.
func SetParameter(p *Password) {
	flag.Uint64Var(&p.Quantity, "N", 1, "The quantity of created passwords.")
	flag.Uint64Var(&p.Length, "l", 16, "The length of password.")
	flag.Uint64Var(&p.Number, "n", 4, "The number of Arabic numerals (0-9).")
	flag.Uint64Var(&p.Character, "c", 4, "The number of special characters (!@#$%^&*+?).")
	flag.Parse()
}
