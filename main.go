package main

// <-- import module yang kita buat dengan go get github.com/koriebruh/say-hello-module
import (
	"fmt"
	"github.com/koriebruh/say-hello-module/v2"
)

// <-- go get github.com/koriebruh/say-hello-module/v2
// jika tidak menambahkan v2 saat import otomatis pake versi 1

func main() {
	test := module_say_hello.SayHello("Jamal")
	fmt.Println(test)

	to := module_say_hello.SayHelloTo("Jamal ", "Eko")
	fmt.Println(to)

}
