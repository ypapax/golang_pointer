package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	type tt struct {
		l string
	}
	var a = &tt{l: "a"}
	log.Printf("a: %+v", a)
	var b = &tt{l: "b"}
	log.Printf("b: %+v", b)
	a = b
	log.Printf("a: %+v", a)
	log.Printf("b: %+v", b)
	a.l = "c"
	log.Printf("a: %+v", a)
	log.Printf("b: %+v", b)
	b.l = "d"
	log.Printf("a: %+v", a)
	log.Printf("b: %+v", b)
	//2024/03/06 07:38:04 main.go:12: a: &{l:a}
	//2024/03/06 07:38:04 main.go:14: b: &{l:b}
	//2024/03/06 07:38:04 main.go:16: a: &{l:b}
	//2024/03/06 07:38:04 main.go:17: b: &{l:b}
	//2024/03/06 07:38:04 main.go:19: a: &{l:c}
	//2024/03/06 07:38:04 main.go:20: b: &{l:c}
	//2024/03/06 07:38:04 main.go:22: a: &{l:d}
	//2024/03/06 07:38:04 main.go:23: b: &{l:d}
}
