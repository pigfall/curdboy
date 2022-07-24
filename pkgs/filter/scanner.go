package filter

import (
	"fmt"
	"github.com/pigfall/gosdk/ascii"
	"strconv"
)

type Scanner struct {
	source      []rune
	cursor      int
	cursorStart int
}

func NewScanner(source []rune) *Scanner {
	return &Scanner{
		source: source,
	}
}

func (this *Scanner) scanToken() (*Token, error) {
	var tk *Token
	var err error
	c := this.cursorAdvance()
	switch c {
	case '"', '\'':
		tk, err = this.scanRestString(c)
	case '(':
		tk = this.makeToken(TokenType_LEFT_PARENTH, this.cursorRange(), this.cursorRange())
	case ')':
		tk = this.makeToken(TokenType_RIGHT_PARENTH, this.cursorRange(), this.cursorRange())
	case '\t', ' ', '\n':
		return nil, nil
	default:
		if ascii.IsAlpha(byte(c)) {
			tk, err = this.scanRestIdentity()
		} else if ascii.IsNumber(byte(c)) {
			tk, err = this.scanRestNum()
		} else {
			return nil, fmt.Errorf("Unsupported character %v ", string(c))
		}
	}

	return tk, err
}

func (this *Scanner) Scan() ([]*Token, error) {
	tks := make([]*Token, 0)
	for !this.cursorIsEnd() {
		this.cursorStart = this.cursor
		tk, err := this.scanToken()
		if err != nil {
			return nil, err
		}
		if tk != nil {
			tks = append(tks, tk)
		}
	}

	return tks, nil
}

func (this *Scanner) scanRestNum() (*Token, error) {
	for ascii.IsNumber(byte(this.cursorPeek())) {
		this.cursorAdvance()
	}

	if this.cursorPeek() == '.' {
		this.cursorAdvance()
		for ascii.IsNumber(byte(this.cursorPeek())) {
			this.cursorAdvance()
		}
	}
	valueStr := this.cursorRange()
	// convert to float
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return nil, fmt.Errorf("parse number token %s to float failed %w", valueStr, err)
	}
	return this.makeToken(TokenType_Number, this.cursorRange(), (value)), nil
}

func (this *Scanner) scanRestIdentity() (*Token, error) {
	for {
		if ascii.IsAlphaNum(byte(this.cursorPeek())) || this.cursorPeek() == '_' || this.cursorPeek() == '.' {
			this.cursorAdvance()
		} else {
			break
		}
	}
	literal := this.cursorRange()
	switch literal {
	case KEYWORD_AND:
		return this.makeToken(TokenType_KW_And, literal, literal), nil
	case KEYWORD_OR:
		return this.makeToken(TokenType_KW_Or, literal, literal), nil
	case KEYWORD_NOT:
		return this.makeToken(TokenType_KW_Not, literal, literal), nil
	case KEYWORD_EQ:
		return this.makeToken(TokenType_KW_Eq, literal, literal), nil
	case KEYWORD_NE:
		return this.makeToken(TokenType_KW_Ne, literal, literal), nil
	case KEYWORD_GT:
		return this.makeToken(TokenType_KW_Gt, literal, literal), nil
	case KEYWORD_GE:
		return this.makeToken(TokenType_KW_Ge, literal, literal), nil
	case KEYWORD_LT:
		return this.makeToken(TokenType_KW_Lt, literal, literal), nil
	case KEYWORD_LE:
		return this.makeToken(TokenType_KW_Le, literal, literal), nil
	case KEYWORD_LIKE:
		return this.makeToken(TokenType_KW_Like, literal, literal), nil
	default:
		return this.makeToken(TokenType_Identity, literal, literal), nil
	}
}

func (this *Scanner) cursorRange() string {
	return string(this.source[this.cursorStart:this.cursor])
}

func (this *Scanner) makeToken(tkTpe TokenType, literal string, value interface{}) *Token {
	return &Token{
		Tpe:     tkTpe,
		Literal: literal,
		Value:   value,
	}
}

func (this *Scanner) scanRestString(matchedC rune) (*Token, error) {
	for !this.cursorIsEnd() {
		if this.cursorPeek() == matchedC {
			this.cursorAdvance()
			return this.makeToken(TokenType_String, this.cursorRange(), string(this.source[this.cursorStart+1:this.cursor-1])), nil
		}
		this.cursorAdvance()
	}

	return nil, fmt.Errorf("unclosed string")
}

func (this *Scanner) cursorAdvance() rune {
	ret := this.source[this.cursor]
	this.cursor++
	return ret
}

func (this *Scanner) cursorPeek() rune {
	if this.cursorIsEnd() {
		return []rune(ascii.NullString())[0]
	}
	return this.source[this.cursor]
}

func (this *Scanner) cursorIsEnd() bool {
	return this.cursor == len(this.source)
}
