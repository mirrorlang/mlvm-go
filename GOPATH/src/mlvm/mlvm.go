package main

import (
	"mirror"
	"mlvm/vm"
	"time"
)

var mem = vm.NewMemory()
var cpu0 = vm.NewRunner()

func test_0() {
	go func() {
		cpu0.Pause()
	}()
	time.Sleep(time.Second)
	vm.TestExpression_null(mem, mirror.Atom{Point_x: 0, Point_y: 0})
	cpu0.Goon()
	time.Sleep(time.Second * 5)
	mem.Print()

}
func test_1() {
	go func() {
		cpu0.Pause()
	}()
	time.Sleep(time.Second)
	vm.TestExpression_not(mem, mirror.Atom{Point_x: 0, Point_y: 0})
	cpu0.Goon()
	time.Sleep(time.Second)
	mem.Print()

}
func main() {
	go cpu0.Do(mem, mirror.Atom{Point_x: 0, Point_y: 0})
	test_0()
}
