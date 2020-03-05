package main

import (
	"fmt"
	"mirror"
	"mlvm/vm"
	"time"
)

var mem = vm.NewMemory()
var cpu0 = vm.NewRunner(mem)

func test_0() {
	go func() {
		cpu0.Pause()
	}()
	time.Sleep(time.Second)
	vm.TestExpression_null(mem, mirror.Atom{Point_x: 0, Point_y: 0})
	vm.TestExpression_not(mem, mirror.Atom{Point_x: 0, Point_y: 2})
	vm.TestExpression_goto(mem, mirror.Atom{Point_x: 0, Point_y: 3})
	mem.Print()
	fmt.Println()
	cpu0.Goon()
	time.Sleep(time.Second * 5)
	mem.Print()

}

func main() {
	go cpu0.Do(mirror.Atom{Point_x: 0, Point_y: 0})
	test_0()
}
