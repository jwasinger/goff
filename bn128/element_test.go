// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by goff (v0.3.11) DO NOT EDIT

package bn128

import (
    "testing"
)

func BenchmarkMontMul_6limbs(b *testing.B) {
    x := Element{0x4d10c051c1fa23c0, 0x2945981a13aec13, 0x3bcea128c5c8d172, 0xdaa35e7a880a2ca}
    y := Element{0x4f0090b4b602e334, 0x670a33daa7a418b4, 0x8b9b1631a9ecad43, 0x15e1e13af71de992}
    var inv uint64
    inv = 0x89f3fffcfffcfffd
    mod := Element{0x6730d2a0f6b0f624, 0x64774b84f38512bf, 0x4b1ba7b6434bacd7, 0x1a0111ea397fe69a}

    for n := 0; n < b.N; n++ {
        x.Mul(&x, &y, &mod, inv)
    }
}
