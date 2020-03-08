package log

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var (
	//showtracefile 是否显示代码
	showtime      bool = false
	showtracefile bool = false
)

func SetTime(show bool) {
	showtime = show

}
func SetCodefile(show bool) {
	showtracefile = show
}

func Print(c Colortext, args ...interface{}) {
	s := string(c) + fmt.Sprint(args...) + string(EndColor)
	fmt.Print(s)
}
func Println(c Colortext, args ...interface{}) {
	s := string(c) + fmt.Sprint(args...) + string(EndColor)
	printtrace()
	fmt.Println(s)
}
func Printtime() {
	fmt.Print(time.Now().String()[:19])
}
func printtrace() {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		//f := runtime.FuncForPC(pc)
		t := file + ":" + strconv.Itoa(line)
		fmt.Print(t)
	}
}
