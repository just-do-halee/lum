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

type Context[Arg Argumenter, Res Resulter] struct {
	t         *testing.T
	FuncName  string
	Arguments Arg
	Result    Res
	reason    string
	logs      Stringify

	isFatal    bool
	isParallel bool
}

func (c *Context[Arg, Res]) IsParallel() bool {
	return c.isParallel
}

func (c *Context[Arg, Res]) setReason(a ...any) {
	c.reason = fmt.Sprint(a...)
}

func (c *Context[Arg, Res]) setReasonf(format string, a ...any) {
	c.reason = fmt.Sprintf(format, a...)
}

func (c *Context[Arg, Res]) String() string {
	return fmt.Sprintf("\n%s\n%s(%v)  -->  %v  \t%s\t", c.logs.String(), c.FuncName, c.Arguments, c.Result, c.reason)
}

func (c *Context[Arg, Res]) Log(a ...any) {
	c.logs.WriteStrings(a...)
	c.logs.Writeln()
}

func (c *Context[Arg, Res]) Logf(format string, a ...any) {
	c.logs.WriteString(fmt.Sprintf(format+"\n", a...))
}

func (c *Context[Arg, Res]) ResetLogs() {
	c.logs.Reset()
}

func (c *Context[Arg, Res]) Fatal(a ...any) {
	var sb Stringify
	sb.WriteString(c.String())
	sb.WriteStrings(a...)
	sb.Writeln()
	c.isFatal = true
	c.t.Fatal(sb.String())
}

func (c *Context[Arg, Res]) Fatalf(format string, a ...any) {
	c.isFatal = true
	c.t.Fatal(c.String(), fmt.Sprintf(format, a...), "\n")
}
func (c *Context[Arg, Res]) Fail(a ...any) {
	var sb Stringify
	sb.WriteString(c.String())
	sb.WriteStrings(a...)
	sb.Writeln()
	c.t.Error(sb.String())
}

func (c *Context[Arg, Res]) Failf(format string, a ...any) {
	c.t.Error(c.String(), fmt.Sprintf(format, a...), "\n")
}

const assertEqualFailFormat = "%v == %v [fail]"
const assertNotEqualFailFormat = "%v != %v [fail]"

func (c *Context[Arg, Res]) AssertResultEqual(rhs Res, a ...any) {
	if c.Result != rhs {
		c.setReasonf(assertEqualFailFormat, c.Result, rhs)
		c.Fail(a...)
	}
}

func (c *Context[Arg, Res]) AssertResultEqualf(rhs Res, format string, a ...any) {
	if c.Result != rhs {
		c.setReasonf(assertEqualFailFormat, c.Result, rhs)
		c.Failf(format, a...)
	}
}

func (c *Context[Arg, Res]) AssertEqual(lhs, rhs Res, a ...any) {
	if lhs != rhs {
		c.setReasonf(assertEqualFailFormat, lhs, rhs)
		c.Fail(a...)
	}
}

func (c *Context[Arg, Res]) AssertEqualf(lhs, rhs Res, format string, a ...any) {
	if lhs != rhs {
		c.setReasonf(assertEqualFailFormat, lhs, rhs)
		c.Failf(format, a...)
	}
}

func (c *Context[Arg, Res]) AssertResultNotEqual(rhs Res, a ...any) {
	if c.Result == rhs {
		c.setReasonf(assertNotEqualFailFormat, c.Result, rhs)
		c.Fail(a...)
	}
}

func (c *Context[Arg, Res]) AssertResultNotEqualf(rhs Res, format string, a ...any) {
	if c.Result == rhs {
		c.setReasonf(assertNotEqualFailFormat, c.Result, rhs)
		c.Failf(format, a...)
	}
}

func (c *Context[Arg, Res]) AssertNotEqual(lhs, rhs Res, a ...any) {
	if lhs == rhs {
		c.setReasonf(assertNotEqualFailFormat, lhs, rhs)
		c.Fail(a...)
	}
}

func (c *Context[Arg, Res]) AssertNotEqualf(lhs, rhs Res, format string, a ...any) {
	if lhs == rhs {
		c.setReasonf(assertNotEqualFailFormat, lhs, rhs)
		c.Failf(format, a...)
	}
}

func (c *Context[Arg, Res]) Assert(b bool, a ...any) {
	if !b {
		c.setReason("<assertion>")
		c.Fail(a...)
	}
}

func (c *Context[Arg, Res]) Assertf(b bool, format string, a ...any) {
	if !b {
		c.setReason("<assertion>")
		c.Failf(format, a...)
	}
}

func (c *Context[Arg, Res]) AssertFatalResultEqual(rhs Res, a ...any) {
	if c.Result != rhs {
		c.setReasonf(assertEqualFailFormat, c.Result, rhs)
		c.Fatal(a...)
	}
}

func (c *Context[Arg, Res]) AssertFatalResultEqualf(rhs Res, format string, a ...any) {
	if c.Result != rhs {
		c.setReasonf(assertEqualFailFormat, c.Result, rhs)
		c.Fatalf(format, a...)
	}
}

func (c *Context[Arg, Res]) AssertFatalEqual(lhs, rhs Res, a ...any) {
	if lhs != rhs {
		c.setReasonf(assertEqualFailFormat, lhs, rhs)
		c.Fatal(a...)
	}
}

func (c *Context[Arg, Res]) AssertFatalEqualf(lhs, rhs Res, format string, a ...any) {
	if lhs != rhs {
		c.setReasonf(assertEqualFailFormat, lhs, rhs)
		c.Fatalf(format, a...)
	}
}

func (c *Context[Arg, Res]) AssertFatalResultNotEqual(rhs Res, a ...any) {
	if c.Result == rhs {
		c.setReasonf(assertNotEqualFailFormat, c.Result, rhs)
		c.Fatal(a...)
	}
}

func (c *Context[Arg, Res]) AssertFatalResultNotEqualf(rhs Res, format string, a ...any) {
	if c.Result == rhs {
		c.setReasonf(assertNotEqualFailFormat, c.Result, rhs)
		c.Fatalf(format, a...)
	}
}

func (c *Context[Arg, Res]) AssertFatalNotEqual(lhs, rhs Res, a ...any) {
	if lhs == rhs {
		c.setReasonf(assertNotEqualFailFormat, lhs, rhs)
		c.Fatal(a...)
	}
}

func (c *Context[Arg, Res]) AssertFatalNotEqualf(lhs, rhs Res, format string, a ...any) {
	if lhs == rhs {
		c.setReasonf(assertNotEqualFailFormat, lhs, rhs)
		c.Fatalf(format, a...)
	}
}

func (c *Context[Arg, Res]) AssertFatal(b bool, a ...any) {
	if !b {
		c.setReason("<assertion>")
		c.Fatal(a...)
	}
}

func (c *Context[Arg, Res]) AssertFatalf(b bool, format string, a ...any) {
	if !b {
		c.setReason("<assertion>")
		c.Fatalf(format, a...)
	}
}

type Fn[Arg Argumenter, Res Resulter] func(Arg) Res

func FnMock[Arg Argumenter, Res Resulter]() Fn[Arg, Res] {
	return func(Arg) Res {
		return *new(Res)
	}
}

type FnPass[Arg Argumenter, Res Resulter] func(*Context[Arg, Res])

func FnPassMock[Arg Argumenter, Res Resulter]() FnPass[Arg, Res] {
	return func(*Context[Arg, Res]) {}
}

type Field[Arg Argumenter, Res Resulter] struct {
	Name string
	Args Arg
	Pass FnPass[Arg, Res]
	Loop uint
	// If Parallel is true, the field will be run in parallel
	// and it does not dedicate Loop in parallel.
	Parallel bool
}

func (f Field[Arg, Res]) Mock() Field[Arg, Res] {
	return Field[Arg, Res]{}
}

func (f Field[Arg, Res]) Run(t *testing.T, fnName string, beforeEach Fn[Arg, Res], afterEach FnPass[Arg, Res]) (ctx *Context[Arg, Res]) {
	t.Run(f.Name, func(t *testing.T) {
		ctx = &Context[Arg, Res]{
			t:         t,
			FuncName:  fnName,
			Arguments: f.Args,
		}
		if beforeEach == nil {
			beforeEach = FnMock[Arg, Res]()
		}
		if afterEach == nil {
			afterEach = FnPassMock[Arg, Res]()
		}
		if f.Pass == nil {
			f.Pass = FnPassMock[Arg, Res]()
		}
		if f.Loop == 0 {
			f.Loop = 1
		}
		if f.Parallel {
			t.Parallel()
		}
		isLoop := f.Loop > 1
		for i := uint(1); i <= f.Loop; i++ {
			ctx.ResetLogs()
			if isLoop {
				ctx.Logf("[LOOP] %d\n", i)
			}
			// before each and execute
			ctx.Result = beforeEach(ctx.Arguments)
			// test assert
			f.Pass(ctx)
			// after each
			afterEach(ctx)
		}
	})
	return
}

type Batch[Arg Argumenter, Res Resulter] []Field[Arg, Res]

func (b Batch[Arg, Res]) Mock() Batch[Arg, Res] {
	return Batch[Arg, Res]{}
}

func (b Batch[Arg, Res]) Run(t *testing.T, fnName string, beforeEach Fn[Arg, Res], afterEach FnPass[Arg, Res]) {
	for _, test := range b {
		if test.Run(t, fnName, beforeEach, afterEach).isFatal {
			break
		}
	}
}

func Todo[Arg Argumenter, Res Resulter](a ...any) FnPass[Arg, Res] {
	return func(c *Context[Arg, Res]) {
		c.Assert(false, a...)
	}
}

func Todof[Arg Argumenter, Res Resulter](format string, a ...any) FnPass[Arg, Res] {
	return func(c *Context[Arg, Res]) {
		c.Assertf(false, format, a...)
	}
}
