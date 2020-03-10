package monitor

import (
	"encoding/json"
	"mlvm/vm"
	"net/http"
)

func mem(runners []*vm.Runner, memory *vm.Memoryspace, writer http.ResponseWriter, request *http.Request) {
	status := struct {
		Mem interface{}
		Cpu interface{}
	}{}
	sp := memory.Rect(0, 0, memory.X(), memory.Y())
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
