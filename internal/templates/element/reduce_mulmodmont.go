package element

// TODO this should  be changed to match ASM constant time version
const ReduceMulModMont = `
{{ define "reduce_mulmodmont" }}
// if z > q --> z -= q
// note: this is NOT constant time
if !({{- range $i := reverse $.all.NbWordsIndexesFull}} z[{{$i}}] < {{$.V3}}[{{$i}}] || ( z[{{$i}}] == {{$.V3}}[{{$i}}] && (
{{- end}}z[0] < {{$.V3}}[0] {{- range $i :=  $.all.NbWordsIndexesFull}} )) {{- end}} ){
	var b uint64
	z[0], b = bits.Sub64(z[0], {{$.V3}}[0], 0)
	{{- range $i := $.all.NbWordsIndexesFull}}
		{{-  if eq $i $.all.NbWordsLastIndex}}
			z[{{$i}}], _ = bits.Sub64(z[{{$i}}], {{$.V3}}[{{$i}}], b)
		{{-  else  }}
			z[{{$i}}], b = bits.Sub64(z[{{$i}}], {{$.V3}}[{{$i}}], b)
		{{- end}}
	{{- end}}
}
{{-  end }}

`
