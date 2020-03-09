package monitor

import (
	"io/ioutil"
	"mlvm/vm"
	"net/http"
	"os"
	"strconv"
)

var Port int = 8000

//func upper(ws *websocket.Conn) {
//	var err error
//	for {
//		var reply string
//
//		if err = websocket.Message.Receive(ws, &reply); err != nil {
//			fmt.Println(err)
//			continue
//		}
//
//		if err = websocket.Message.Send(ws, strings.ToUpper(reply)); err != nil {
//			fmt.Println(err)
//			continue
//		}
//	}
//}
func Run(runners []*vm.Runner, memory *vm.Memoryspace) {
	webpath := os.Args[1]
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		st, err := os.Stat((webpath + request.URL.Path))
		if err == nil {
			if st.IsDir() {
				return
			}
			bs, err := ioutil.ReadFile(webpath + request.URL.Path)
			if err != nil {
				panic(err)
			}
			writer.Write(bs)
			return
		}
		switch request.URL.Path[1:] {
		case "mem":
			mem(runners, memory, writer, request)
		}
	})

	http.ListenAndServe(":"+strconv.Itoa(Port), nil)
}
