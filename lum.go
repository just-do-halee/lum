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
	logs      Stringify
}

func (c *Ctx[Arg, Res]) setReason(a ...any) {
	c.reason = fmt.Sprint(a...)
}

func (c *Ctx[Arg, Res]) setReasonf(format string, a ...any) {
	c.reason = fmt.Sprintf(format, a...)
}

func (c *Ctx[Arg, Res]) String() string {
	return fmt.Sprintf("\n%s\n%s(%v)  -->  %v  \t%s\t", c.logs.String(), c.FuncName, c.Arguments, c.Result, c.reason)
}

func (c *Ctx[Arg, Res]) Log(a ...any) {
	c.logs.WriteStrings(a...)
	c.logs.Writeln()
}

func (c *Ctx[Arg, Res]) Logf(format string, a ...any) {
	c.logs.WriteString(fmt.Sprintf(format+"\n", a...))
}

func (c *Ctx[Arg, Res]) ResetLogs() {
	c.logs.Reset()
}

func (c *Ctx[Arg, Res]) Fail(a ...any) {
	var sb Stringify
	sb.WriteString(c.String())
	sb.WriteStrings(a...)
	sb.Writeln()
	c.t.Error(sb.String())
}

func (c *Ctx[Arg, Res]) Failf(format string, a ...any) {
	c.t.Error(c.String(), fmt.Sprintf(format, a...), "\n")
}

func (c *Ctx[Arg, Res]) AssertResultEqual(rhs Res, a ...any) {
	if c.Result != rhs {
		c.setReasonf("%v == %v [fail]", c.Result, rhs)
		c.Fail(a...)
	}
}

func (c *Ctx[Arg, Res]) AssertResultEqualf(rhs Res, format string, a ...any) {
	if c.Result != rhs {
		c.setReasonf("%v == %v [fail]", c.Result, rhs)
		c.Failf(format, a...)
	}
}

func (c *Ctx[Arg, Res]) AssertEqual(lhs, rhs Res, a ...any) {
	if lhs != rhs {
		c.setReasonf("%v == %v [fail]", lhs, rhs)
		c.Fail(a...)
	}
}

func (c *Ctx[Arg, Res]) AssertEqualf(lhs, rhs Res, format string, a ...any) {
	if lhs != rhs {
		c.setReasonf("%v == %v [fail]", lhs, rhs)
		c.Failf(format, a...)
	}
}

func (c *Ctx[Arg, Res]) AssertResultNotEqual(rhs Res, a ...any) {
	if c.Result == rhs {
		c.setReasonf("%v == %v [fail]", c.Result, rhs)
		c.Fail(a...)
	}
}

func (c *Ctx[Arg, Res]) AssertResultNotEqualf(rhs Res, format string, a ...any) {
	if c.Result == rhs {
		c.setReasonf("%v == %v [fail]", c.Result, rhs)
		c.Failf(format, a...)
	}
}

func (c *Ctx[Arg, Res]) AssertNotEqual(lhs, rhs Res, a ...any) {
	if lhs == rhs {
		c.setReasonf("%v != %v [fail]", lhs, rhs)
		c.Fail(a...)
	}
}

func (c *Ctx[Arg, Res]) AssertNotEqualf(lhs, rhs Res, format string, a ...any) {
	if lhs == rhs {
		c.setReasonf("%v != %v [fail]", lhs, rhs)
		c.Failf(format, a...)
	}
}

func (c *Ctx[Arg, Res]) Assert(b bool, a ...any) {
	if !b {
		c.setReason("<assertion>")
		c.Fail(a...)
	}
}

func (c *Ctx[Arg, Res]) Assertf(b bool, format string, a ...any) {
	if !b {
		c.setReason("<assertion>")
		c.Failf(format, a...)
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
		}
		if fn != nil {
			ctx.Result = fn(f.Args)
		}
		if f.Pass != nil {
			f.Pass(ctx)
		}
	})
}

type Batch[A Argumenter, R Resulter] []Field[A, R]

func (b Batch[A, R]) Run(t *testing.T, fnName string, fn func(a A) R) {
	for _, test := range b {
		test.Run(t, fnName, fn)
	}
}

func Todo[Arg Argumenter, Res Resulter](a ...any) FnPass[Arg, Res] {
	return func(c *Ctx[Arg, Res]) {
		c.Assert(false, a...)
	}
}

func Todof[Arg Argumenter, Res Resulter](format string, a ...any) FnPass[Arg, Res] {
	return func(c *Ctx[Arg, Res]) {
		c.Assertf(false, format, a...)
	}
}
