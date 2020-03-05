package main

import (
	"fmt"
	"mirror"
	"mlvm/vm"
	"time"
)

func main() {
	mem := vm.NewMemory()
	cpu0 := vm.NewRunner()
	vm.TestExpression(mem, mirror.Atom{Point_x: 0, Point_y: 0})
	mem.Print()
	go cpu0.Do(mem, mirror.Atom{Point_x: 0, Point_y: 0})
	fmt.Println()
	time.Sleep(time.Second)
	mem.Print()
}
