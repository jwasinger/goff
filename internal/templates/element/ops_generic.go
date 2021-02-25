package element

const OpsNoAsm = `
// /!\ WARNING /!\
// this code has not been audited and is provided as-is. In particular, 
// there is no security guarantees such as constant time implementation 
// or side-channel attack resistance
// /!\ WARNING /!\


func mul(z, x, y, mod *{{.ElementName}}, modinv uint64) {
	_mulGeneric(z, x, y, mod, modinv)
}

func add(z,  x, y, mod *{{.ElementName}}) {
	_addGeneric(z,x,y, mod)
}

func sub(z,  x, y, mod *{{.ElementName}}) {
	_subGeneric(z,x,y, mod)
}

`
