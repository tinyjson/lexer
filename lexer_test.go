package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer_ReadString(t *testing.T) {
	for _, tt := range []struct {
		name string
		json string
		want string
	}{
		{
			name: "string",
			json: `{"v":"string"}`,
			want: `string`,
		},
		{
			name: "escaped_string",
			json: `{"v":"str\"ing"}`,
			want: `str"ing`,
		},
		{
			name: "bool",
			json: `{"v": true}`,
			want: `true`,
		},
		{
			name: "number",
			json: `{"v": 1}`,
			want: `1`,
		},
		{
			name: "null",
			json: `{"v": null}`,
			want: `null`,
		},
		{
			name: "array",
			json: `{"v": []}`,
			want: `[]`,
		},
		{
			name: "array_complex",
			json: `{"v": [1,"2",null]}`,
			want: `[1,"2",null]`,
		},
		{
			name: "object",
			json: `{"v": {}}`,
			want: `{}`,
		},
		{
			name: "object_complex",
			json: `{"v": {"k":"v"}}`,
			want: `{"k":"v"}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			lex := NewLexer([]byte(tt.json))
			lex.Parse()
			for i, c := range lex.Controls {
				if c == Key {
					lex.Controls = lex.Controls[i+1:]
					lex.Actions = lex.Actions[4+2:]
					break
				}
			}
			got, err := lex.ReadString()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLexer_ReadInt(t *testing.T) {
	for _, tt := range []struct {
		name    string
		json    string
		want    int64
		wantErr error
	}{
		{
			name: "string",
			json: `{"v":"1"}`,
			want: 1,
		},
		{
			name: "bool_false",
			json: `{"v": false}`,
			want: 0,
		},
		{
			name: "bool_true",
			json: `{"v": true}`,
			want: 1,
		},
		{
			name: "number",
			json: `{"v": 1}`,
			want: 1,
		},
		{
			name: "null",
			json: `{"v": null}`,
			want: 0,
		},
		{
			name:    "array",
			json:    `{"v": []}`,
			wantErr: ErrorParseObject,
		},
		{
			name:    "object",
			json:    `{"v": {}}`,
			wantErr: ErrorParseObject,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			lex := NewLexer([]byte(tt.json))
			lex.Parse()
			for i, c := range lex.Controls {
				if c == Key {
					lex.Controls = lex.Controls[i+1:]
					lex.Actions = lex.Actions[4+2:]
					break
				}
			}
			got, err := lex.ReadInt()
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLexer_ReadFloat(t *testing.T) {
	for _, tt := range []struct {
		name    string
		json    string
		want    float64
		wantErr error
	}{
		{
			name: "string",
			json: `{"v":"1"}`,
			want: 1,
		},
		{
			name: "bool_false",
			json: `{"v": false}`,
			want: 0,
		},
		{
			name: "bool_true",
			json: `{"v": true}`,
			want: 1,
		},
		{
			name: "number",
			json: `{"v": 1}`,
			want: 1,
		},
		{
			name: "number_fraction",
			json: `{"v": 1.2134}`,
			want: 1.2134,
		},
		{
			name: "number_epsilon",
			json: `{"v": 1e-42}`,
			want: 1e-42,
		},
		{
			name: "number_fraction_epsilon",
			json: `{"v": 1.342e-42}`,
			want: 1.342e-42,
		},
		{
			name: "null",
			json: `{"v": null}`,
			want: 0,
		},
		{
			name:    "array",
			json:    `{"v": []}`,
			wantErr: ErrorParseObject,
		},
		{
			name:    "object",
			json:    `{"v": {}}`,
			wantErr: ErrorParseObject,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			lex := NewLexer([]byte(tt.json))
			lex.Parse()
			for i, c := range lex.Controls {
				if c == Key {
					lex.Controls = lex.Controls[i+1:]
					lex.Actions = lex.Actions[4+2:]
					break
				}
			}
			got, err := lex.ReadFloat()
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLexer_ReadBool(t *testing.T) {
	for _, tt := range []struct {
		name    string
		json    string
		want    bool
		wantErr error
	}{
		{
			name: "string_num_true",
			json: `{"v":"1"}`,
			want: true,
		},
		{
			name: "string_num_false",
			json: `{"v":"0"}`,
			want: false,
		},
		{
			name: "string_true",
			json: `{"v":"true"}`,
			want: true,
		},
		{
			name: "string_false",
			json: `{"v":"false"}`,
			want: false,
		},
		{
			name: "bool_false",
			json: `{"v": false}`,
			want: false,
		},
		{
			name: "bool_true",
			json: `{"v": true}`,
			want: true,
		},
		{
			name: "null",
			json: `{"v": null}`,
			want: false,
		},
		{
			name:    "array",
			json:    `{"v": []}`,
			want:    true,
			wantErr: ErrorParseObject,
		},
		{
			name:    "object",
			json:    `{"v": {}}`,
			want:    true,
			wantErr: ErrorParseObject,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			lex := NewLexer([]byte(tt.json))
			lex.Parse()
			for i, c := range lex.Controls {
				if c == Key {
					lex.Controls = lex.Controls[i+1:]
					lex.Actions = lex.Actions[4+2:]
					break
				}
			}
			got, err := lex.ReadBool()
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
