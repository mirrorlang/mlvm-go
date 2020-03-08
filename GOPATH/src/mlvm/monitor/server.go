package monitor

import (
	"encoding/json"
	"mlvm/vm"
	"net/http"
	"strconv"
)

var Port int = 8000

func Run(runners []*vm.Runner, memory *vm.Memoryspace) {

	http.HandleFunc("cpu", func(writer http.ResponseWriter, request *http.Request) {
		bs, err := json.Marshal(runners)
		if err != nil {
			panic(err)
		}
		_, err = writer.Write(bs)
		if err != nil {
			panic(err)
		}
	})
	http.HandleFunc("mem", func(writer http.ResponseWriter, request *http.Request) {
		bs, err := json.Marshal(memory.Space)
		if err != nil {
			panic(err)
		}
		_, err = writer.Write(bs)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":"+strconv.Itoa(Port), nil)
}
