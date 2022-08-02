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
		// Output: 24655e23
	})
}
