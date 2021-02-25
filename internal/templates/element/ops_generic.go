package element

const OpsNoAsm = `
// /!\ WARNING /!\
// this code has not been audited and is provided as-is. In particular, 
// there is no security guarantees such as constant time implementation 
// or side-channel attack resistance
// /!\ WARNING /!\


// MulBy3 x *= 3
func MulBy3(x *{{.ElementName}}) {
	mulByConstant(x, 3)
}

// MulBy5 x *= 5
func MulBy5(x *{{.ElementName}}) {
	mulByConstant(x, 5)
}


func mul(z, x, y, mod *{{.ElementName}}) {
	_mulGeneric(z, x, y, mod)
}

// FromMont converts z in place (i.e. mutates) from Montgomery to regular representation
// sets and returns z = z * 1
func fromMont(z *{{.ElementName}} ) {
	_fromMontGeneric(z)
}

func add(z,  x, y, mod *{{.ElementName}}) {
	_addGeneric(z,x,y, mod)
}

func double(z,  x *{{.ElementName}}) {
	_doubleGeneric(z,x)
}


func sub(z,  x, y, mod *{{.ElementName}}) {
	_subGeneric(z,x,y, mod)
}

func neg(z,  x *{{.ElementName}}) {
	_negGeneric(z,x)
}


func reduce(z *{{.ElementName}})  {
	_reduceGeneric(z)
}


`
