package element

const Base = `

// /!\ WARNING /!\
// this code has not been audited and is provided as-is. In particular, 
// there is no security guarantees such as constant time implementation 
// or side-channel attack resistance
// /!\ WARNING /!\

import (
	"math/bits"
)

// {{.ElementName}} represents a field element stored on {{.NbWords}} words (uint64)
// {{.ElementName}} are assumed to be in Montgomery form in all methods
// field modulus q =
// 
// {{.Modulus}}
type {{.ElementName}} [{{.NbWords}}]uint64

// Limbs number of 64 bits words needed to represent {{.ElementName}}
const Limbs = {{.NbWords}}

// Bits number bits needed to represent {{.ElementName}}
const Bits = {{.NbBits}}

// Bytes number bytes needed to represent {{.ElementName}}
const Bytes = Limbs * 8

// Mul z = x * y mod q 
// see https://hackmd.io/@zkteam/modular_multiplication
func (z *{{.ElementName}}) Mul(x, y, mod *{{.ElementName}}, modinv uint64) *{{.ElementName}} {
	mul(z, x, y, mod, modinv)
	return z
}

// Add z = x + y mod q
func (z *{{.ElementName}}) Add( x, y, mod *{{.ElementName}}) *{{.ElementName}} {
	add(z, x, y, mod)
	return z 
}

// Sub  z = x - y mod q
func (z *{{.ElementName}}) Sub( x, y, mod *{{.ElementName}}) *{{.ElementName}} {
	sub(z, x, y, mod)
	return z
}

// Generic (no ADX instructions, no AMD64) versions of multiplication and squaring algorithms

func _mulGeneric(z,x,y,mod *{{.ElementName}}, modinv uint64) {
	{{ if .NoCarry}}
		{{ template "mul_nocarry" dict "all" . "V1" "x" "V2" "y" "V3" "mod" "V4" "modinv"}}
	{{ else }}
		{{ template "mul_cios" dict "all" . "V1" "x" "V2" "y" "V3" "mod" "V4" "modinv" "NoReturn" true}}
	{{ end }}
    // TODO can make the following faster and constant time
	{{ template "reduce_mulmodmont" dict "all" . "V3" "mod"}}
}

func _addGeneric(z,  x, y, mod *{{.ElementName}}) {
	var carry uint64
	{{$k := sub $.NbWords 1}}
	z[0], carry = bits.Add64(x[0], y[0], 0)
	{{- range $i := .NbWordsIndexesNoZero}}
		{{- if eq $i $.NbWordsLastIndex}}
		{{- else}}
			z[{{$i}}], carry = bits.Add64(x[{{$i}}], y[{{$i}}], carry)
		{{- end}}
	{{- end}}
	{{- if .NoCarry}}
		z[{{$k}}], _ = bits.Add64(x[{{$k}}], y[{{$k}}], carry)
	{{- else }}
		z[{{$k}}], carry = bits.Add64(x[{{$k}}], y[{{$k}}], carry)
		// if we overflowed the last addition, z >= q
		// if z >= q, z = z - q
		if carry != 0 {
			// we overflowed, so z >= q
			z[0], carry = bits.Sub64(z[0], mod[0], 0)
			{{- range $i := .NbWordsIndexesNoZero}}
				z[{{$i}}], carry = bits.Sub64(z[{{$i}}], mod[{{$i}}], carry)
			{{- end}}
			return 
		}
	{{- end}}

	{{ template "reduce_mulmodmont" dict "all" . "V3" "mod"}}
}

func _subGeneric(z,  x, y, mod *{{.ElementName}}) {
	var b uint64
	z[0], b = bits.Sub64(x[0], y[0], 0)
	{{- range $i := .NbWordsIndexesFull}}
		z[{{$i}}], b = bits.Sub64(x[{{$i}}], y[{{$i}}], b)
	{{- end}}
	if b != 0 {
		var c uint64
		z[0], c = bits.Add64(z[0], mod[0], 0)
		{{- range $i := .NbWordsIndexesFull}}
			{{- if eq $i $.NbWordsLastIndex}}
				z[{{$i}}], _ = bits.Add64(z[{{$i}}], mod[{{$i}}], c)
			{{- else}}
				z[{{$i}}], c = bits.Add64(z[{{$i}}], mod[{{$i}}], c)
			{{- end}}
		{{- end}}
	}
}
`
