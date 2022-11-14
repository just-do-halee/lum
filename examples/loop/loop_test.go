package loop_test

import (
	"testing"

	"github.com/just-do-halee/lum"
)

func Sum(a, b int) int { return a + b }

func TestSum(t *testing.T) {
	type Args struct {
		a, b int
	}
	lum.Batch[Args, int]{
		{
			Name: "Increment lhs by 1, 100 times",
			Args: Args{1, 1},
			Pass: func(c *lum.Ctx[Args, int]) {
				c.Logf("%v + %v = %v", c.Arguments.a, c.Arguments.b, c.Result)
				c.AssertResultEqual(c.Arguments.a + c.Arguments.b)
				c.Arguments.a++
			},
			Loop: 100,
		},
	}.Run(t, "Sum", func(a Args) int {
		return Sum(a.a, a.b)
	})
}
