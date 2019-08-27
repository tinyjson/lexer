package lexer

import (
	"errors"
	"strconv"
)

//go:generate ragel -Z lexer-machine.rl

// Lexer is the struct that contains json data and links to this data of parsed
// lexemes.
type Lexer struct {
	data        []byte
	p, pe, cs   int
	ts, te, act int
	stack       []int
	top         int
	eof         int
	// bracerPosPos is position in actions where to write closing bracer
	// position.
	bracerPosPos int

	pushArray bool
	// Controls contains sequence of lexem types of JSON.
	Controls []Control
	// Actions contains values depending on Controls:
	// Key:       data_start_pos, data_end_pos;
	// Nil:       nothing;
	// Bool:      0 for false, 1 for true;
	// Number:    data_start_pos, data_end_pos;
	// String:    data_start_pos, data_end_pos;
	// ArrayIn:   data_start_pos, data_end_pos, Controls_ArrayOut_pos,
	// Actions_pos - for fast skip;
	// ArrayOut:  nothing;
	// ObjectIn:  data_obj_start_pos, data_obj_end_pos, Controls_ObjectOut_pos,
	// Actions_pos - for fast skip;
	// ObjectOut: nothing.
	Actions []int
	// Just the same as Controls.
	controls []Control
	// Just the same as Actions.
	actions []int
}

var (
	ErrorUnexpectedEnd  = errors.New("unexpected end")
	ErrorUnexpectedType = errors.New("unexpected type")
	ErrorParseObject    = errors.New("can not parse object to int64")
	ErrorNilValue       = errors.New("nil value")
)

type Control int

const (
	Key Control = iota
	Nil
	Bool
	Number
	String
	ArrayIn
	ArrayOut
	ObjectIn
	ObjectOut
)

func NewLexer(data []byte) *Lexer {
	lex := &Lexer{}
	lex.Init(data)
	return lex
}

func (lex *Lexer) Init(data []byte) {
	lex.data = data
	lex.pe = len(lex.data)
	lex.eof = lex.pe
}

func (lex *Lexer) Error() bool {
	return lex.cs == 0
}

func (lex *Lexer) Data() []byte {
	return lex.data
}

func (lex *Lexer) ResetControls() {
	lex.Actions = lex.actions
	lex.Controls = lex.controls
}

func (lex *Lexer) ReadString() (string, error) {
	n := len(lex.Controls)
	if n == 0 {
		return "", ErrorUnexpectedEnd
	}
	switch lex.Controls[0] {
	case String:
		// TODO: unsafe to reduce string allocation?
		data := string(lex.data[lex.Actions[0]:lex.Actions[1]])
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[2:]
		return strconv.Unquote(data)
	case Number:
		data := string(lex.data[lex.Actions[0]:lex.Actions[1]])
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[2:]
		return data, nil
	case Nil:
		lex.Controls = lex.Controls[1:]
		return "null", ErrorNilValue
	case Bool:
		lex.Controls = lex.Controls[1:]
		t := lex.Actions[0] == 1
		lex.Actions = lex.Actions[1:]
		if t {
			return "true", nil
		} else {
			return "false", nil
		}
	case ArrayIn, ObjectIn:
		s := string(lex.data[lex.Actions[0]:lex.Actions[1]])
		lex.Controls = lex.controls[lex.Actions[2]:]
		lex.Actions = lex.actions[lex.Actions[3]:]
		return s, nil
	}

	return "", ErrorUnexpectedEnd
}

func (lex *Lexer) ReadInt() (int64, error) {
	n := len(lex.Controls)
	if n == 0 {
		return 0, ErrorUnexpectedEnd
	}
	switch lex.Controls[0] {
	case String:
		// TODO: unsafe to reduce string allocation?
		data := string(lex.data[lex.Actions[0]+1 : lex.Actions[1]-1])
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[2:]
		return strconv.ParseInt(data, 10, 64)
	case Number:
		data := string(lex.data[lex.Actions[0]:lex.Actions[1]])
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[2:]
		return strconv.ParseInt(data, 10, 64)
	case Nil:
		lex.Controls = lex.Controls[1:]
		return 0, ErrorNilValue
	case Bool:
		lex.Controls = lex.Controls[1:]
		res := int64(lex.Actions[0])
		lex.Actions = lex.Actions[1:]
		return res, nil
	case ArrayIn, ObjectIn:
		lex.Controls = lex.controls[lex.Actions[2]:]
		lex.Actions = lex.actions[lex.Actions[3]:]
		return 0, ErrorParseObject
	}

	return 0, ErrorUnexpectedEnd
}

func (lex *Lexer) ReadFloat() (float64, error) {
	n := len(lex.Controls)
	if n == 0 {
		return 0, ErrorUnexpectedEnd
	}
	switch lex.Controls[0] {
	case String:
		// TODO: unsafe to reduce string allocation?
		data := string(lex.data[lex.Actions[0]+1 : lex.Actions[1]-1])
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[2:]
		return strconv.ParseFloat(data, 64)
	case Number:
		data := string(lex.data[lex.Actions[0]:lex.Actions[1]])
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[2:]
		return strconv.ParseFloat(data, 64)
	case Nil:
		lex.Controls = lex.Controls[1:]
		return 0, ErrorNilValue
	case Bool:
		lex.Controls = lex.Controls[1:]
		res := float64(lex.Actions[0])
		lex.Actions = lex.Actions[1:]
		return res, nil
	case ArrayIn, ObjectIn:
		lex.Controls = lex.controls[lex.Actions[2]:]
		lex.Actions = lex.actions[lex.Actions[3]:]
		return 0, ErrorParseObject
	}

	return 0, ErrorUnexpectedEnd
}

func (lex *Lexer) ReadBool() (bool, error) {
	n := len(lex.Controls)
	if n == 0 {
		return false, ErrorUnexpectedEnd
	}
	switch lex.Controls[0] {
	case String:
		// TODO: unsafe to reduce string allocation?
		data := string(lex.data[lex.Actions[0]+1 : lex.Actions[1]-1])
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[2:]
		return data != "0" && data != "false", nil
	case Number:
		data := string(lex.data[lex.Actions[0]:lex.Actions[1]])
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[2:]
		return data != "0", nil
	case Nil:
		lex.Controls = lex.Controls[1:]
		return false, ErrorNilValue
	case Bool:
		res := int64(lex.Actions[0])
		lex.Actions = lex.Actions[1:]
		lex.Controls = lex.Controls[1:]
		return res == 1, nil
	case ArrayIn, ObjectIn:
		lex.Controls = lex.controls[lex.Actions[2]:]
		lex.Actions = lex.actions[lex.Actions[3]:]
		return true, ErrorParseObject
	}

	return false, ErrorUnexpectedEnd
}

func (lex *Lexer) SkipValue() {
	n := len(lex.Controls)
	if n == 0 {
		return
	}
	switch lex.Controls[0] {
	case String, Number:
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[2:]
	case Nil:
		lex.Controls = lex.Controls[1:]
	case Bool:
		lex.Actions = lex.Actions[1:]
		lex.Controls = lex.Controls[1:]
	case ArrayIn, ObjectIn:
		lex.Controls = lex.controls[lex.Actions[2]:]
		lex.Actions = lex.actions[lex.Actions[3]:]
	}
}
