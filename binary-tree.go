package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

type B []int

func (b B) Contains(a int) bool {
	for _, v := range b {
		if v == a {
			return true
		}
	}
	return false
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var v1, v2 B
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for i := range ch1 {
		v1 = append(v1, i)
	}

	for i := range ch2 {
		v2 = append(v2, i)
	}

	fmt.Println("v", v1)

	for _, v := range v1 {
		result := v2.Contains(v)
		if !result {
			return false
		}

	}
	return true

}

func main() {
	// ch := make(chan int)
	// t := tree.New(1)
	// go func() {
	// 	Walk(t, ch)
	// 	close(ch)
	// }()
	// fmt.Println(t, "original")

	// for i := range ch {
	// 	fmt.Println(i, "i")
	// }

	result := Same(tree.New(1), tree.New(2))
	fmt.Println("result", result)
}
