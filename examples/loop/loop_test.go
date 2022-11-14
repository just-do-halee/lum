package loop_test

import (
	"testing"

	"github.com/just-do-halee/lum"
)

func Sum(a, b int) int { return a + b }

func TestSum(t *testing.T) {
	type Args struct {
		a, b int
		// for distinguishing test cases
		autoInc bool
	}
	type Ctx = *lum.Context[Args, int]

	lum.Batch[Args, int]{
		{
			Name: "Increment lhs by 1, 100 times",
			Args: Args{1, 1, true},
			Pass: func(c Ctx) {
				// assert
				c.Logf("%v + %v = %v", c.Arguments.a, c.Arguments.b, c.Result)
				c.AssertResultEqual(c.Arguments.a + c.Arguments.b)
			},
			Loop: 100,
		},
		{
			Name: "Increment rhs by 2, 10 times",
			Args: Args{1, 1, false},
			Pass: func(c Ctx) {
				// assert
				c.Logf("%v + %v = %v", c.Arguments.a, c.Arguments.b, c.Result)
				c.AssertResultEqual(c.Arguments.a + c.Arguments.b)
				c.Arguments.b += 2
			},
			Loop: 10,
		},
	}.Run(t, "Sum", func(a Args) int {
		// before each
		return Sum(a.a, a.b)
	}, func(c Ctx) {
		// after each
		if c.Arguments.autoInc {
			c.Arguments.a++
		}
	})
}
