package lum

import (
	"fmt"
	"testing"
)

type Argumenter interface {
	any
}
type Resulter interface {
	comparable
}

type Void struct{}

type Ctx[Arg Argumenter, Res Resulter] struct {
	t         *testing.T
	FuncName  string
	Arguments Arg
	Result    Res
	reason    string
	logs      string
}

func (c *Ctx[Arg, Res]) String() string {
	return fmt.Sprintf("\n%s\n%s(%v)  -->  %v  \t%s\t", c.logs, c.FuncName, c.Arguments, c.Result, c.reason)
}

func (c *Ctx[Arg, Res]) Log(a ...any) {
	newLine := "\n"
	switch len(a) {
	case 0:
		c.logs += newLine
	case 1:
		c.logs += fmt.Sprint(a[0], newLine)
	default:
		format := fmt.Sprint(a[0])
		c.logs += fmt.Sprintf(format+newLine, a[1:]...)
	}
}

func (c *Ctx[Arg, Res]) ResetLogs() {
	c.logs = ""
}

func (c *Ctx[Arg, Res]) Fail(a ...any) {
	newLine := "\n"
	switch len(a) {
	case 0:
		c.t.Error(c.String(), newLine)
	case 1:
		c.t.Error(c.String(), a[0], newLine)
	default:
		format := fmt.Sprint(a[0])
		c.t.Error(c.String(), fmt.Sprintf(format, a[1:]...), newLine)
	}
}

func (c *Ctx[Arg, Res]) AssertResultEqual(a Res, format ...any) {
	if c.Result != a {
		c.reason = fmt.Sprintf("%v == %v [fail]", c.Result, a)
		c.Fail(format...)
	}
}

func (c *Ctx[Arg, Res]) AssertEqual(a, b Res, format ...any) {
	if a != b {
		c.reason = fmt.Sprintf("%v == %v [fail]", a, b)
		c.Fail(format...)
	}
}

func (c *Ctx[Arg, Res]) AssertResultNotEqual(a Res, format ...any) {
	if c.Result == a {
		c.reason = fmt.Sprintf("%v == %v [fail]", c.Result, a)
		c.Fail(format...)
	}
}

func (c *Ctx[Arg, Res]) AssertNotEqual(a, b Res, format ...any) {
	if a == b {
		c.reason = fmt.Sprintf("%v != %v [fail]", a, b)
		c.Fail(format...)
	}
}

func (c *Ctx[Arg, Res]) Assert(b bool, a ...any) {
	if !b {
		c.reason = "<assertion>"
		c.Fail(a...)
	}
}

type Fn[Arg Argumenter, Res Resulter] func(Arg) Res

type FnPass[Arg Argumenter, Res Resulter] func(*Ctx[Arg, Res])

type Field[Arg Argumenter, Res Resulter] struct {
	Name string
	Args Arg
	Pass FnPass[Arg, Res]
}

func (f Field[Arg, Res]) Run(t *testing.T, fnName string, fn Fn[Arg, Res]) {
	t.Run(f.Name, func(t *testing.T) {
		ctx := &Ctx[Arg, Res]{
			t:         t,
			FuncName:  fnName,
			Arguments: f.Args,
			Result:    fn(f.Args),
		}
		f.Pass(ctx)
	})
}

type Batch[A Argumenter, R Resulter] []Field[A, R]

func (b Batch[A, R]) Run(t *testing.T, fnName string, fn func(a A) R) {
	for _, test := range b {
		test.Run(t, fnName, fn)
	}
}
