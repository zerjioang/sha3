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
	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/sha3"
	"testing"
)

const hextable = "0123456789abcdef"

type MethodSignatureHex [8]byte

func GenerateMethodSignature(h *sha3.State, fname []byte) string {
	_, _ = h.Write(fname)
	hash := h.Hash4()
	h.Reset()
	var dst MethodSignatureHex
	_ = hash[3]
	dst[0] = hextable[hash[0]>>4]
	dst[1] = hextable[hash[0]&0x0f]

	dst[2] = hextable[hash[1]>>4]
	dst[3] = hextable[hash[1]&0x0f]

	dst[4] = hextable[hash[2]>>4]
	dst[5] = hextable[hash[2]&0x0f]

	dst[6] = hextable[hash[3]>>4]
	dst[7] = hextable[hash[3]&0x0f]
	return string(dst[:])
}

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
	t.Run("use-multiple-hashers", func(t *testing.T) {
		var h = sha3.NewSha3()
		assert.Equal(t, GenerateMethodSignature(&h, []byte("f09140466846285922(address,bytes)")), "00000000")
		h1 := sha3.NewSha3()
		assert.Equal(t, GenerateMethodSignature(&h1, []byte("destroy()")), "83197ef0")
	})
	t.Run("reuse-hasher", func(t *testing.T) {
		var h = sha3.NewSha3()
		assert.Equal(t, GenerateMethodSignature(&h, []byte("f09140466846285922(address,bytes)")), "00000000")
		assert.Equal(t, GenerateMethodSignature(&h, []byte("destroy()")), "83197ef0")
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
