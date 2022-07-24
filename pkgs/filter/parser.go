package filter

import (
	"fmt"
)

type Parser struct {
	tokens  []*Token
	current int
}

func NewParser(tokens []*Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (this *Parser) Parse() (Expr, error) {
	expr, err := this.expression()
	if err != nil {
		return nil, err
	}
	if !this.isAtEnd() {
		return nil, fmt.Errorf("unexpect token in pos %+v", this.current)
	}

	return expr, nil
}

func (this *Parser) expression() (Expr, error) {
	return this.binaryLogicalExpression()
}

func (this *Parser) binaryLogicalExpression() (Expr, error) {
	expr, err := this.unaryLogicalExpression()
	if err != nil {
		return nil, err
	}
	for this.match(TokenType_KW_And, TokenType_KW_Or) {
		token := this.prevToken()
		right, err := this.unaryLogicalExpression()
		if err != nil {
			return nil, err
		}
		expr = NewBinaryLogicalExpr(expr, right, token)
	}

	return expr, nil
}

func (this *Parser) unaryLogicalExpression() (Expr, error) {
	if this.match(TokenType_KW_Not) {
		bangToken := this.prevToken()
		//if err := this.expect(TokenType_LEFT_PARENTH); err != nil {
		//	return nil, err
		//}
		expr, err := this.unaryLogicalExpression()
		if err != nil {
			return nil, err
		}
		//if err := this.expect(TokenType_RIGHT_PARENTH); err != nil {
		//	return nil, err
		//}
		return NewUnaryExpr(bangToken, expr), nil
	}
	return this.primary()
}

func (this *Parser) primary() (Expr, error) {
	if this.match(TokenType_LEFT_PARENTH) {
		expr, err := this.expression()
		if err != nil {
			return nil, err
		}
		if err := this.expect(TokenType_RIGHT_PARENTH); err != nil {
			return nil, err
		}
		this.advance()
		return expr, nil
	}

	return this.comparision()
}

func (this *Parser) expect(tpe TokenType) error {
	if !this.check(tpe) {
		//panic("here")
		var names = make([]string, 0, len(this.tokens))
		for _, tk := range this.tokens {
			names = append(names, tk.Literal)
		}

		return fmt.Errorf("expected TokenType %+v in pos %d, token list: %+v", tpe.String(), this.current, names)
	}
	return nil
}

func (this *Parser) formatTks(tks []TokenType) []string {
	ret := make([]string, 0, len(tks))
	for _, v := range tks {
		ret = append(ret, v.String())
	}

	return ret
}

func (this *Parser) comparision() (Expr, error) {
	ok := this.match(TokenType_Identity)
	if !ok {
		return nil, fmt.Errorf("expected %+v in pos %d", TokenType_Identity.String(), this.current)
	}
	identifier := this.prevToken()
	wants := []TokenType{TokenType_KW_Eq, TokenType_KW_Ge, TokenType_KW_Gt, TokenType_KW_Lt, TokenType_KW_Le, TokenType_KW_Ne, TokenType_KW_Like}
	ok = this.match(wants...)
	if !ok {
		return nil, fmt.Errorf("expected %+v in pos %d", this.formatTks(wants), this.current)
	}
	op := this.prevToken()
	wants = []TokenType{TokenType_String, TokenType_Number}
	ok = this.match(wants...)
	if !ok {
		return nil, fmt.Errorf("expected %+v in pos %d", this.formatTks(wants), this.current)
	}
	literal := this.prevToken()
	return NewComparisionExpr(identifier, literal, op), nil
}

func (this *Parser) peek() *Token {
	return this.tokens[this.current]
}

func (this *Parser) check(expected TokenType) bool {
	if this.isAtEnd() {
		return false
	}
	return this.peek().Tpe == expected
}

func (this *Parser) isAtEnd() bool {
	return this.current >= len(this.tokens)
}

func (this *Parser) advance() *Token {
	tk := this.tokens[this.current]
	this.current++
	return tk
}

func (this *Parser) match(tpes ...TokenType) bool {
	for _, tpe := range tpes {
		if this.check(tpe) {
			this.advance()
			return true
		}
	}
	return false

}

func (this *Parser) prevToken() *Token {
	return this.tokens[this.current-1]
}
