package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	sl1, sl2 := make([]int, cap(ch1)), make([]int, cap(ch2))
	
	for i := 0; i < cap(ch1); i++ {
		sl1[i] = <- ch1
		sl2[i] = <- ch2
	}
	
	isSame := false
	
	for _, v1 := range sl1 {
		for _, v2 := range sl2 {
			if v1 == v2 {
				isSame = true
				break
			} else {
				isSame = false
			}
		}
		
		if isSame == false {
			break	
		}
	}
		
	return isSame
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	
	for ii := 0; ii < cap(ch); ii++ {
		fmt.Println(<- ch)		
	}
	
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(3), tree.New(32)))
}
