package inner

import "fmt"

type Inner struct{}

func (i Inner) Print() {
	fmt.Println("Inner")
}
