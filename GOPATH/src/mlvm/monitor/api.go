package monitor

import (
	"encoding/json"
	"mlvm/vm"
	"net/http"
	"strconv"
)

func arg(r *http.Request) (X, Y, size_X, size_Y int) {
	vs := r.URL.Query()
	var err error

	if "" != vs.Get("x") {
		X, err = strconv.Atoi(vs.Get("x"))
		if err != nil {
			panic(err)
		}
	}

	if vs.Get("y") != "" {
		Y, err = strconv.Atoi(vs.Get("y"))
		if err != nil {
			panic(err)
		}
	}

	if "" != vs.Get("size_x") {
		size_X, err = strconv.Atoi(vs.Get("size_x"))
		if err != nil {
			panic(err)
		}
	}
	if vs.Get("size_y") != "" {
		size_Y, err = strconv.Atoi(vs.Get("size_y"))
		if err != nil {
			panic(err)
		}
	}
	return
}
func code(runners []*vm.Runner, memory *vm.Memoryspace, writer http.ResponseWriter, request *http.Request) {
	sp := memory.Rect(arg(request))
	str := vm.Code(sp)
	_, err := writer.Write([]byte(str))
	if err != nil {
		panic(err)
	}
}
func mem(runners []*vm.Runner, memory *vm.Memoryspace, writer http.ResponseWriter, request *http.Request) {
	status := struct {
		Mem interface{}
		Cpu interface{}
	}{}

	sp := memory.Rect(arg(request))
	status.Cpu = runners
	status.Mem = sp
	bs, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}
	_, err = writer.Write(bs)
	if err != nil {
		panic(err)
	}
}
