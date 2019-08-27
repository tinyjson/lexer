
//line lexer-machine.rl:1
package lexer

//line lexer-machine.go:6
var _json_lexer_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 2, 9, 14, 2, 10, 15, 
	2, 11, 13, 2, 12, 15, 3, 2, 
	10, 15, 3, 2, 12, 15, 3, 3, 
	10, 15, 3, 3, 12, 15, 3, 4, 
	10, 15, 3, 4, 12, 15, 3, 6, 
	10, 15, 3, 6, 12, 15, 3, 8, 
	10, 15, 3, 8, 12, 15, 
}

var _json_lexer_key_offsets []int16 = []int16{
	0, 0, 14, 16, 25, 31, 37, 43, 
	46, 48, 52, 54, 55, 56, 57, 58, 
	59, 60, 61, 62, 63, 64, 79, 81, 
	87, 93, 107, 110, 119, 121, 131, 135, 
	137, 145, 156, 157, 158, 159, 160, 166, 
	167, 168, 169, 175, 176, 177, 178, 184, 
	193, 199, 205, 211, 217, 219, 224, 229, 
	243, 245, 251, 257, 262, 271, 277, 283, 
	289, 292, 301, 303, 313, 317, 319, 327, 
	338, 339, 340, 341, 342, 348, 349, 350, 
	351, 357, 358, 359, 360, 366, 375, 381, 
	387, 393, 397, 401, 408, 416, 422, 431, 
	435, 439, 443, 443, 
}

var _json_lexer_trans_keys []byte = []byte{
	32, 34, 45, 48, 91, 102, 110, 114, 
	116, 123, 9, 10, 49, 57, 34, 92, 
	34, 47, 92, 98, 102, 110, 114, 116, 
	117, 48, 57, 65, 70, 97, 102, 48, 
	57, 65, 70, 97, 102, 48, 57, 65, 
	70, 97, 102, 48, 49, 57, 48, 57, 
	43, 45, 48, 57, 48, 57, 97, 108, 
	115, 101, 117, 108, 108, 114, 117, 101, 
	32, 34, 45, 48, 91, 93, 102, 110, 
	114, 116, 123, 9, 10, 49, 57, 34, 
	92, 32, 44, 93, 114, 9, 10, 32, 
	44, 93, 114, 9, 10, 32, 34, 45, 
	48, 91, 102, 110, 114, 116, 123, 9, 
	10, 49, 57, 48, 49, 57, 32, 44, 
	46, 69, 93, 101, 114, 9, 10, 48, 
	57, 32, 44, 69, 93, 101, 114, 9, 
	10, 48, 57, 43, 45, 48, 57, 48, 
	57, 32, 44, 93, 114, 9, 10, 48, 
	57, 32, 44, 46, 69, 93, 101, 114, 
	9, 10, 48, 57, 97, 108, 115, 101, 
	32, 44, 93, 114, 9, 10, 117, 108, 
	108, 32, 44, 93, 114, 9, 10, 114, 
	117, 101, 32, 44, 93, 114, 9, 10, 
	34, 47, 92, 98, 102, 110, 114, 116, 
	117, 48, 57, 65, 70, 97, 102, 48, 
	57, 65, 70, 97, 102, 48, 57, 65, 
	70, 97, 102, 32, 34, 114, 125, 9, 
	10, 34, 92, 32, 58, 114, 9, 10, 
	32, 58, 114, 9, 10, 32, 34, 45, 
	48, 91, 102, 110, 114, 116, 123, 9, 
	10, 49, 57, 34, 92, 32, 44, 114, 
	125, 9, 10, 32, 44, 114, 125, 9, 
	10, 32, 34, 114, 9, 10, 34, 47, 
	92, 98, 102, 110, 114, 116, 117, 48, 
	57, 65, 70, 97, 102, 48, 57, 65, 
	70, 97, 102, 48, 57, 65, 70, 97, 
	102, 48, 49, 57, 32, 44, 46, 69, 
	101, 114, 125, 9, 10, 48, 57, 32, 
	44, 69, 101, 114, 125, 9, 10, 48, 
	57, 43, 45, 48, 57, 48, 57, 32, 
	44, 114, 125, 9, 10, 48, 57, 32, 
	44, 46, 69, 101, 114, 125, 9, 10, 
	48, 57, 97, 108, 115, 101, 32, 44, 
	114, 125, 9, 10, 117, 108, 108, 32, 
	44, 114, 125, 9, 10, 114, 117, 101, 
	32, 44, 114, 125, 9, 10, 34, 47, 
	92, 98, 102, 110, 114, 116, 117, 48, 
	57, 65, 70, 97, 102, 48, 57, 65, 
	70, 97, 102, 48, 57, 65, 70, 97, 
	102, 32, 114, 9, 10, 32, 114, 9, 
	10, 32, 46, 69, 101, 114, 9, 10, 
	32, 69, 101, 114, 9, 10, 48, 57, 
	32, 114, 9, 10, 48, 57, 32, 46, 
	69, 101, 114, 9, 10, 48, 57, 32, 
	114, 9, 10, 32, 114, 9, 10, 32, 
	114, 9, 10, 
}

var _json_lexer_single_lengths []byte = []byte{
	0, 10, 2, 9, 0, 0, 0, 1, 
	0, 2, 0, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 11, 2, 4, 
	4, 10, 1, 7, 0, 6, 2, 0, 
	4, 7, 1, 1, 1, 1, 4, 1, 
	1, 1, 4, 1, 1, 1, 4, 9, 
	0, 0, 0, 4, 2, 3, 3, 10, 
	2, 4, 4, 3, 9, 0, 0, 0, 
	1, 7, 0, 6, 2, 0, 4, 7, 
	1, 1, 1, 1, 4, 1, 1, 1, 
	4, 1, 1, 1, 4, 9, 0, 0, 
	0, 2, 2, 5, 4, 2, 5, 2, 
	2, 2, 0, 0, 
}

var _json_lexer_range_lengths []byte = []byte{
	0, 2, 0, 0, 3, 3, 3, 1, 
	1, 1, 1, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 2, 0, 1, 
	1, 2, 1, 1, 1, 2, 1, 1, 
	2, 2, 0, 0, 0, 0, 1, 0, 
	0, 0, 1, 0, 0, 0, 1, 0, 
	3, 3, 3, 1, 0, 1, 1, 2, 
	0, 1, 1, 1, 0, 3, 3, 3, 
	1, 1, 1, 2, 1, 1, 2, 2, 
	0, 0, 0, 0, 1, 0, 0, 0, 
	1, 0, 0, 0, 1, 0, 3, 3, 
	3, 1, 1, 1, 2, 2, 2, 1, 
	1, 1, 0, 0, 
}

var _json_lexer_index_offsets []int16 = []int16{
	0, 0, 13, 16, 26, 30, 34, 38, 
	41, 43, 47, 49, 51, 53, 55, 57, 
	59, 61, 63, 65, 67, 69, 83, 86, 
	92, 98, 111, 114, 123, 125, 134, 138, 
	140, 147, 157, 159, 161, 163, 165, 171, 
	173, 175, 177, 183, 185, 187, 189, 195, 
	205, 209, 213, 217, 223, 226, 231, 236, 
	249, 252, 258, 264, 269, 279, 283, 287, 
	291, 294, 303, 305, 314, 318, 320, 327, 
	337, 339, 341, 343, 345, 351, 353, 355, 
	357, 363, 365, 367, 369, 375, 385, 389, 
	393, 397, 401, 405, 412, 419, 424, 432, 
	436, 440, 444, 445, 
}

var _json_lexer_indicies []byte = []byte{
	0, 2, 3, 4, 6, 7, 8, 0, 
	9, 10, 0, 5, 1, 12, 13, 11, 
	11, 11, 11, 11, 11, 11, 11, 11, 
	14, 1, 15, 15, 15, 1, 16, 16, 
	16, 1, 11, 11, 11, 1, 17, 18, 
	1, 19, 1, 20, 20, 21, 1, 21, 
	1, 22, 1, 23, 1, 24, 1, 25, 
	1, 26, 1, 27, 1, 28, 1, 29, 
	1, 30, 1, 31, 1, 32, 33, 34, 
	35, 37, 38, 39, 40, 32, 41, 42, 
	32, 36, 1, 44, 45, 43, 46, 47, 
	48, 46, 46, 1, 49, 50, 38, 49, 
	49, 1, 50, 33, 34, 35, 37, 39, 
	40, 50, 41, 42, 50, 36, 1, 51, 
	52, 1, 53, 54, 55, 56, 57, 56, 
	53, 53, 1, 58, 1, 53, 54, 56, 
	57, 56, 53, 53, 58, 1, 59, 59, 
	60, 1, 60, 1, 53, 54, 57, 53, 
	53, 60, 1, 53, 54, 55, 56, 57, 
	56, 53, 53, 52, 1, 61, 1, 62, 
	1, 63, 1, 64, 1, 65, 66, 67, 
	65, 65, 1, 68, 1, 69, 1, 70, 
	1, 71, 72, 73, 71, 71, 1, 74, 
	1, 75, 1, 76, 1, 77, 78, 79, 
	77, 77, 1, 43, 43, 43, 43, 43, 
	43, 43, 43, 80, 1, 81, 81, 81, 
	1, 82, 82, 82, 1, 43, 43, 43, 
	1, 83, 84, 83, 85, 83, 1, 87, 
	88, 86, 89, 90, 89, 89, 1, 91, 
	92, 91, 91, 1, 92, 93, 94, 95, 
	97, 98, 99, 92, 100, 101, 92, 96, 
	1, 103, 104, 102, 105, 106, 105, 107, 
	105, 1, 108, 109, 108, 85, 108, 1, 
	109, 84, 109, 109, 1, 102, 102, 102, 
	102, 102, 102, 102, 102, 110, 1, 111, 
	111, 111, 1, 112, 112, 112, 1, 102, 
	102, 102, 1, 113, 114, 1, 115, 116, 
	117, 118, 118, 115, 119, 115, 1, 120, 
	1, 115, 116, 118, 118, 115, 119, 115, 
	120, 1, 121, 121, 122, 1, 122, 1, 
	115, 116, 115, 119, 115, 122, 1, 115, 
	116, 117, 118, 118, 115, 119, 115, 114, 
	1, 123, 1, 124, 1, 125, 1, 126, 
	1, 127, 128, 127, 129, 127, 1, 130, 
	1, 131, 1, 132, 1, 133, 134, 133, 
	135, 133, 1, 136, 1, 137, 1, 138, 
	1, 139, 140, 139, 141, 139, 1, 86, 
	86, 86, 86, 86, 86, 86, 86, 142, 
	1, 143, 143, 143, 1, 144, 144, 144, 
	1, 86, 86, 86, 1, 145, 145, 145, 
	1, 146, 146, 146, 1, 147, 148, 149, 
	149, 147, 147, 1, 147, 149, 149, 147, 
	147, 19, 1, 147, 147, 147, 21, 1, 
	147, 148, 149, 149, 147, 147, 18, 1, 
	150, 150, 150, 1, 151, 151, 151, 1, 
	152, 152, 152, 1, 1, 1, 
}

var _json_lexer_trans_targs []byte = []byte{
	1, 0, 2, 7, 91, 94, 90, 11, 
	15, 18, 90, 2, 89, 3, 4, 5, 
	6, 91, 94, 92, 10, 93, 12, 13, 
	14, 95, 16, 17, 96, 19, 20, 97, 
	21, 22, 26, 27, 33, 24, 98, 34, 
	39, 43, 24, 22, 23, 47, 24, 25, 
	98, 24, 25, 27, 33, 24, 25, 28, 
	30, 98, 29, 31, 32, 35, 36, 37, 
	38, 24, 25, 98, 40, 41, 42, 24, 
	25, 98, 44, 45, 46, 24, 25, 98, 
	48, 49, 50, 51, 52, 99, 52, 53, 
	85, 54, 55, 54, 55, 56, 64, 65, 
	71, 58, 72, 77, 81, 58, 56, 57, 
	60, 58, 59, 99, 58, 59, 61, 62, 
	63, 65, 71, 58, 59, 66, 68, 99, 
	67, 69, 70, 73, 74, 75, 76, 58, 
	59, 99, 78, 79, 80, 58, 59, 99, 
	82, 83, 84, 58, 59, 99, 86, 87, 
	88, 90, 90, 90, 8, 9, 90, 90, 
	90, 
}

var _json_lexer_trans_actions []byte = []byte{
	0, 0, 15, 11, 11, 11, 19, 0, 
	0, 0, 25, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 15, 11, 11, 11, 19, 22, 0, 
	0, 0, 25, 0, 0, 0, 17, 17, 
	63, 0, 0, 0, 0, 13, 13, 0, 
	0, 55, 0, 0, 0, 0, 0, 0, 
	0, 7, 7, 39, 0, 0, 0, 5, 
	5, 31, 0, 0, 0, 9, 9, 47, 
	0, 0, 0, 0, 1, 28, 0, 0, 
	0, 3, 3, 0, 0, 15, 11, 11, 
	11, 19, 0, 0, 0, 25, 0, 0, 
	0, 17, 17, 67, 0, 0, 0, 0, 
	0, 0, 0, 13, 13, 0, 0, 59, 
	0, 0, 0, 0, 0, 0, 0, 7, 
	7, 43, 0, 0, 0, 5, 5, 35, 
	0, 0, 0, 9, 9, 51, 0, 0, 
	0, 17, 0, 13, 0, 0, 7, 5, 
	9, 
}

var _json_lexer_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 17, 0, 13, 13, 13, 13, 7, 
	5, 9, 0, 0, 
}

const json_lexer_start int = 1
const json_lexer_first_final int = 89
const json_lexer_error int = 0

const json_lexer_en_mainArray int = 21
const json_lexer_en_mainObject int = 51
const json_lexer_en_main int = 1


//line lexer-machine.rl:81


func (lex *Lexer) Parse() {
    
//line lexer-machine.go:277
	{
	 lex.cs = json_lexer_start
	 lex.top = 0
	}

//line lexer-machine.go:283
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	if  lex.cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_json_lexer_key_offsets[ lex.cs])
	_trans = int(_json_lexer_index_offsets[ lex.cs])

	_klen = int(_json_lexer_single_lengths[ lex.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case  lex.data[( lex.p)] < _json_lexer_trans_keys[_mid]:
				_upper = _mid - 1
			case  lex.data[( lex.p)] > _json_lexer_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_json_lexer_range_lengths[ lex.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case  lex.data[( lex.p)] < _json_lexer_trans_keys[_mid]:
				_upper = _mid - 2
			case  lex.data[( lex.p)] > _json_lexer_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_json_lexer_indicies[_trans])
	 lex.cs = int(_json_lexer_trans_targs[_trans])

	if _json_lexer_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_json_lexer_trans_actions[_trans])
	_nacts = uint(_json_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _json_lexer_actions[_acts-1] {
		case 0:
//line lexer-machine.rl:6
 lex.Controls, lex.Actions = append(lex.Controls, Key), append(lex.Actions, lex.p) 
		case 1:
//line lexer-machine.rl:7
 lex.Actions = append(lex.Actions, lex.p) 
		case 2:
//line lexer-machine.rl:8
 lex.Controls = append(lex.Controls, Nil) 
		case 3:
//line lexer-machine.rl:9
 lex.Controls, lex.Actions = append(lex.Controls, Bool), append(lex.Actions, 0) 
		case 4:
//line lexer-machine.rl:10
 lex.Controls, lex.Actions = append(lex.Controls, Bool), append(lex.Actions, 1) 
		case 5:
//line lexer-machine.rl:11
 lex.Controls, lex.Actions = append(lex.Controls, Number), append(lex.Actions, lex.p) 
		case 6:
//line lexer-machine.rl:12
 lex.Actions = append(lex.Actions, lex.p) 
		case 7:
//line lexer-machine.rl:13
 lex.Controls, lex.Actions = append(lex.Controls, String), append(lex.Actions, lex.p) 
		case 8:
//line lexer-machine.rl:14
 lex.Actions = append(lex.Actions, lex.p) 
		case 9:
//line lexer-machine.rl:15
 lex.Controls, lex.Actions = append(lex.Controls, ArrayIn), append(lex.Actions, lex.p, 0, 0, 0) 
		case 10:
//line lexer-machine.rl:16
 lex.Controls, lex.Actions[lex.bracerPosPos] = append(lex.Controls, ArrayOut), lex.p+1 
		case 11:
//line lexer-machine.rl:17
 lex.Controls, lex.Actions = append(lex.Controls, ObjectIn), append(lex.Actions, lex.p, 0, 0, 0) 
		case 12:
//line lexer-machine.rl:18
 lex.Controls, lex.Actions[lex.bracerPosPos] = append(lex.Controls, ObjectOut), lex.p+1 
		case 13:
//line lexer-machine.rl:20
 { 
		lex.stack = append(lex.stack, len(lex.Actions), 0)
		lex.top++
	 lex.stack[ lex.top] =  lex.cs;  lex.top++;  lex.cs = 51; goto _again
 } 
		case 14:
//line lexer-machine.rl:21
 { 
		lex.stack = append(lex.stack, len(lex.Actions), 0)
		lex.top++
	 lex.stack[ lex.top] =  lex.cs;  lex.top++;  lex.cs = 21; goto _again
 } 
		case 15:
//line lexer-machine.rl:22
  lex.top--;  lex.cs =  lex.stack[ lex.top]
{ 
	    actPos := lex.stack[lex.top-1]
	    lex.Actions[actPos-2], lex.Actions[actPos-1] = len(lex.Controls), len(lex.Actions)
	    lex.bracerPosPos = actPos-3
		lex.stack = lex.stack[0:len(lex.stack)-2]
		lex.top--
	 }
goto _again
 
//line lexer-machine.go:427
		}
	}

_again:
	if  lex.cs == 0 {
		goto _out
	}
	( lex.p)++
	if ( lex.p) != ( lex.pe) {
		goto _resume
	}
	_test_eof: {}
	if ( lex.p) == ( lex.eof) {
		__acts := _json_lexer_eof_actions[ lex.cs]
		__nacts := uint(_json_lexer_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _json_lexer_actions[__acts-1] {
			case 2:
//line lexer-machine.rl:8
 lex.Controls = append(lex.Controls, Nil) 
			case 3:
//line lexer-machine.rl:9
 lex.Controls, lex.Actions = append(lex.Controls, Bool), append(lex.Actions, 0) 
			case 4:
//line lexer-machine.rl:10
 lex.Controls, lex.Actions = append(lex.Controls, Bool), append(lex.Actions, 1) 
			case 6:
//line lexer-machine.rl:12
 lex.Actions = append(lex.Actions, lex.p) 
			case 8:
//line lexer-machine.rl:14
 lex.Actions = append(lex.Actions, lex.p) 
//line lexer-machine.go:461
			}
		}
	}

	_out: {}
	}

//line lexer-machine.rl:91

	lex.actions = lex.Actions
	lex.controls = lex.Controls
}
