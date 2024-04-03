package spec

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	r, _ := Search("w00a1")
	fmt.Println(r)

}

func TestSpecByModel(t *testing.T) {
	fmt.Println(SpecByModel("scc.light.w00a1"))
}
