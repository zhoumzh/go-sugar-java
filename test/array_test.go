package test

import (
	"fmt"
	"github.com/zhoumzh/go-sugar-java/arrays"
	"testing"
)

func TestContains(t *testing.T) {

	aii := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arrays.Contains(aii, 4))

}
