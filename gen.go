// genpkg creates a package that is hard for the compiler to compile.
//
// usage:
//
//      mkdir x
//      go run gen.go -n 5000 > x/types.go
//      time go build ./x
package main

import (
	"crypto/sha1"
	"encoding/binary"
	"flag"
	"fmt"
)

func hash(i int) string {
	hash := sha1.New()
	binary.Write(hash, binary.BigEndian, uint32(i))
	return fmt.Sprintf("Type%x", hash.Sum(nil))
}

func main() {
	n := flag.Int("n", 10000, "number of types to generate")
	noinitfn := flag.Bool("noinitfn", false, "don't generate n init fuctions")
	flag.Parse()

	fmt.Println("package x")
	fmt.Println(`import "reflect"`)

	for i := 0; i < *n; i++ {
		fmt.Printf("type %s struct { A int; B string; c []string; D float64 }\n", hash(i))
		if !*noinitfn {
			fmt.Printf("func init() { t[%q] = reflect.TypeOf((*%s)(nil)).Elem() }\n", hash(i), hash(i))
		}
	}

	if !*noinitfn {
		fmt.Println("var t = map[string]reflect.Type{}")
	} else {
		fmt.Println("var t = map[string]reflect.Type{")
		for i := 0; i < *n; i++ {
			fmt.Printf("%q: reflect.TypeOf((*%s)(nil)).Elem(),\n", hash(i), hash(i))
		}
		fmt.Println("}")
	}

}
