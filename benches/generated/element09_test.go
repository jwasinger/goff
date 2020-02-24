// Code generated by goff DO NOT EDIT
package generated

import (
	"crypto/rand"
	"math/big"
	"math/bits"
	mrand "math/rand"
	"testing"
)

func TestELEMENT09CorrectessAgainstBigInt(t *testing.T) {
	modulus, _ := new(big.Int).SetString("103184574917104427901333969349794958333929702908907610310000668649441237878132450269633384409814565451349830727596073204454861419191038748430243571780797913131162618958407779", 10)
	cmpEandB := func(e *Element09, b *big.Int, name string) {
		var _e big.Int
		if e.FromMont().ToBigInt(&_e).Cmp(b) != 0 {
			t.Fatal(name, "failed")
		}
	}
	var modulusMinusOne, one big.Int
	one.SetUint64(1)

	modulusMinusOne.Sub(modulus, &one)

	for i := 0; i < 10000; i++ {

		// sample 2 random big int
		b1, _ := rand.Int(rand.Reader, modulus)
		b2, _ := rand.Int(rand.Reader, modulus)
		rExp := mrand.Uint64()

		// adding edge cases
		// TODO need more edge cases
		switch i {
		case 0:
			rExp = 0
			b1.SetUint64(0)
		case 1:
			b2.SetUint64(0)
		case 2:
			b1.SetUint64(0)
			b2.SetUint64(0)
		case 3:
			rExp = 0
		case 4:
			rExp = 1
		case 5:
			rExp = ^uint64(0) // max uint
		case 6:
			rExp = 2
			b1.Set(&modulusMinusOne)
		case 7:
			b2.Set(&modulusMinusOne)
		case 8:
			b1.Set(&modulusMinusOne)
			b2.Set(&modulusMinusOne)
		}

		rbExp := new(big.Int).SetUint64(rExp)

		var bMul, bAdd, bSub, bDiv, bNeg, bLsh, bInv, bExp, bSquare big.Int

		// e1 = mont(b1), e2 = mont(b2)
		var e1, e2, eMul, eAdd, eSub, eDiv, eNeg, eLsh, eInv, eExp, eSquare, eMulAssign, eSubAssign, eAddAssign Element09
		e1.SetBigInt(b1)
		e2.SetBigInt(b2)

		// (e1*e2).FromMont() === b1*b2 mod q ... etc
		eSquare.Square(&e1)
		eMul.Mul(&e1, &e2)
		eMulAssign.Set(&e1)
		eMulAssign.MulAssign(&e2)
		eAdd.Add(&e1, &e2)
		eAddAssign.Set(&e1)
		eAddAssign.AddAssign(&e2)
		eSub.Sub(&e1, &e2)
		eSubAssign.Set(&e1)
		eSubAssign.SubAssign(&e2)
		eDiv.Div(&e1, &e2)
		eNeg.Neg(&e1)
		eInv.Inverse(&e1)
		eExp.Exp(e1, rExp)
		eLsh.Double(&e1)

		// same operations with big int
		bAdd.Add(b1, b2).Mod(&bAdd, modulus)
		bMul.Mul(b1, b2).Mod(&bMul, modulus)
		bSquare.Mul(b1, b1).Mod(&bSquare, modulus)
		bSub.Sub(b1, b2).Mod(&bSub, modulus)
		bDiv.ModInverse(b2, modulus)
		bDiv.Mul(&bDiv, b1).
			Mod(&bDiv, modulus)
		bNeg.Neg(b1).Mod(&bNeg, modulus)

		bInv.ModInverse(b1, modulus)
		bExp.Exp(b1, rbExp, modulus)
		bLsh.Lsh(b1, 1).Mod(&bLsh, modulus)

		cmpEandB(&eSquare, &bSquare, "Square")
		cmpEandB(&eMul, &bMul, "Mul")
		cmpEandB(&eMulAssign, &bMul, "MulAssign")
		cmpEandB(&eAdd, &bAdd, "Add")
		cmpEandB(&eAddAssign, &bAdd, "AddAssign")
		cmpEandB(&eSub, &bSub, "Sub")
		cmpEandB(&eSubAssign, &bSub, "SubAssign")
		cmpEandB(&eDiv, &bDiv, "Div")
		cmpEandB(&eNeg, &bNeg, "Neg")
		cmpEandB(&eInv, &bInv, "Inv")
		cmpEandB(&eExp, &bExp, "Exp")
		cmpEandB(&eLsh, &bLsh, "Lsh")
	}
}

func TestELEMENT09IsRandom(t *testing.T) {
	for i := 0; i < 1000; i++ {
		var x, y Element09
		x.SetRandom()
		y.SetRandom()
		if x.Equal(&y) {
			t.Fatal("2 random numbers are unlikely to be equal")
		}
	}
}

// -------------------------------------------------------------------------------------------------
// benchmarks
// most benchmarks are rudimentary and should sample a large number of random inputs
// or be run multiple times to ensure it didn't measure the fastest path of the function
// TODO: clean up and push benchmarking branch

var benchResElement09 Element09

func BenchmarkInverseELEMENT09(b *testing.B) {
	var x Element09
	x.SetRandom()
	benchResElement09.SetRandom()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchResElement09.Inverse(&x)
	}

}
func BenchmarkExpELEMENT09(b *testing.B) {
	var x Element09
	x.SetRandom()
	benchResElement09.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.Exp(x, mrand.Uint64())
	}
}

func BenchmarkDoubleELEMENT09(b *testing.B) {
	benchResElement09.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.Double(&benchResElement09)
	}
}

func BenchmarkAddELEMENT09(b *testing.B) {
	var x Element09
	x.SetRandom()
	benchResElement09.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.Add(&x, &benchResElement09)
	}
}

func BenchmarkSubELEMENT09(b *testing.B) {
	var x Element09
	x.SetRandom()
	benchResElement09.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.Sub(&x, &benchResElement09)
	}
}

func BenchmarkNegELEMENT09(b *testing.B) {
	benchResElement09.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.Neg(&benchResElement09)
	}
}

func BenchmarkDivELEMENT09(b *testing.B) {
	var x Element09
	x.SetRandom()
	benchResElement09.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.Div(&x, &benchResElement09)
	}
}

func BenchmarkFromMontELEMENT09(b *testing.B) {
	benchResElement09.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.FromMont()
	}
}

func BenchmarkToMontELEMENT09(b *testing.B) {
	benchResElement09.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.ToMont()
	}
}
func BenchmarkSquareELEMENT09(b *testing.B) {
	benchResElement09.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.Square(&benchResElement09)
	}
}

func BenchmarkMulAssignELEMENT09(b *testing.B) {
	x := Element09{
		13830955647730413669,
		6927004744728180,
		6111362518727439073,
		1139408670260740882,
		3895080857423388830,
		17074512152868828260,
		8192131671142038703,
		16959221168518730559,
		4996576493377651554,
	}
	benchResElement09.SetOne()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.MulAssign(&x)
	}
}

// Montgomery multiplication benchmarks
func (z *Element09) mulCIOS(x *Element09) *Element09 {

	var t [10]uint64
	var D uint64
	var m, C uint64
	// -----------------------------------
	// First loop

	C, t[0] = bits.Mul64(x[0], z[0])
	C, t[1] = madd1(x[0], z[1], C)
	C, t[2] = madd1(x[0], z[2], C)
	C, t[3] = madd1(x[0], z[3], C)
	C, t[4] = madd1(x[0], z[4], C)
	C, t[5] = madd1(x[0], z[5], C)
	C, t[6] = madd1(x[0], z[6], C)
	C, t[7] = madd1(x[0], z[7], C)
	C, t[8] = madd1(x[0], z[8], C)

	D = C

	// m = t[0]n'[0] mod W
	m = t[0] * 13715081754621093557

	// -----------------------------------
	// Second loop
	C = madd0(m, 1897235049931565155, t[0])

	C, t[0] = madd2(m, 13948212332252695022, t[1], C)

	C, t[1] = madd2(m, 5931880109645070248, t[2], C)

	C, t[2] = madd2(m, 13767938734030060565, t[3], C)

	C, t[3] = madd2(m, 1578249351545375860, t[4], C)

	C, t[4] = madd2(m, 13837360922177919458, t[5], C)

	C, t[5] = madd2(m, 13195475758251816711, t[6], C)

	C, t[6] = madd2(m, 18260081051345282122, t[7], C)

	C, t[7] = madd3(m, 7695857179358191508, t[8], C, t[9])

	t[8], t[9] = bits.Add64(D, C, 0)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[1], z[0], t[0])
	C, t[1] = madd2(x[1], z[1], t[1], C)
	C, t[2] = madd2(x[1], z[2], t[2], C)
	C, t[3] = madd2(x[1], z[3], t[3], C)
	C, t[4] = madd2(x[1], z[4], t[4], C)
	C, t[5] = madd2(x[1], z[5], t[5], C)
	C, t[6] = madd2(x[1], z[6], t[6], C)
	C, t[7] = madd2(x[1], z[7], t[7], C)
	C, t[8] = madd2(x[1], z[8], t[8], C)

	D = C

	// m = t[0]n'[0] mod W
	m = t[0] * 13715081754621093557

	// -----------------------------------
	// Second loop
	C = madd0(m, 1897235049931565155, t[0])

	C, t[0] = madd2(m, 13948212332252695022, t[1], C)

	C, t[1] = madd2(m, 5931880109645070248, t[2], C)

	C, t[2] = madd2(m, 13767938734030060565, t[3], C)

	C, t[3] = madd2(m, 1578249351545375860, t[4], C)

	C, t[4] = madd2(m, 13837360922177919458, t[5], C)

	C, t[5] = madd2(m, 13195475758251816711, t[6], C)

	C, t[6] = madd2(m, 18260081051345282122, t[7], C)

	C, t[7] = madd3(m, 7695857179358191508, t[8], C, t[9])

	t[8], t[9] = bits.Add64(D, C, 0)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[2], z[0], t[0])
	C, t[1] = madd2(x[2], z[1], t[1], C)
	C, t[2] = madd2(x[2], z[2], t[2], C)
	C, t[3] = madd2(x[2], z[3], t[3], C)
	C, t[4] = madd2(x[2], z[4], t[4], C)
	C, t[5] = madd2(x[2], z[5], t[5], C)
	C, t[6] = madd2(x[2], z[6], t[6], C)
	C, t[7] = madd2(x[2], z[7], t[7], C)
	C, t[8] = madd2(x[2], z[8], t[8], C)

	D = C

	// m = t[0]n'[0] mod W
	m = t[0] * 13715081754621093557

	// -----------------------------------
	// Second loop
	C = madd0(m, 1897235049931565155, t[0])

	C, t[0] = madd2(m, 13948212332252695022, t[1], C)

	C, t[1] = madd2(m, 5931880109645070248, t[2], C)

	C, t[2] = madd2(m, 13767938734030060565, t[3], C)

	C, t[3] = madd2(m, 1578249351545375860, t[4], C)

	C, t[4] = madd2(m, 13837360922177919458, t[5], C)

	C, t[5] = madd2(m, 13195475758251816711, t[6], C)

	C, t[6] = madd2(m, 18260081051345282122, t[7], C)

	C, t[7] = madd3(m, 7695857179358191508, t[8], C, t[9])

	t[8], t[9] = bits.Add64(D, C, 0)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[3], z[0], t[0])
	C, t[1] = madd2(x[3], z[1], t[1], C)
	C, t[2] = madd2(x[3], z[2], t[2], C)
	C, t[3] = madd2(x[3], z[3], t[3], C)
	C, t[4] = madd2(x[3], z[4], t[4], C)
	C, t[5] = madd2(x[3], z[5], t[5], C)
	C, t[6] = madd2(x[3], z[6], t[6], C)
	C, t[7] = madd2(x[3], z[7], t[7], C)
	C, t[8] = madd2(x[3], z[8], t[8], C)

	D = C

	// m = t[0]n'[0] mod W
	m = t[0] * 13715081754621093557

	// -----------------------------------
	// Second loop
	C = madd0(m, 1897235049931565155, t[0])

	C, t[0] = madd2(m, 13948212332252695022, t[1], C)

	C, t[1] = madd2(m, 5931880109645070248, t[2], C)

	C, t[2] = madd2(m, 13767938734030060565, t[3], C)

	C, t[3] = madd2(m, 1578249351545375860, t[4], C)

	C, t[4] = madd2(m, 13837360922177919458, t[5], C)

	C, t[5] = madd2(m, 13195475758251816711, t[6], C)

	C, t[6] = madd2(m, 18260081051345282122, t[7], C)

	C, t[7] = madd3(m, 7695857179358191508, t[8], C, t[9])

	t[8], t[9] = bits.Add64(D, C, 0)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[4], z[0], t[0])
	C, t[1] = madd2(x[4], z[1], t[1], C)
	C, t[2] = madd2(x[4], z[2], t[2], C)
	C, t[3] = madd2(x[4], z[3], t[3], C)
	C, t[4] = madd2(x[4], z[4], t[4], C)
	C, t[5] = madd2(x[4], z[5], t[5], C)
	C, t[6] = madd2(x[4], z[6], t[6], C)
	C, t[7] = madd2(x[4], z[7], t[7], C)
	C, t[8] = madd2(x[4], z[8], t[8], C)

	D = C

	// m = t[0]n'[0] mod W
	m = t[0] * 13715081754621093557

	// -----------------------------------
	// Second loop
	C = madd0(m, 1897235049931565155, t[0])

	C, t[0] = madd2(m, 13948212332252695022, t[1], C)

	C, t[1] = madd2(m, 5931880109645070248, t[2], C)

	C, t[2] = madd2(m, 13767938734030060565, t[3], C)

	C, t[3] = madd2(m, 1578249351545375860, t[4], C)

	C, t[4] = madd2(m, 13837360922177919458, t[5], C)

	C, t[5] = madd2(m, 13195475758251816711, t[6], C)

	C, t[6] = madd2(m, 18260081051345282122, t[7], C)

	C, t[7] = madd3(m, 7695857179358191508, t[8], C, t[9])

	t[8], t[9] = bits.Add64(D, C, 0)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[5], z[0], t[0])
	C, t[1] = madd2(x[5], z[1], t[1], C)
	C, t[2] = madd2(x[5], z[2], t[2], C)
	C, t[3] = madd2(x[5], z[3], t[3], C)
	C, t[4] = madd2(x[5], z[4], t[4], C)
	C, t[5] = madd2(x[5], z[5], t[5], C)
	C, t[6] = madd2(x[5], z[6], t[6], C)
	C, t[7] = madd2(x[5], z[7], t[7], C)
	C, t[8] = madd2(x[5], z[8], t[8], C)

	D = C

	// m = t[0]n'[0] mod W
	m = t[0] * 13715081754621093557

	// -----------------------------------
	// Second loop
	C = madd0(m, 1897235049931565155, t[0])

	C, t[0] = madd2(m, 13948212332252695022, t[1], C)

	C, t[1] = madd2(m, 5931880109645070248, t[2], C)

	C, t[2] = madd2(m, 13767938734030060565, t[3], C)

	C, t[3] = madd2(m, 1578249351545375860, t[4], C)

	C, t[4] = madd2(m, 13837360922177919458, t[5], C)

	C, t[5] = madd2(m, 13195475758251816711, t[6], C)

	C, t[6] = madd2(m, 18260081051345282122, t[7], C)

	C, t[7] = madd3(m, 7695857179358191508, t[8], C, t[9])

	t[8], t[9] = bits.Add64(D, C, 0)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[6], z[0], t[0])
	C, t[1] = madd2(x[6], z[1], t[1], C)
	C, t[2] = madd2(x[6], z[2], t[2], C)
	C, t[3] = madd2(x[6], z[3], t[3], C)
	C, t[4] = madd2(x[6], z[4], t[4], C)
	C, t[5] = madd2(x[6], z[5], t[5], C)
	C, t[6] = madd2(x[6], z[6], t[6], C)
	C, t[7] = madd2(x[6], z[7], t[7], C)
	C, t[8] = madd2(x[6], z[8], t[8], C)

	D = C

	// m = t[0]n'[0] mod W
	m = t[0] * 13715081754621093557

	// -----------------------------------
	// Second loop
	C = madd0(m, 1897235049931565155, t[0])

	C, t[0] = madd2(m, 13948212332252695022, t[1], C)

	C, t[1] = madd2(m, 5931880109645070248, t[2], C)

	C, t[2] = madd2(m, 13767938734030060565, t[3], C)

	C, t[3] = madd2(m, 1578249351545375860, t[4], C)

	C, t[4] = madd2(m, 13837360922177919458, t[5], C)

	C, t[5] = madd2(m, 13195475758251816711, t[6], C)

	C, t[6] = madd2(m, 18260081051345282122, t[7], C)

	C, t[7] = madd3(m, 7695857179358191508, t[8], C, t[9])

	t[8], t[9] = bits.Add64(D, C, 0)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[7], z[0], t[0])
	C, t[1] = madd2(x[7], z[1], t[1], C)
	C, t[2] = madd2(x[7], z[2], t[2], C)
	C, t[3] = madd2(x[7], z[3], t[3], C)
	C, t[4] = madd2(x[7], z[4], t[4], C)
	C, t[5] = madd2(x[7], z[5], t[5], C)
	C, t[6] = madd2(x[7], z[6], t[6], C)
	C, t[7] = madd2(x[7], z[7], t[7], C)
	C, t[8] = madd2(x[7], z[8], t[8], C)

	D = C

	// m = t[0]n'[0] mod W
	m = t[0] * 13715081754621093557

	// -----------------------------------
	// Second loop
	C = madd0(m, 1897235049931565155, t[0])

	C, t[0] = madd2(m, 13948212332252695022, t[1], C)

	C, t[1] = madd2(m, 5931880109645070248, t[2], C)

	C, t[2] = madd2(m, 13767938734030060565, t[3], C)

	C, t[3] = madd2(m, 1578249351545375860, t[4], C)

	C, t[4] = madd2(m, 13837360922177919458, t[5], C)

	C, t[5] = madd2(m, 13195475758251816711, t[6], C)

	C, t[6] = madd2(m, 18260081051345282122, t[7], C)

	C, t[7] = madd3(m, 7695857179358191508, t[8], C, t[9])

	t[8], t[9] = bits.Add64(D, C, 0)
	// -----------------------------------
	// First loop

	C, t[0] = madd1(x[8], z[0], t[0])
	C, t[1] = madd2(x[8], z[1], t[1], C)
	C, t[2] = madd2(x[8], z[2], t[2], C)
	C, t[3] = madd2(x[8], z[3], t[3], C)
	C, t[4] = madd2(x[8], z[4], t[4], C)
	C, t[5] = madd2(x[8], z[5], t[5], C)
	C, t[6] = madd2(x[8], z[6], t[6], C)
	C, t[7] = madd2(x[8], z[7], t[7], C)
	C, t[8] = madd2(x[8], z[8], t[8], C)

	D = C

	// m = t[0]n'[0] mod W
	m = t[0] * 13715081754621093557

	// -----------------------------------
	// Second loop
	C = madd0(m, 1897235049931565155, t[0])

	C, t[0] = madd2(m, 13948212332252695022, t[1], C)

	C, t[1] = madd2(m, 5931880109645070248, t[2], C)

	C, t[2] = madd2(m, 13767938734030060565, t[3], C)

	C, t[3] = madd2(m, 1578249351545375860, t[4], C)

	C, t[4] = madd2(m, 13837360922177919458, t[5], C)

	C, t[5] = madd2(m, 13195475758251816711, t[6], C)

	C, t[6] = madd2(m, 18260081051345282122, t[7], C)

	C, t[7] = madd3(m, 7695857179358191508, t[8], C, t[9])

	t[8], t[9] = bits.Add64(D, C, 0)

	if t[9] != 0 {
		// we need to reduce, we have a result on 10 words
		var b uint64
		z[0], b = bits.Sub64(t[0], 1897235049931565155, 0)
		z[1], b = bits.Sub64(t[1], 13948212332252695022, b)
		z[2], b = bits.Sub64(t[2], 5931880109645070248, b)
		z[3], b = bits.Sub64(t[3], 13767938734030060565, b)
		z[4], b = bits.Sub64(t[4], 1578249351545375860, b)
		z[5], b = bits.Sub64(t[5], 13837360922177919458, b)
		z[6], b = bits.Sub64(t[6], 13195475758251816711, b)
		z[7], b = bits.Sub64(t[7], 18260081051345282122, b)
		z[8], _ = bits.Sub64(t[8], 7695857179358191508, b)
		return z
	}

	// copy t into z
	z[0] = t[0]
	z[1] = t[1]
	z[2] = t[2]
	z[3] = t[3]
	z[4] = t[4]
	z[5] = t[5]
	z[6] = t[6]
	z[7] = t[7]
	z[8] = t[8]

	// if z > q --> z -= q
	if !(z[8] < 7695857179358191508 || (z[8] == 7695857179358191508 && (z[7] < 18260081051345282122 || (z[7] == 18260081051345282122 && (z[6] < 13195475758251816711 || (z[6] == 13195475758251816711 && (z[5] < 13837360922177919458 || (z[5] == 13837360922177919458 && (z[4] < 1578249351545375860 || (z[4] == 1578249351545375860 && (z[3] < 13767938734030060565 || (z[3] == 13767938734030060565 && (z[2] < 5931880109645070248 || (z[2] == 5931880109645070248 && (z[1] < 13948212332252695022 || (z[1] == 13948212332252695022 && (z[0] < 1897235049931565155))))))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 1897235049931565155, 0)
		z[1], b = bits.Sub64(z[1], 13948212332252695022, b)
		z[2], b = bits.Sub64(z[2], 5931880109645070248, b)
		z[3], b = bits.Sub64(z[3], 13767938734030060565, b)
		z[4], b = bits.Sub64(z[4], 1578249351545375860, b)
		z[5], b = bits.Sub64(z[5], 13837360922177919458, b)
		z[6], b = bits.Sub64(z[6], 13195475758251816711, b)
		z[7], b = bits.Sub64(z[7], 18260081051345282122, b)
		z[8], _ = bits.Sub64(z[8], 7695857179358191508, b)
	}
	return z
}

func (z *Element09) mulNoCarry(x *Element09) *Element09 {

	var t [9]uint64
	var c [3]uint64
	{
		// round 0
		v := z[0]
		c[1], c[0] = bits.Mul64(v, x[0])
		m := c[0] * 13715081754621093557
		c[2] = madd0(m, 1897235049931565155, c[0])
		c[1], c[0] = madd1(v, x[1], c[1])
		c[2], t[0] = madd2(m, 13948212332252695022, c[2], c[0])
		c[1], c[0] = madd1(v, x[2], c[1])
		c[2], t[1] = madd2(m, 5931880109645070248, c[2], c[0])
		c[1], c[0] = madd1(v, x[3], c[1])
		c[2], t[2] = madd2(m, 13767938734030060565, c[2], c[0])
		c[1], c[0] = madd1(v, x[4], c[1])
		c[2], t[3] = madd2(m, 1578249351545375860, c[2], c[0])
		c[1], c[0] = madd1(v, x[5], c[1])
		c[2], t[4] = madd2(m, 13837360922177919458, c[2], c[0])
		c[1], c[0] = madd1(v, x[6], c[1])
		c[2], t[5] = madd2(m, 13195475758251816711, c[2], c[0])
		c[1], c[0] = madd1(v, x[7], c[1])
		c[2], t[6] = madd2(m, 18260081051345282122, c[2], c[0])
		c[1], c[0] = madd1(v, x[8], c[1])
		t[8], t[7] = madd3(m, 7695857179358191508, c[0], c[2], c[1])
	}
	{
		// round 1
		v := z[1]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 13715081754621093557
		c[2] = madd0(m, 1897235049931565155, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13948212332252695022, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 5931880109645070248, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 13767938734030060565, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 1578249351545375860, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 13837360922177919458, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		c[2], t[5] = madd2(m, 13195475758251816711, c[2], c[0])
		c[1], c[0] = madd2(v, x[7], c[1], t[7])
		c[2], t[6] = madd2(m, 18260081051345282122, c[2], c[0])
		c[1], c[0] = madd2(v, x[8], c[1], t[8])
		t[8], t[7] = madd3(m, 7695857179358191508, c[0], c[2], c[1])
	}
	{
		// round 2
		v := z[2]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 13715081754621093557
		c[2] = madd0(m, 1897235049931565155, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13948212332252695022, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 5931880109645070248, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 13767938734030060565, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 1578249351545375860, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 13837360922177919458, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		c[2], t[5] = madd2(m, 13195475758251816711, c[2], c[0])
		c[1], c[0] = madd2(v, x[7], c[1], t[7])
		c[2], t[6] = madd2(m, 18260081051345282122, c[2], c[0])
		c[1], c[0] = madd2(v, x[8], c[1], t[8])
		t[8], t[7] = madd3(m, 7695857179358191508, c[0], c[2], c[1])
	}
	{
		// round 3
		v := z[3]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 13715081754621093557
		c[2] = madd0(m, 1897235049931565155, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13948212332252695022, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 5931880109645070248, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 13767938734030060565, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 1578249351545375860, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 13837360922177919458, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		c[2], t[5] = madd2(m, 13195475758251816711, c[2], c[0])
		c[1], c[0] = madd2(v, x[7], c[1], t[7])
		c[2], t[6] = madd2(m, 18260081051345282122, c[2], c[0])
		c[1], c[0] = madd2(v, x[8], c[1], t[8])
		t[8], t[7] = madd3(m, 7695857179358191508, c[0], c[2], c[1])
	}
	{
		// round 4
		v := z[4]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 13715081754621093557
		c[2] = madd0(m, 1897235049931565155, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13948212332252695022, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 5931880109645070248, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 13767938734030060565, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 1578249351545375860, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 13837360922177919458, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		c[2], t[5] = madd2(m, 13195475758251816711, c[2], c[0])
		c[1], c[0] = madd2(v, x[7], c[1], t[7])
		c[2], t[6] = madd2(m, 18260081051345282122, c[2], c[0])
		c[1], c[0] = madd2(v, x[8], c[1], t[8])
		t[8], t[7] = madd3(m, 7695857179358191508, c[0], c[2], c[1])
	}
	{
		// round 5
		v := z[5]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 13715081754621093557
		c[2] = madd0(m, 1897235049931565155, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13948212332252695022, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 5931880109645070248, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 13767938734030060565, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 1578249351545375860, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 13837360922177919458, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		c[2], t[5] = madd2(m, 13195475758251816711, c[2], c[0])
		c[1], c[0] = madd2(v, x[7], c[1], t[7])
		c[2], t[6] = madd2(m, 18260081051345282122, c[2], c[0])
		c[1], c[0] = madd2(v, x[8], c[1], t[8])
		t[8], t[7] = madd3(m, 7695857179358191508, c[0], c[2], c[1])
	}
	{
		// round 6
		v := z[6]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 13715081754621093557
		c[2] = madd0(m, 1897235049931565155, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13948212332252695022, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 5931880109645070248, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 13767938734030060565, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 1578249351545375860, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 13837360922177919458, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		c[2], t[5] = madd2(m, 13195475758251816711, c[2], c[0])
		c[1], c[0] = madd2(v, x[7], c[1], t[7])
		c[2], t[6] = madd2(m, 18260081051345282122, c[2], c[0])
		c[1], c[0] = madd2(v, x[8], c[1], t[8])
		t[8], t[7] = madd3(m, 7695857179358191508, c[0], c[2], c[1])
	}
	{
		// round 7
		v := z[7]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 13715081754621093557
		c[2] = madd0(m, 1897235049931565155, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13948212332252695022, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 5931880109645070248, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 13767938734030060565, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 1578249351545375860, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 13837360922177919458, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		c[2], t[5] = madd2(m, 13195475758251816711, c[2], c[0])
		c[1], c[0] = madd2(v, x[7], c[1], t[7])
		c[2], t[6] = madd2(m, 18260081051345282122, c[2], c[0])
		c[1], c[0] = madd2(v, x[8], c[1], t[8])
		t[8], t[7] = madd3(m, 7695857179358191508, c[0], c[2], c[1])
	}
	{
		// round 8
		v := z[8]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 13715081754621093557
		c[2] = madd0(m, 1897235049931565155, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], z[0] = madd2(m, 13948212332252695022, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], z[1] = madd2(m, 5931880109645070248, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], z[2] = madd2(m, 13767938734030060565, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], z[3] = madd2(m, 1578249351545375860, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], z[4] = madd2(m, 13837360922177919458, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		c[2], z[5] = madd2(m, 13195475758251816711, c[2], c[0])
		c[1], c[0] = madd2(v, x[7], c[1], t[7])
		c[2], z[6] = madd2(m, 18260081051345282122, c[2], c[0])
		c[1], c[0] = madd2(v, x[8], c[1], t[8])
		z[8], z[7] = madd3(m, 7695857179358191508, c[0], c[2], c[1])
	}

	// if z > q --> z -= q
	if !(z[8] < 7695857179358191508 || (z[8] == 7695857179358191508 && (z[7] < 18260081051345282122 || (z[7] == 18260081051345282122 && (z[6] < 13195475758251816711 || (z[6] == 13195475758251816711 && (z[5] < 13837360922177919458 || (z[5] == 13837360922177919458 && (z[4] < 1578249351545375860 || (z[4] == 1578249351545375860 && (z[3] < 13767938734030060565 || (z[3] == 13767938734030060565 && (z[2] < 5931880109645070248 || (z[2] == 5931880109645070248 && (z[1] < 13948212332252695022 || (z[1] == 13948212332252695022 && (z[0] < 1897235049931565155))))))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 1897235049931565155, 0)
		z[1], b = bits.Sub64(z[1], 13948212332252695022, b)
		z[2], b = bits.Sub64(z[2], 5931880109645070248, b)
		z[3], b = bits.Sub64(z[3], 13767938734030060565, b)
		z[4], b = bits.Sub64(z[4], 1578249351545375860, b)
		z[5], b = bits.Sub64(z[5], 13837360922177919458, b)
		z[6], b = bits.Sub64(z[6], 13195475758251816711, b)
		z[7], b = bits.Sub64(z[7], 18260081051345282122, b)
		z[8], _ = bits.Sub64(z[8], 7695857179358191508, b)
	}
	return z
}

func (z *Element09) mulFIPS(x *Element09) *Element09 {

	var p [9]uint64
	var t, u, v uint64
	u, v = bits.Mul64(z[0], x[0])
	p[0] = v * 13715081754621093557
	u, v, _ = madd(p[0], 1897235049931565155, 0, u, v)
	t, u, v = madd(z[0], x[1], 0, u, v)
	t, u, v = madd(p[0], 13948212332252695022, t, u, v)
	t, u, v = madd(z[1], x[0], t, u, v)
	p[1] = v * 13715081754621093557
	u, v, _ = madd(p[1], 1897235049931565155, t, u, v)
	t, u, v = madd(z[0], x[2], 0, u, v)
	t, u, v = madd(p[0], 5931880109645070248, t, u, v)
	t, u, v = madd(z[1], x[1], t, u, v)
	t, u, v = madd(p[1], 13948212332252695022, t, u, v)
	t, u, v = madd(z[2], x[0], t, u, v)
	p[2] = v * 13715081754621093557
	u, v, _ = madd(p[2], 1897235049931565155, t, u, v)
	t, u, v = madd(z[0], x[3], 0, u, v)
	t, u, v = madd(p[0], 13767938734030060565, t, u, v)
	t, u, v = madd(z[1], x[2], t, u, v)
	t, u, v = madd(p[1], 5931880109645070248, t, u, v)
	t, u, v = madd(z[2], x[1], t, u, v)
	t, u, v = madd(p[2], 13948212332252695022, t, u, v)
	t, u, v = madd(z[3], x[0], t, u, v)
	p[3] = v * 13715081754621093557
	u, v, _ = madd(p[3], 1897235049931565155, t, u, v)
	t, u, v = madd(z[0], x[4], 0, u, v)
	t, u, v = madd(p[0], 1578249351545375860, t, u, v)
	t, u, v = madd(z[1], x[3], t, u, v)
	t, u, v = madd(p[1], 13767938734030060565, t, u, v)
	t, u, v = madd(z[2], x[2], t, u, v)
	t, u, v = madd(p[2], 5931880109645070248, t, u, v)
	t, u, v = madd(z[3], x[1], t, u, v)
	t, u, v = madd(p[3], 13948212332252695022, t, u, v)
	t, u, v = madd(z[4], x[0], t, u, v)
	p[4] = v * 13715081754621093557
	u, v, _ = madd(p[4], 1897235049931565155, t, u, v)
	t, u, v = madd(z[0], x[5], 0, u, v)
	t, u, v = madd(p[0], 13837360922177919458, t, u, v)
	t, u, v = madd(z[1], x[4], t, u, v)
	t, u, v = madd(p[1], 1578249351545375860, t, u, v)
	t, u, v = madd(z[2], x[3], t, u, v)
	t, u, v = madd(p[2], 13767938734030060565, t, u, v)
	t, u, v = madd(z[3], x[2], t, u, v)
	t, u, v = madd(p[3], 5931880109645070248, t, u, v)
	t, u, v = madd(z[4], x[1], t, u, v)
	t, u, v = madd(p[4], 13948212332252695022, t, u, v)
	t, u, v = madd(z[5], x[0], t, u, v)
	p[5] = v * 13715081754621093557
	u, v, _ = madd(p[5], 1897235049931565155, t, u, v)
	t, u, v = madd(z[0], x[6], 0, u, v)
	t, u, v = madd(p[0], 13195475758251816711, t, u, v)
	t, u, v = madd(z[1], x[5], t, u, v)
	t, u, v = madd(p[1], 13837360922177919458, t, u, v)
	t, u, v = madd(z[2], x[4], t, u, v)
	t, u, v = madd(p[2], 1578249351545375860, t, u, v)
	t, u, v = madd(z[3], x[3], t, u, v)
	t, u, v = madd(p[3], 13767938734030060565, t, u, v)
	t, u, v = madd(z[4], x[2], t, u, v)
	t, u, v = madd(p[4], 5931880109645070248, t, u, v)
	t, u, v = madd(z[5], x[1], t, u, v)
	t, u, v = madd(p[5], 13948212332252695022, t, u, v)
	t, u, v = madd(z[6], x[0], t, u, v)
	p[6] = v * 13715081754621093557
	u, v, _ = madd(p[6], 1897235049931565155, t, u, v)
	t, u, v = madd(z[0], x[7], 0, u, v)
	t, u, v = madd(p[0], 18260081051345282122, t, u, v)
	t, u, v = madd(z[1], x[6], t, u, v)
	t, u, v = madd(p[1], 13195475758251816711, t, u, v)
	t, u, v = madd(z[2], x[5], t, u, v)
	t, u, v = madd(p[2], 13837360922177919458, t, u, v)
	t, u, v = madd(z[3], x[4], t, u, v)
	t, u, v = madd(p[3], 1578249351545375860, t, u, v)
	t, u, v = madd(z[4], x[3], t, u, v)
	t, u, v = madd(p[4], 13767938734030060565, t, u, v)
	t, u, v = madd(z[5], x[2], t, u, v)
	t, u, v = madd(p[5], 5931880109645070248, t, u, v)
	t, u, v = madd(z[6], x[1], t, u, v)
	t, u, v = madd(p[6], 13948212332252695022, t, u, v)
	t, u, v = madd(z[7], x[0], t, u, v)
	p[7] = v * 13715081754621093557
	u, v, _ = madd(p[7], 1897235049931565155, t, u, v)
	t, u, v = madd(z[0], x[8], 0, u, v)
	t, u, v = madd(p[0], 7695857179358191508, t, u, v)
	t, u, v = madd(z[1], x[7], t, u, v)
	t, u, v = madd(p[1], 18260081051345282122, t, u, v)
	t, u, v = madd(z[2], x[6], t, u, v)
	t, u, v = madd(p[2], 13195475758251816711, t, u, v)
	t, u, v = madd(z[3], x[5], t, u, v)
	t, u, v = madd(p[3], 13837360922177919458, t, u, v)
	t, u, v = madd(z[4], x[4], t, u, v)
	t, u, v = madd(p[4], 1578249351545375860, t, u, v)
	t, u, v = madd(z[5], x[3], t, u, v)
	t, u, v = madd(p[5], 13767938734030060565, t, u, v)
	t, u, v = madd(z[6], x[2], t, u, v)
	t, u, v = madd(p[6], 5931880109645070248, t, u, v)
	t, u, v = madd(z[7], x[1], t, u, v)
	t, u, v = madd(p[7], 13948212332252695022, t, u, v)
	t, u, v = madd(z[8], x[0], t, u, v)
	p[8] = v * 13715081754621093557
	u, v, _ = madd(p[8], 1897235049931565155, t, u, v)
	t, u, v = madd(z[1], x[8], 0, u, v)
	t, u, v = madd(p[1], 7695857179358191508, t, u, v)
	t, u, v = madd(z[2], x[7], t, u, v)
	t, u, v = madd(p[2], 18260081051345282122, t, u, v)
	t, u, v = madd(z[3], x[6], t, u, v)
	t, u, v = madd(p[3], 13195475758251816711, t, u, v)
	t, u, v = madd(z[4], x[5], t, u, v)
	t, u, v = madd(p[4], 13837360922177919458, t, u, v)
	t, u, v = madd(z[5], x[4], t, u, v)
	t, u, v = madd(p[5], 1578249351545375860, t, u, v)
	t, u, v = madd(z[6], x[3], t, u, v)
	t, u, v = madd(p[6], 13767938734030060565, t, u, v)
	t, u, v = madd(z[7], x[2], t, u, v)
	t, u, v = madd(p[7], 5931880109645070248, t, u, v)
	t, u, v = madd(z[8], x[1], t, u, v)
	u, v, p[0] = madd(p[8], 13948212332252695022, t, u, v)
	t, u, v = madd(z[2], x[8], 0, u, v)
	t, u, v = madd(p[2], 7695857179358191508, t, u, v)
	t, u, v = madd(z[3], x[7], t, u, v)
	t, u, v = madd(p[3], 18260081051345282122, t, u, v)
	t, u, v = madd(z[4], x[6], t, u, v)
	t, u, v = madd(p[4], 13195475758251816711, t, u, v)
	t, u, v = madd(z[5], x[5], t, u, v)
	t, u, v = madd(p[5], 13837360922177919458, t, u, v)
	t, u, v = madd(z[6], x[4], t, u, v)
	t, u, v = madd(p[6], 1578249351545375860, t, u, v)
	t, u, v = madd(z[7], x[3], t, u, v)
	t, u, v = madd(p[7], 13767938734030060565, t, u, v)
	t, u, v = madd(z[8], x[2], t, u, v)
	u, v, p[1] = madd(p[8], 5931880109645070248, t, u, v)
	t, u, v = madd(z[3], x[8], 0, u, v)
	t, u, v = madd(p[3], 7695857179358191508, t, u, v)
	t, u, v = madd(z[4], x[7], t, u, v)
	t, u, v = madd(p[4], 18260081051345282122, t, u, v)
	t, u, v = madd(z[5], x[6], t, u, v)
	t, u, v = madd(p[5], 13195475758251816711, t, u, v)
	t, u, v = madd(z[6], x[5], t, u, v)
	t, u, v = madd(p[6], 13837360922177919458, t, u, v)
	t, u, v = madd(z[7], x[4], t, u, v)
	t, u, v = madd(p[7], 1578249351545375860, t, u, v)
	t, u, v = madd(z[8], x[3], t, u, v)
	u, v, p[2] = madd(p[8], 13767938734030060565, t, u, v)
	t, u, v = madd(z[4], x[8], 0, u, v)
	t, u, v = madd(p[4], 7695857179358191508, t, u, v)
	t, u, v = madd(z[5], x[7], t, u, v)
	t, u, v = madd(p[5], 18260081051345282122, t, u, v)
	t, u, v = madd(z[6], x[6], t, u, v)
	t, u, v = madd(p[6], 13195475758251816711, t, u, v)
	t, u, v = madd(z[7], x[5], t, u, v)
	t, u, v = madd(p[7], 13837360922177919458, t, u, v)
	t, u, v = madd(z[8], x[4], t, u, v)
	u, v, p[3] = madd(p[8], 1578249351545375860, t, u, v)
	t, u, v = madd(z[5], x[8], 0, u, v)
	t, u, v = madd(p[5], 7695857179358191508, t, u, v)
	t, u, v = madd(z[6], x[7], t, u, v)
	t, u, v = madd(p[6], 18260081051345282122, t, u, v)
	t, u, v = madd(z[7], x[6], t, u, v)
	t, u, v = madd(p[7], 13195475758251816711, t, u, v)
	t, u, v = madd(z[8], x[5], t, u, v)
	u, v, p[4] = madd(p[8], 13837360922177919458, t, u, v)
	t, u, v = madd(z[6], x[8], 0, u, v)
	t, u, v = madd(p[6], 7695857179358191508, t, u, v)
	t, u, v = madd(z[7], x[7], t, u, v)
	t, u, v = madd(p[7], 18260081051345282122, t, u, v)
	t, u, v = madd(z[8], x[6], t, u, v)
	u, v, p[5] = madd(p[8], 13195475758251816711, t, u, v)
	t, u, v = madd(z[7], x[8], 0, u, v)
	t, u, v = madd(p[7], 7695857179358191508, t, u, v)
	t, u, v = madd(z[8], x[7], t, u, v)
	u, v, p[6] = madd(p[8], 18260081051345282122, t, u, v)
	t, u, v = madd(z[8], x[8], t, u, v)
	u, v, p[7] = madd(p[8], 7695857179358191508, t, u, v)

	p[8] = v
	z[8] = p[8]
	z[7] = p[7]
	z[6] = p[6]
	z[5] = p[5]
	z[4] = p[4]
	z[3] = p[3]
	z[2] = p[2]
	z[1] = p[1]
	z[0] = p[0]
	// copy(z[:], p[:])

	// if z > q --> z -= q
	if !(z[8] < 7695857179358191508 || (z[8] == 7695857179358191508 && (z[7] < 18260081051345282122 || (z[7] == 18260081051345282122 && (z[6] < 13195475758251816711 || (z[6] == 13195475758251816711 && (z[5] < 13837360922177919458 || (z[5] == 13837360922177919458 && (z[4] < 1578249351545375860 || (z[4] == 1578249351545375860 && (z[3] < 13767938734030060565 || (z[3] == 13767938734030060565 && (z[2] < 5931880109645070248 || (z[2] == 5931880109645070248 && (z[1] < 13948212332252695022 || (z[1] == 13948212332252695022 && (z[0] < 1897235049931565155))))))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 1897235049931565155, 0)
		z[1], b = bits.Sub64(z[1], 13948212332252695022, b)
		z[2], b = bits.Sub64(z[2], 5931880109645070248, b)
		z[3], b = bits.Sub64(z[3], 13767938734030060565, b)
		z[4], b = bits.Sub64(z[4], 1578249351545375860, b)
		z[5], b = bits.Sub64(z[5], 13837360922177919458, b)
		z[6], b = bits.Sub64(z[6], 13195475758251816711, b)
		z[7], b = bits.Sub64(z[7], 18260081051345282122, b)
		z[8], _ = bits.Sub64(z[8], 7695857179358191508, b)
	}
	return z
}

func BenchmarkMulCIOSELEMENT09(b *testing.B) {
	x := Element09{
		13830955647730413669,
		6927004744728180,
		6111362518727439073,
		1139408670260740882,
		3895080857423388830,
		17074512152868828260,
		8192131671142038703,
		16959221168518730559,
		4996576493377651554,
	}
	benchResElement09.SetOne()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.mulCIOS(&x)
	}
}

func BenchmarkMulFIPSELEMENT09(b *testing.B) {
	x := Element09{
		13830955647730413669,
		6927004744728180,
		6111362518727439073,
		1139408670260740882,
		3895080857423388830,
		17074512152868828260,
		8192131671142038703,
		16959221168518730559,
		4996576493377651554,
	}
	benchResElement09.SetOne()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.mulFIPS(&x)
	}
}

func BenchmarkMulNoCarryELEMENT09(b *testing.B) {
	x := Element09{
		13830955647730413669,
		6927004744728180,
		6111362518727439073,
		1139408670260740882,
		3895080857423388830,
		17074512152868828260,
		8192131671142038703,
		16959221168518730559,
		4996576493377651554,
	}
	benchResElement09.SetOne()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResElement09.mulNoCarry(&x)
	}
}
