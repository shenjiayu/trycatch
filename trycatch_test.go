package trycatch

import (
	"fmt"
	"testing"
)

func TestTry(t *testing.T) {
	//case one, it does not panics
	Try(func() {
		fmt.Println("Hello")
	},
		func(e interface{}) {
			fmt.Println("Catch", e)
		},
	)
	//case two, it panics
	Try(func() {
		fmt.Println("Hello")
		panic("But i paniced here")
	},
		func(e interface{}) {
			fmt.Println("Catch: ", e)
		},
	)
}
