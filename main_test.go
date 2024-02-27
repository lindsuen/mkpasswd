// mkpasswd - main_test.go
// Copyright (C) 2023 LindSuen <lindsuen@foxmail.com>
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"testing"

	"github.com/lindsuen/mkpasswd/src"
)

const (
	N uint64 = 999999
	l uint64 = 16
	n uint64 = 4
	c uint64 = 4
)

func BenchmarkMain(b *testing.B) {
	var i uint64
	for i = 0; i < N; i++ {
		src.GenerateRandomString(l, n, c)
	}
}
