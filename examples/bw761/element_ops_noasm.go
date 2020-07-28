// +build !amd64

// Copyright 2020 ConsenSys AG
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

// Code generated by goff (v0.3.1) DO NOT EDIT

// Package fp contains field arithmetic operations
package fp

// /!\ WARNING /!\
// this code has not been audited and is provided as-is. In particular,
// there is no security guarantees such as constant time implementation
// or side-channel attack resistance
// /!\ WARNING /!\

import "math/bits"

func Mul(z, x, y *Element) {
	_mulGeneric(z, x, y)
}

func Square(z, x *Element) {
	_squareGeneric(z, x)
}

// FromMont converts z in place (i.e. mutates) from Montgomery to regular representation
// sets and returns z = z * 1
func FromMont(z *Element) {
	_fromMontGeneric(z)
}

// Add z = x + y mod q
func Add(z, x, y *Element) {
	var carry uint64

	z[0], carry = bits.Add64(x[0], y[0], 0)
	z[1], carry = bits.Add64(x[1], y[1], carry)
	z[2], carry = bits.Add64(x[2], y[2], carry)
	z[3], carry = bits.Add64(x[3], y[3], carry)
	z[4], carry = bits.Add64(x[4], y[4], carry)
	z[5], carry = bits.Add64(x[5], y[5], carry)
	z[6], carry = bits.Add64(x[6], y[6], carry)
	z[7], carry = bits.Add64(x[7], y[7], carry)
	z[8], carry = bits.Add64(x[8], y[8], carry)
	z[9], carry = bits.Add64(x[9], y[9], carry)
	z[10], carry = bits.Add64(x[10], y[10], carry)
	z[11], _ = bits.Add64(x[11], y[11], carry)

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[11] < 81882988782276106 || (z[11] == 81882988782276106 && (z[10] < 15098257552581525310 || (z[10] == 15098257552581525310 && (z[9] < 13341377791855249032 || (z[9] == 13341377791855249032 && (z[8] < 5945444129596489281 || (z[8] == 5945444129596489281 && (z[7] < 8105254717682411801 || (z[7] == 8105254717682411801 && (z[6] < 274362232328168196 || (z[6] == 274362232328168196 && (z[5] < 9694500593442880912 || (z[5] == 9694500593442880912 && (z[4] < 8204665564953313070 || (z[4] == 8204665564953313070 && (z[3] < 10998096788944562424 || (z[3] == 10998096788944562424 && (z[2] < 1588918198704579639 || (z[2] == 1588918198704579639 && (z[1] < 16614129118623039618 || (z[1] == 16614129118623039618 && (z[0] < 17626244516597989515))))))))))))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 17626244516597989515, 0)
		z[1], b = bits.Sub64(z[1], 16614129118623039618, b)
		z[2], b = bits.Sub64(z[2], 1588918198704579639, b)
		z[3], b = bits.Sub64(z[3], 10998096788944562424, b)
		z[4], b = bits.Sub64(z[4], 8204665564953313070, b)
		z[5], b = bits.Sub64(z[5], 9694500593442880912, b)
		z[6], b = bits.Sub64(z[6], 274362232328168196, b)
		z[7], b = bits.Sub64(z[7], 8105254717682411801, b)
		z[8], b = bits.Sub64(z[8], 5945444129596489281, b)
		z[9], b = bits.Sub64(z[9], 13341377791855249032, b)
		z[10], b = bits.Sub64(z[10], 15098257552581525310, b)
		z[11], _ = bits.Sub64(z[11], 81882988782276106, b)
	}
}

// Double z = x + x mod q, aka Lsh 1
func Double(z, x *Element) {
	var carry uint64

	z[0], carry = bits.Add64(x[0], x[0], 0)
	z[1], carry = bits.Add64(x[1], x[1], carry)
	z[2], carry = bits.Add64(x[2], x[2], carry)
	z[3], carry = bits.Add64(x[3], x[3], carry)
	z[4], carry = bits.Add64(x[4], x[4], carry)
	z[5], carry = bits.Add64(x[5], x[5], carry)
	z[6], carry = bits.Add64(x[6], x[6], carry)
	z[7], carry = bits.Add64(x[7], x[7], carry)
	z[8], carry = bits.Add64(x[8], x[8], carry)
	z[9], carry = bits.Add64(x[9], x[9], carry)
	z[10], carry = bits.Add64(x[10], x[10], carry)
	z[11], _ = bits.Add64(x[11], x[11], carry)

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[11] < 81882988782276106 || (z[11] == 81882988782276106 && (z[10] < 15098257552581525310 || (z[10] == 15098257552581525310 && (z[9] < 13341377791855249032 || (z[9] == 13341377791855249032 && (z[8] < 5945444129596489281 || (z[8] == 5945444129596489281 && (z[7] < 8105254717682411801 || (z[7] == 8105254717682411801 && (z[6] < 274362232328168196 || (z[6] == 274362232328168196 && (z[5] < 9694500593442880912 || (z[5] == 9694500593442880912 && (z[4] < 8204665564953313070 || (z[4] == 8204665564953313070 && (z[3] < 10998096788944562424 || (z[3] == 10998096788944562424 && (z[2] < 1588918198704579639 || (z[2] == 1588918198704579639 && (z[1] < 16614129118623039618 || (z[1] == 16614129118623039618 && (z[0] < 17626244516597989515))))))))))))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 17626244516597989515, 0)
		z[1], b = bits.Sub64(z[1], 16614129118623039618, b)
		z[2], b = bits.Sub64(z[2], 1588918198704579639, b)
		z[3], b = bits.Sub64(z[3], 10998096788944562424, b)
		z[4], b = bits.Sub64(z[4], 8204665564953313070, b)
		z[5], b = bits.Sub64(z[5], 9694500593442880912, b)
		z[6], b = bits.Sub64(z[6], 274362232328168196, b)
		z[7], b = bits.Sub64(z[7], 8105254717682411801, b)
		z[8], b = bits.Sub64(z[8], 5945444129596489281, b)
		z[9], b = bits.Sub64(z[9], 13341377791855249032, b)
		z[10], b = bits.Sub64(z[10], 15098257552581525310, b)
		z[11], _ = bits.Sub64(z[11], 81882988782276106, b)
	}
}

// Sub  z = x - y mod q
func Sub(z, x, y *Element) {
	var b uint64
	z[0], b = bits.Sub64(x[0], y[0], 0)
	z[1], b = bits.Sub64(x[1], y[1], b)
	z[2], b = bits.Sub64(x[2], y[2], b)
	z[3], b = bits.Sub64(x[3], y[3], b)
	z[4], b = bits.Sub64(x[4], y[4], b)
	z[5], b = bits.Sub64(x[5], y[5], b)
	z[6], b = bits.Sub64(x[6], y[6], b)
	z[7], b = bits.Sub64(x[7], y[7], b)
	z[8], b = bits.Sub64(x[8], y[8], b)
	z[9], b = bits.Sub64(x[9], y[9], b)
	z[10], b = bits.Sub64(x[10], y[10], b)
	z[11], b = bits.Sub64(x[11], y[11], b)
	if b != 0 {
		var c uint64
		z[0], c = bits.Add64(z[0], 17626244516597989515, 0)
		z[1], c = bits.Add64(z[1], 16614129118623039618, c)
		z[2], c = bits.Add64(z[2], 1588918198704579639, c)
		z[3], c = bits.Add64(z[3], 10998096788944562424, c)
		z[4], c = bits.Add64(z[4], 8204665564953313070, c)
		z[5], c = bits.Add64(z[5], 9694500593442880912, c)
		z[6], c = bits.Add64(z[6], 274362232328168196, c)
		z[7], c = bits.Add64(z[7], 8105254717682411801, c)
		z[8], c = bits.Add64(z[8], 5945444129596489281, c)
		z[9], c = bits.Add64(z[9], 13341377791855249032, c)
		z[10], c = bits.Add64(z[10], 15098257552581525310, c)
		z[11], _ = bits.Add64(z[11], 81882988782276106, c)
	}
}

// Neg z = q - x
func Neg(z, x *Element) {
	if x.IsZero() {
		z.SetZero()
		return
	}
	var borrow uint64
	z[0], borrow = bits.Sub64(17626244516597989515, x[0], 0)
	z[1], borrow = bits.Sub64(16614129118623039618, x[1], borrow)
	z[2], borrow = bits.Sub64(1588918198704579639, x[2], borrow)
	z[3], borrow = bits.Sub64(10998096788944562424, x[3], borrow)
	z[4], borrow = bits.Sub64(8204665564953313070, x[4], borrow)
	z[5], borrow = bits.Sub64(9694500593442880912, x[5], borrow)
	z[6], borrow = bits.Sub64(274362232328168196, x[6], borrow)
	z[7], borrow = bits.Sub64(8105254717682411801, x[7], borrow)
	z[8], borrow = bits.Sub64(5945444129596489281, x[8], borrow)
	z[9], borrow = bits.Sub64(13341377791855249032, x[9], borrow)
	z[10], borrow = bits.Sub64(15098257552581525310, x[10], borrow)
	z[11], _ = bits.Sub64(81882988782276106, x[11], borrow)
}
