// mkpasswd - main_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"testing"
)

const (
	l uint64 = 16
	n uint64 = 4
	c uint64 = 4
)

func BenchmarkCreateRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateRandomString(l, n, c)
	}
}
