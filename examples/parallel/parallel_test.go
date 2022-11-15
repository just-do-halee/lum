package parallel

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

	lum.Batch[Args, Ret]{
		{
			Name: "Increment lhs by 1, 100000 times",
			Args: Args{1, 1},
			Pass: func(c Ctx) {
				c.AssertFatalResultEqual(c.Arguments.a + c.Arguments.b)
				c.Arguments.a++
			},
			Loop:     100000,
			Parallel: true,
		},
		{
			Name: "Increment lhs by 2, 100000 times",
			Args: Args{1, 1},
			Pass: func(c Ctx) {
				c.AssertFatalResultEqual(c.Arguments.a + c.Arguments.b)
				c.Arguments.b += 2
			},
			Loop:     100000,
			Parallel: true,
		},
		{
			Name: "Increment lhs by 3, 100000 times",
			Args: Args{1, 1},
			Pass: func(c Ctx) {
				c.AssertFatalResultEqual(c.Arguments.a + c.Arguments.b)
				c.Arguments.b += 3
			},
			Loop:     100000,
			Parallel: true,
		},
		{
			Name: "Increment lhs by 4, 100000 times",
			Args: Args{1, 1},
			Pass: func(c Ctx) {
				c.AssertFatalResultEqual(c.Arguments.a + c.Arguments.b)
				c.Arguments.b += 4
			},
			Loop:     100000,
			Parallel: true,
		},
		{
			Name: "Increment lhs by 5, 100000 times",
			Args: Args{1, 1},
			Pass: func(c Ctx) {
				c.AssertFatalResultEqual(c.Arguments.a + c.Arguments.b)
				c.Arguments.b += 5
			},
			Loop:     100000,
			Parallel: true,
		},
	}.Run(t, "Sum", func(a Args) Ret {
		return Sum(a.a, a.b)
	}, nil)
}
