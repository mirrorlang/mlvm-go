package main

import (
	"github.com/beevik/etree"
	"io/ioutil"
	"os"
)

type Atom interface {
	Type() string
}

func Loadapp() {
	var ml_home = os.Getenv("ML_HOME")
	fs, err := ioutil.ReadDir(ml_home)
	if err != nil {
		panic(err)
	}
	for _, f := range fs {
		f.IsDir()
	}
}
func a(apppath string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(apppath); err != nil {
		panic(err)
	}

}
func main() {

	Loadapp()
}
