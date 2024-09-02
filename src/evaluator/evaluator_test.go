package evaluator

import (
	"bytes"
	"testing"

	"github.com/Ayobami0/yoruba/src/lexer"
	"github.com/Ayobami0/yoruba/src/object"
	"github.com/Ayobami0/yoruba/src/parser"
)

func testEval(input string) object.Object {
	bInput := bytes.NewBufferString(input)
	l := lexer.New(bInput)
	p := parser.New(l)
	program := p.ParseProgram()

	env := object.NewEnvironment()

	return Eval(program, env)
}

func testNilValue(t *testing.T, object object.Object) {
	if object != nil {
		t.Errorf("expected object to be nil, got %v", object)
	}
}

func testNumberObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Number)
	if !ok {
		t.Errorf("object is not Number. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}
	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%s, want=%s",
			result.Value, expected)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
		return false
	}
	return true
}

func TestEvalStringExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"Ayobami"`, "Ayobami"},
		{`"10"`, "10"},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObject(t, evaluated, tt.expected)
	}
}

func TestEvalNumberExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * {5 + 10}", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * {3 * 3} + 10", 37},
		{"{5 + 10 * 2 + 15 / 3} * 2 + -10", 50},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ooto", true},
		{"eke", false},
		{"5 tobiju 6", false},
		{"2 kereju 10", true},
		{"17 baje 17", true},
		{"97 kobaje 17", true},
		{"{1 kereju 2} baje ooto", true},
		{"{1 kereju 2} baje eke", false},
		{"{1 tobiju 2} kobaje ooto", true},
		{"{1 tobiju 2} baje eke", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"jeki a je 5 a", 5},
		{"jeki a je 5 * 5 a", 25},
		{"jeki a je 5 jeki b je a b", 5},
		{"jeki a je 5 jeki b je a jeki c je a + b + 5 c", 15},
	}
	for _, tt := range tests {
		testNumberObject(t, testEval(tt.input), tt.expected)
	}
}

func TestIfStatement(t *testing.T) {
	input := `
    jeki i je 3
    ti {i baje 3 ati ooto} lehinna jeki e je "three" abi jeki e je "not three" pari
    e
    `
	testStringObject(t, testEval(input), "three")
}

func TestFunctionObject(t *testing.T) {
	input := "ise addTwo x se x + 2 pari"
	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}
	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}
	if fn.Parameters[0].TokenLiteral() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}
}

func TestFunctionCalls(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`ise add a, b se da a + b pada pari pe add pelu 1, 2 pa`, 3},
		{`ise sub a, b se da a - b pada pari pe sub pelu 1, 2 pa`, -1},
		{`ise orukoMiNi oruko se da oruko pada pari pe orukoMiNi pelu "Ayobami" pa`, "Ayobami"},
		{`ise orukoMiNiAyobami se da ooto pada pari pe orukoMiNiAyobami pa`, true},
		{`
      ise moDaGbaTo age se
        ti age kereju 18 lehinna
          da eke pada
        abi
          da ooto pada
        pari
      pari
      pe moDaGbaTo pelu 17 pa`, false},
		{`ise doNothing se pari pe doNothing pa`, nil},
	}

	for _, v := range tests {
		switch s := v.expected.(type) {
		case string:
			testStringObject(t, testEval(v.input), s)
		case int:
			testNumberObject(t, testEval(v.input), int64(s))
		case bool:
			testBooleanObject(t, testEval(v.input), s)
		case nil:
			testNilValue(t, testEval(v.input))
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"da 10 pada", 10},
		{"da 10 pada 9", 10},
		{"da 2 * 5 pada 9", 10},
		{"9 da {2 * 5} pada 9", 10},
		{"ti {3 baje 3} lehinna da 10 pada abi da 0 pada pari", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}

func TestLoopStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"jeki i je 0 titi i baje 5 se jeki i je i + 1 pari i", 5},
		{"jeki i je 0 jeki m je 0 titi i baje 5 se jeki j je 0 titi j baje 5 se jeki m je m + 1 jeki j je j + 1 pari jeki i je i + 1 pari m", 25},                        // Pew :-(
		{"jeki i je 0 jeki m je 0 titi i baje 5 se jeki j je 0 titi eke se ti j baje 5 lehinna fo pari jeki m je m + 1 jeki j je j + 1 pari jeki i je i + 1 pari m", 25}, // And i made it longer
		{"jeki i je 0 titi i baje 5 se ti i baje 3 lehinna fo pari jeki i je i + 1 pari i", 3},
		{"jeki i je 0 titi i baje 5 se jeki i je i + 1 ti i baje 3 lehinna fo pari pari i", 3},
		{"jeki i je 0 titi eke se jeki i je i + 1 ti i baje 100 lehinna fo pari pari i", 100},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}
