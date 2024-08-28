package parser

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/Ayobami0/yoruba/src/ast"
	"github.com/Ayobami0/yoruba/src/lexer"
	"github.com/Ayobami0/yoruba/src/token"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "jeki" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func testNumberLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.NumberLiteral)
	if !ok {
		t.Errorf("il not *ast.NumberLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}
	return true
}

func testStringLiteral(t *testing.T, sl ast.Expression, value string) bool {
	str, ok := sl.(*ast.StringLiteral)
	if !ok {
		t.Errorf("sl not *ast.StringLiteral. got=%T", str)
		return false
	}

	if str.Value != value {
		t.Errorf("str.Value not %s. got=%s", value, str.Value)
		return false
	}

	if str.Token.Literal != value {
		t.Errorf("str.TokenLiteral not %s. got=%s", value, str.TokenLiteral())
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, sl ast.Expression, value bool) bool {
	bl, ok := sl.(*ast.BooleanLiteral)
	if !ok {
		t.Errorf("sl not *ast.BooleanLiteral. got=%T", bl)
		return false
	}

	if bl.Value != value {
		t.Errorf("bl.Value not %v. got=%v", value, bl.Value)
		return false
	}
	var val string

	if value {
		val = "ooto"
	} else {
		val = "eke"
	}

	if bl.Token.Literal != val {
		t.Errorf("bool.TokenLiteral not %s. got=%s", fmt.Sprintf("%v", value), bl.TokenLiteral())
		return false
	}
	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		nExp := exp.(*ast.NumberLiteral)
		return testNumberLiteral(t, nExp, int64(v))
	case string:
		switch exp.(type) {
		case *ast.StringLiteral:
			sExp := exp.(*ast.StringLiteral)
			return testStringLiteral(t, sExp, v)
		case *ast.Identifier:
			iExp := exp.(*ast.Identifier)
			return testIdentifier(t, iExp, v)
		}
	case bool:
		bExp := exp.(*ast.BooleanLiteral)
		return testBooleanLiteral(t, bExp, v)
	}

	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func TestLetStatements(t *testing.T) {
	b := bytes.NewBufferString(`jeki name je "Ayobami"
jeki age je 24
jeki school je "FUTA"
  `)

	l := lexer.New(b)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"name"},
		{"age"},
		{"school"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	b := bytes.NewBufferString(`da name pada
da "Ayobami" pada
da 69 pada
da (69 + 240) pada
`)

	l := lexer.New(b)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 4 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "da pada" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	b := bytes.NewBufferString(`print`)

	l := lexer.New(b)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "print" {
		t.Errorf("ident.Value not %s. got=%s", "print", ident.Value)
	}
	if ident.TokenLiteral() != "print" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "print",
			ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	b := bytes.NewBufferString("5")

	l := lexer.New(b)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.NumberLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
}

func TestIfStatement(t *testing.T) {
	b := bytes.NewBufferString(`
    ti x baje y lehinna
      "abi"
      da x pada
    abi
      ti x baje z lehinna
        "test"
      pari
      5
    pari
    `)

	l := lexer.New(b)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	ifStmt, ok := program.Statements[0].(*ast.IfStatement)

	if !ok {
		t.Fatalf("program.Statement[0] is not *ast.IfStatement")
	}

	consqBlk := ifStmt.Consequence

	if len(consqBlk.Statements) != 2 {
		t.Fatalf("Consequence block statements does not equal 2. got=%d",
			len(program.Statements))
	}

	consqTests := []struct{ eLiteral string }{{"abi"}, {"da pada"}}

	for i, v := range consqTests {
		if consqBlk.Statements[i].TokenLiteral() != v.eLiteral {
			t.Fatalf("consqBlk.Statements[%d].TokenLiteral() is not %s, got %s", i, v.eLiteral, consqBlk.TokenLiteral())
		}
	}

	altBlk := ifStmt.Alternative
	if len(altBlk.Statements) != 2 {
		t.Fatalf("Alternative block statements does not equal 2. got=%d",
			len(program.Statements))
	}
	altTests := []struct{ eLiteral string }{{"ti"}, {"5"}}

	for i, v := range altTests {
		if altBlk.Statements[i].TokenLiteral() != v.eLiteral {
			t.Fatalf("altBlk.Statements[%d].TokenLiteral() is not %s, got %s", i, v.eLiteral, altBlk.Statements[i].TokenLiteral())
		}
	}

}

func TestBooleanLiteralExpression(t *testing.T) {
	b := bytes.NewBufferString(`
    ooto
    eke
    `)

	l := lexer.New(b)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)
	if len(program.Statements) != 2 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		eType    token.TokenType
		eValue   bool
		eLiteral string
	}{
		{token.TRUE, true, "ooto"},
		{token.FALSE, false, "eke"},
	}

	for i, v := range tests {
		stmt, ok := program.Statements[i].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[%d] is not ast.ExpressionStatement. got=%T",
				i, program.Statements[i])
		}

		exp, ok := stmt.Expression.(*ast.BooleanLiteral)

		if !ok {
			t.Fatalf("exp not *ast.BooleanLiteral. got=%T", stmt.Expression)
		}

		if exp.Token.Type != v.eType {
			t.Fatalf("exp.Token.Type not %s. got %s", v.eType, exp.Token.Type)
		}
		if exp.TokenLiteral() != v.eLiteral {
			t.Fatalf("exp.TokenLiteral() not %s. got %s", v.eLiteral, exp.TokenLiteral())
		}
		if exp.Value != v.eValue {
			t.Fatalf("exp.Value not %v. got %v", v.eValue, exp.Value)
		}
	}

}

func TestStringLiteralExpression(t *testing.T) {
	b := bytes.NewBufferString("\"Ayobami\"")

	l := lexer.New(b)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != "Ayobami" {
		t.Errorf("literal.Value not %s. got=%s", "Ayobami", literal.Value)
	}

	if literal.TokenLiteral() != "Ayobami" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "Ayob",
			literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"-15", "-", 15},
	}
	for _, tt := range prefixTests {
		b := bytes.NewBufferString(tt.input)
		l := lexer.New(b)

		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testNumberLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 tobiju 5", 5, "tobiju", 5},
		{"5 kereju 5", 5, "kereju", 5},
		{"5 baje 5", 5, "baje", 5},
		{"5 kobaje 5", 5, "kobaje", 5},
		{"ooto kobaje eke", true, "kobaje", false},
	}
	for _, tt := range infixTests {
		b := bytes.NewBufferString(tt.input)
		l := lexer.New(b)

		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
		}
		if !testLiteralExpression(t, exp.Left, tt.leftValue) {
			return
		}
		if !testLiteralExpression(t, exp.Right, tt.rightValue) {
			return
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
	}
}

func TestFunctionStatement(t *testing.T) {
	b := bytes.NewBufferString(`
    ise print fname, lname se
      "nothing"
    pari
    `)

	l := lexer.New(b)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.FunctionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.FunctionStatement. got=%T",
			program.Statements[0])
	}
	if len(stmt.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n",
			len(stmt.Parameters))
	}
	testLiteralExpression(t, stmt.Parameters[0], "fname")
	testLiteralExpression(t, stmt.Parameters[1], "lname")
	if len(stmt.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n",
			len(stmt.Body.Statements))
	}
	bodyStmt, ok := stmt.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement. got=%T",
			stmt.Body.Statements[0])
	}

  testStringLiteral(t, bodyStmt.Expression, "nothing")
}
