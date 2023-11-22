// mkpasswd - password.go
// Copyright (C) 2023 LindSuen <lindsuen@foxmail.com>
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package src

// Password defines the property of a password.
type Password struct {
	Quantity  uint64
	Length    uint64
	Number    uint64
	Character uint64
}
