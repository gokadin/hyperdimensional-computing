package fhrr

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())

	a := Rand()     // color
	b := Rand()     // black
	x := Bind(a, b) // color x black

	a2 := Rand()       // size
	b2 := Rand()       // small
	x2 := Bind(a2, b2) // size x small

	f := Rand()
	y := Bundle(x, x2, f)

	//for i := 0; i < x.Size(); i++ {
	//	if x.At(i) < -math.Pi || x.At(i) > math.Pi {
	//		fmt.Println(x.At(i))
	//	}
	//}

	c := Unbind(y, a)   // what is color
	c2 := Unbind(y, a2) // what is size

	sim := Similarity(b, c)
	fmt.Println(sim)
	sim2 := Similarity(b2, c2)
	fmt.Println(sim2)
}
