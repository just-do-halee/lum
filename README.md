
# **`Lum`**

`Lum` is a simple, ergonomic test tool in Go (Amazingly small package).

[![Go Reference](https://pkg.go.dev/badge/github.com/just-do-halee/lum.svg)](https://pkg.go.dev/github.com/just-do-halee/lum)
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

```shell
go get -u github.com/just-do-halee/lum@latest
```

## **`Template`**

```go
import (
    "testing"

    "github.com/just-do-halee/lum"
)

func TestFn(t *testing.T) {
    type Args struct {
        a, b int
    }
    type Ret = int
    type Ctx = *lum.Context[Args, Ret]
    lum.Batch[Args, Ret]{
        {
            Name: "description",
            Args: Args{0, 0},
            Pass: lum.Todo[Args, Ret]()
        },
    }.Run(t, "Fn", nil, nil)
}
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
    // ... Before All ...

    // Test Function Arguments
    type Args struct {
        a, b int
    }
    // Return Type
    type Ret = int 
    // Context Alias
    type Ctx = *lum.Context[Args, Ret]

    lum.Batch[Args, Ret]{
        {
            Name: "1 + 1 = 2",
            Args: Args{1, 1},
            Pass: func(c Ctx) {
                // ... Assert ...
                c.AssertResultEqual(2)
            },
        },
        {
            Name: "3 > 1 + 3 < 5",
            Args: Args{1, 3},
            Pass: func(c Ctx) {
                // ... Assert ...
                c.Log(c.Arguments)
                c.Logf("result: %v", c.Result)

                c.Assert(c.Result > 3, "should be more than 3")
                c.Assertf(c.Result < 5, "should be less than %v", 5)
            },
        },
        {
            Name: "3 + 5 ...?",
            Args: Args{3, 5},
            // This is unimplemented, so it won't occur any error
        },
        {
            Name: "2 + 4 != 7",
            Args: Args{2, 4},
            // This is todo, so it will occur an error has meta message
            Pass: lum.Todo[Args, Ret]("not equal testing"),
        },
    }.Run(t, "Sum", func(a Args) Ret {
        // ... Before Each ...

        // Call The Actual Function
        return Sum(a.a, a.b)

    }, func(c Ctx) {
        // ... After Each ...
    })
    
    // ... After All ...
}
```
