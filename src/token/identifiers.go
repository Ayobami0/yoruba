package token

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	IDENT = "IDENT"

	NUM = "NUM"
	STR = "STR"

	R_GROUP = "}"
	L_GROUP = "{"

	PLUS   = "+"
	MINUS  = "-"
	DIVIDE = "/"
	TIMES  = "*"

	QUOTE = "QUOTE"
	COMMA   = "COMMA"

	ASSIGNMENT    = "ASSIGNMENT"
	FUNCTION      = "FUNCTION"
	LET           = "LET"
	EQL           = "EQUAL"
	NOTEQL        = "NOT_EQUAL"
	TRUE          = "TRUE"
	FALSE         = "FALSE"
	END           = "END"
	AND           = "AND"
	OR            = "OR"
	IF            = "IF"
	ELSE          = "ELSE"
	THEN          = "THEN"
	EXECUTE       = "EXECUTE"
	RTN_PREFIX    = "RETURN_PREFIX"
	RTN_SURFIX    = "RETURN_SURFIX"
	BREAK         = "BREAK"
	TILL          = "TILL"
	G_THAN        = "GREATER_THAN"
	L_THAN        = "LESS_THAN"
	F_CALL_PREFIX = "FUNCTION_CALL_PREFIX"
	F_CALL_INFIX  = "FUNCTION_CALL_INFIX"
	F_CALL_SURFIX = "FUNCTION_CALL_SURFIX"
)

var keywords = map[string]TokenType{
	"jeki":    LET,
	"je":      ASSIGNMENT,
	"ise":     FUNCTION,
	"baje":    EQL,
	"kobaje":  NOTEQL,
	"ooto":    TRUE,
	"eke":     FALSE,
	"pari":    END,
	"ati":     AND,
	"tabi":    OR,
	"ti":      IF,
	"abi":     ELSE,
	"se":      EXECUTE,
	"lehinna": THEN,
	"da":      RTN_PREFIX,
	"pada":    RTN_SURFIX,
	"titi":    TILL,
	"fo":      BREAK,
	"tobiju":  G_THAN,
	"kereju":  L_THAN,
	"pe":      F_CALL_PREFIX,
	"pelu":    F_CALL_INFIX,
	"pa":      F_CALL_SURFIX,
}

func LookUp(ident string) TokenType {

	if t, ok := keywords[ident]; ok {
		return t
	}

	return IDENT
}
