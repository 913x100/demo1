package outer

import "fmt"

type Outer struct {}

func (o Outer) Print() {
	fmt.Println("Outer")
}