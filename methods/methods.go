package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	v := value{"number"}
	v.pointerReceiver("1")
	fmt.Println(v.name) // 1
	v.valueReceiver("2")
	fmt.Println(v.name) // 1

	p := &v
	p.pointerReceiver("3")
	fmt.Println(p.name) // 3
	p.valueReceiver("4")
	fmt.Println(p.name) // 3
}

type value struct{ name string }

func (s *value) pointerReceiver(name string) {
	s.name = name
}

func (s value) valueReceiver(name string) {
	s.name = name
}
