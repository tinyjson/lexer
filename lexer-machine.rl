package lexer
%%{
    machine json_lexer;

    # Actions
    action KeyIn { lex.Controls, lex.Actions = append(lex.Controls, Key), append(lex.Actions, lex.p) }
    action KeyOut { lex.Actions = append(lex.Actions, lex.p) }
    action NilOut { lex.Controls = append(lex.Controls, Nil) }
    action BoolOutFalse { lex.Controls, lex.Actions = append(lex.Controls, Bool), append(lex.Actions, 0) }
    action BoolOutTrue { lex.Controls, lex.Actions = append(lex.Controls, Bool), append(lex.Actions, 1) }
    action NumberIn { lex.Controls, lex.Actions = append(lex.Controls, Number), append(lex.Actions, lex.p) }
    action NumberOut { lex.Actions = append(lex.Actions, lex.p) }
    action StringIn { lex.Controls, lex.Actions = append(lex.Controls, String), append(lex.Actions, lex.p) }
    action StringOut { lex.Actions = append(lex.Actions, lex.p) }
    action ArrIn { lex.Controls, lex.Actions = append(lex.Controls, ArrayIn), append(lex.Actions, lex.p, 0, 0, 0) }
    action ArrOut { lex.Controls, lex.Actions[lex.bracerPosPos] = append(lex.Controls, ArrayOut), lex.p+1 }
    action ObjIn { lex.Controls, lex.Actions = append(lex.Controls, ObjectIn), append(lex.Actions, lex.p, 0, 0, 0) }
    action ObjOut { lex.Controls, lex.Actions[lex.bracerPosPos] = append(lex.Controls, ObjectOut), lex.p+1 }

    action CallObj { fcall mainObject; }
    action CallArr { fcall mainArray; }
    action Ret { fret; }

    # States
    objectStart = '{';
    objectRef = objectStart >ObjIn @CallObj;
    arrayStart = '[';
    arrayRef = arrayStart >ArrIn @CallArr;

    ws = [ \t\nr]*;

    onenine = [1-9];
    digits = digit+;
    integer = '-'? (digit | onenine digits);
    fraction = ('.' digits)?;
    exponent = ([eE] [+\-]? digits)?;
    number = (integer fraction exponent);

    escape = ["\\/bfnrt] | ('u' xdigit xdigit xdigit);
    # TODO: [^"\\] is not as in documentation. Original input char can be up to 24bit.
    character = ('\\' escape) |  [^"\\];
    characters = character*;
    string = '"' characters '"' ;

    value = ( objectRef ) | ( arrayRef ) | (string >StringIn %StringOut) | (number >NumberIn %NumberOut) | ('true' %BoolOutTrue) | ('false' %BoolOutFalse) | ('null' %NilOut);
    element = ws value ws;

    key = string >KeyIn %KeyOut;
    member = ws key ws ':' element;
    members = member (',' member)*;

    # Object
    innerObject = ws | members;

    # Array
    elements = element (',' element)*;
    innerArray = ws | elements;

    mainArray := innerArray ']' @ArrOut @Ret;
    mainObject := innerObject '}' @ObjOut @Ret;
    main :=  element;

 	# Handle stack growth
	prepush {
		lex.stack = append(lex.stack, len(lex.Actions), 0)
		lex.top++
	}

 	# Handle stack shrinking
	postpop {
	    actPos := lex.stack[lex.top-1]
	    lex.Actions[actPos-2], lex.Actions[actPos-1] = len(lex.Controls), len(lex.Actions)
	    lex.bracerPosPos = actPos-3
		lex.stack = lex.stack[0:len(lex.stack)-2]
		lex.top--
	}

	# Data
	write data;

}%%

func (lex *Lexer) Parse() {
    %%{
        access lex.;
        variable p lex.p;
        variable pe lex.pe;
        variable eof lex.eof;
	    write init;
	    write exec;
	}%%
	lex.actions = lex.Actions
	lex.controls = lex.Controls
}
