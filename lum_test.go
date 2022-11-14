package lum

import (
	"testing"
)

func TestBatchRun(t *testing.T) {
	sum := func(a, b int) int { return a + b }
	type Args struct {
		a, b int
	}
	Batch[Args, int]{
		{
			Name: "Sum proper test",
			Args: Args{1, 1},
			Pass: func(c *Ctx[Args, int]) {
				c.Logf("%v + %v = %v", c.Arguments.a, c.Arguments.b, c.Result)
				c.AssertResultEqual(c.Arguments.a + c.Arguments.b)
				c.Arguments.a++
			},
			Loop: 10,
		},
		{
			Name: "It should be",
			Args: Args{1, 3},
			Pass: func(c *Ctx[Args, int]) {
				c.Log(c.Arguments.a, " + ", c.Arguments.b, " = ", c.Result)
				c.Assertf(c.Result > 3, "should be more than %v", 3)
			},
		},
		{
			Name: "it is mock",
			Args: Args{1, 3},
		},
		{
			Name: "It should be",
			Args: Args{1, 3},
			// Pass: Todo[Args, int]("qwdqwd"),
		},
	}.Run(t, "Sum", func(a Args) int {
		return sum(a.a, a.b)
	})
}

func TestFieldRunPanic(t *testing.T) {
	hello := func() string { return "hello" }

	Field[Void, string]{
		Pass: func(c *Ctx[Void, string]) {
			c.AssertNotEqual(c.Result, "hello!")
		},
	}.Run(t, "Hello", func(Void) string {
		return hello()
	})
}

func TestFieldRun(t *testing.T) {
	sum := func(a, b int) int { return a + b }
	type Args struct {
		a, b int
	}
	Field[Args, int]{
		Args: Args{1, 1},
		Pass: func(c *Ctx[Args, int]) {
			c.AssertResultEqual(2, "should be 2 but %v", c.Result)
		},
	}.Run(t, "Sum", func(a Args) int {
		return sum(a.a, a.b)
	})
}
