//
// Copyright zerjioang. 2022 All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package sha3_test

import (
	"encoding/hex"
	"fmt"
	"github.com/zerjioang/sha3"
	"testing"
)

func TestNewSha3(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		var h = sha3.NewSha3()
		fname := "approveTOP(bytes32,uint64)"
		_, _ = h.Write([]byte(fname))
		hash := h.Hash4()
		fmt.Println("function name:", fname)
		fmt.Println("sha3 hash value:", hex.EncodeToString(hash[:]))
		// Output: 24655e23
	})
	t.Run("loop-example", func(t *testing.T) {
		var h = sha3.NewSha3()
		fname := "approveTOP(bytes32,uint64)"
		for i := 0; i < 1000; i++ {
			_, _ = h.Write([]byte(fname))
			hash := h.Hash4()
			fmt.Println("function name:", fname)
			fmt.Println("sha3 hash value:", hex.EncodeToString(hash[:]))
			// Output: 24655e23
			h.Reset()
		}
	})
}

func BenchmarkNewSha3(b *testing.B) {
	// baseline
	// BenchmarkNewSha3/h-not-reuse-8         	 2558073	       467.8 ns/op	  55.58 MB/s	     480 B/op	       2 allocs/op
	b.Run("h-not-reuse", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		fname := "approveTOP(bytes32,uint64)"
		b.SetBytes(int64(len(fname)))
		for i := 0; i < b.N; i++ {
			var h = sha3.NewSha3()
			_, _ = h.Write([]byte(fname))
			hash := h.Hash4()
			if &hash == nil {
				panic("error")
			}
		}
	})
	// baseline
	// BenchmarkNewSha3/h-reuse-8             	 2981540	       401.0 ns/op	  64.83 MB/s	      32 B/op	       1 allocs/op
	b.Run("h-reuse", func(b *testing.B) {
		b.ReportAllocs()
		var h = sha3.NewSha3()
		fname := "approveTOP(bytes32,uint64)"
		b.ResetTimer()
		b.SetBytes(int64(len(fname)))
		for i := 0; i < b.N; i++ {
			_, _ = h.Write([]byte(fname))
			hash := h.Hash4()
			if &hash == nil {
				panic("error")
			}
			h.Reset()
		}
	})
}
