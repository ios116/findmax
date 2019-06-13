package findmax

import (
	"testing"
)

func TestFindMax(t *testing.T) {
	sl:=[]int{1,2,3,4,5,6,7}
	FindMax(sl, func(i,j int) bool{
		return true
	})
}