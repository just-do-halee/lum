package lum

import (
	"fmt"
	"strings"
)

type Stringify struct {
	strings.Builder
}

type Stringer interface {
	String() string
}

func (sb *Stringify) Writeln() {
	sb.WriteString("\n")
}

func (sb *Stringify) WriteStrings(s ...any) (total int, err error) {
	var num int
	for _, str := range s {
		switch str.(type) {
		case string:
			num, err = sb.WriteString(str.(string))
		case Stringer:
			num, err = sb.WriteString(str.(Stringer).String())
		default:
			num, err = sb.WriteString(fmt.Sprint(str))

		}
		total += num
		if err != nil {
			return total, err
		}
	}
	return total, nil
}
