package vm

import (
	"mlvm_go/mirror"
	"strings"
)

func Code(funcarea [][]mirror.Atom) (r string) {
	for j := 0; j < len(funcarea); j++ {
		for i := 0; i < len(funcarea[0]); i++ {
			atom := funcarea[j][i]
			r += atom.String() + "\t"
		}
		r += "\n"

	}
	return
}
func Load(fstr string) (funcarea [][]mirror.Atom) {

	for _, line := range strings.Split(fstr, "\n") {
		exp := make([]mirror.Atom, 0)
		for _, _ = range strings.Split(line, "\t") {
			//exp = append(exp, mirror.Atom{Type: "string", V_string: atom})
		}
		funcarea = append(funcarea, exp)
	}

	return
}
