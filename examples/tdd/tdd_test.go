package tdd_test

import (
	"testing"

	"github.com/just-do-halee/lum"
)

func Sum(a, b int) int { return a + b }

func TestSum(t *testing.T) {
	type Args struct {
		a, b int
	}
	type Ret = int
	type Ctx = *lum.Context[Args, Ret]

	lum.Batch[Args, int]{
		{
			Name: "Increment lhs by 1, 100 times",
			Args: Args{1, 1},
			Pass: lum.Todo[Args, Ret](),
		},
		{
			Name: "Boundary test",
			Args: Args{1, 1},
			Pass: lum.Todo[Args, Ret]("TODO: implement this test"),
		},
		{
			Name: "Fuzzy test",
			Args: Args{},
			Pass: lum.Todof[Args, Ret]("What can i do to test this function? Args.. %v", Args{1, 2}),
		},
	}.Mock().Run(t, "Sum", nil, nil)
}
