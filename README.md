
# **`Lum`**

`Lum` is a simple, expandable, ergonomic test tool in Go (Amazingly small package lol).

[![CI][ci-badge]][ci-url]
[![Licensed][license-badge]][license-url]
[![Twitter][twitter-badge]][twitter-url]

[ci-badge]: https://github.com/just-do-halee/lum/actions/workflows/ci.yml/badge.svg
[license-badge]: https://img.shields.io/github/license/just-do-halee/lum?labelColor=383636
[twitter-badge]: https://img.shields.io/twitter/follow/do_halee?style=flat&logo=twitter&color=4a4646&labelColor=333131&label=just-do-halee
[ci-url]: https://github.com/just-do-halee/lum/actions
[twitter-url]: https://twitter.com/do_halee
[license-url]: https://github.com/just-do-halee/lum

| [Examples](./examples/) | [Latest Note](./CHANGELOG.md) |

```toml
lum = "0.1.0"
```

## **`How to use,`**

```go
package example

func Sum(a, b int) int { return a + b }
```

```go
package example

import (
    "testing"
    "github.com/just-do-halee/lum"
)

func TestSum(t *testing.T) {
	type Arg struct {
		a, b int
	}
	// [Argument, Result Type]
	lum.Batch[Arg, int]{
		{
			Arg: Arg{1, 1},
			Pass: func(c *lum.Ctx[Arg, int]) {
				want := 2
				c.AssertResultEqual(2, "should be %v", want)
			},
		},
		{
			Arg: Arg{1, 3},
			Pass: func(c *lum.Ctx[Arg, int]) {
				c.Log(c.Arguments)
 				c.Log("result: %v", c.Result)
				c.Assert(c.Result > 3, "should be more than 3")
 				c.Assert(c.Result < 5, "should be less than 5")
                
			},
		},
	}.Run(t, "Sum", func(a Arg) int {
 		// Call the actual function
		return Sum(a.a, a.b)
	})
}
```
