package echo

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

//go test -bench=EchoWithFor -args 123 3 123
func BenchmarkEchoWithFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithFor()
	}
}

func echoWithFor() {
	s := ""
	for _, v := range os.Args[1:] {
		s += v + " "
	}
	fmt.Println(s)
}

//go test -bench=EchoWithJoin -args 123 3 123
func BenchmarkEchoWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithJoin()
	}
}

func echoWithJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
