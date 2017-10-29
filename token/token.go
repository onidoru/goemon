package token

// TokenType is the type of token
type TokenType string

// Token is a token type
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LSQBR     = "["
	RSQBR     = "]"
	LT        = "<"
	GT        = ">"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
	EQ        = "=="
	NOTEQ     = "!="
	DEC       = "--"
	INC       = "++"
	ADDTO     = "+="
	SUBTO     = "-="
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent chechs the keywords table to see whether the given identifier
// has a keyword. If it is, it returns the keyword's TokenType const.
// If not, it returns token.IDENT which is the TokenType for all user-defined id
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	if isNumber(ident) {
		return INT
	}
	if isWord(ident) {
		return IDENT
	}
	return ILLEGAL
}

func isWord(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r != '_') {
			return false
		}
	}
	return true
}

func isNumber(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
