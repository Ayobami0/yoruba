package token

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	IDENT = "IDENT"

	NUM = "NUM"
	STR = "STR"

	PLUS   = "PLUS"
	MINUS  = "MINUS"
	DIVIDE = "DIVIDE"
	TIMES  = "MULTIPLY"

	QUOTE = "QUOTE"
	COMMA = "COMMA"

	ASSIGNMENT = "ASSIGNMENT"
	FUNCTION   = "FUNCTION"
	LET        = "LET"
	EQL        = "EQUAL"
	NOTEQL     = "NOT_EQUAL"
	TRUE       = "TRUE"
	FALSE      = "FALSE"
	END        = "END"
	AND        = "AND"
	OR         = "OR"
	IF         = "IF"
	ELSE       = "ELSE"
	ELSEIF     = "ELSEIF"
	THEN       = "THEN"
	EXECUTE    = "EXECUTE"
	RTN_PREFIX = "RETURN_PREFIX"
	RTN_SURFIX = "RETURN_SURFIX"
	BREAK      = "BREAK"
	TILL       = "TILL"
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
	"abiti":  ELSEIF,
	"abi":     ELSE,
	"se":      EXECUTE,
	"lehinna": THEN,
	"da":      RTN_PREFIX,
	"pada":    RTN_SURFIX,
	"titi":    TILL,
	"fo":      BREAK,
}

func LookUp(ident string) TokenType {

	if t, ok := keywords[ident]; ok {
		return t
	}

	return IDENT
}
