package filter



type TokenType int

const(
	TokenType_LEFT_PARENTH TokenType = iota+1
	TokenType_RIGHT_PARENTH 
	TokenType_Number 
	TokenType_String
	TokenType_Identity
	TokenType_KW_And 
	TokenType_KW_Or
	TokenType_KW_Not
	TokenType_KW_Eq
	TokenType_KW_Ne
	TokenType_KW_Gt
	TokenType_KW_Ge
	TokenType_KW_Lt
	TokenType_KW_Le
	TokenType_KW_Like
)

type Token struct{
	Tpe TokenType
	Literal string
	Value interface{}
}

func (this *Token) GetValue()interface{}{
	return this.Value
}

func (this *Token) IsNumber()bool{
	return this.Tpe == TokenType_Number 
}

func (this *Token) GetNumberValue()float64{
	return this.Value.(float64)
}

func (this *Token) GetStringValue()string{
	return this.Value.(string)
}



const(
	KEYWORD_AND = "and"
	KEYWORD_NOT = "not"
	KEYWORD_OR = "or"
	KEYWORD_EQ = "eq"
	KEYWORD_NE = "ne"
	KEYWORD_GT = "gt"
	KEYWORD_GE = "ge"
	KEYWORD_LT = "lt"
	KEYWORD_LE= "le"
	KEYWORD_LIKE= "like"
)

func (this TokenType) String()string{
	switch this{
	case TokenType_LEFT_PARENTH :
		return "TokenType_LEFT_PARENTH"
	case TokenType_RIGHT_PARENTH :
		return "TokenType_RIGHT_PARENTH"
	case TokenType_Number :
		return "TokenType_Number"
	case TokenType_String:
		return " TokenType_String"
	case TokenType_Identity:
		return "TokenType_Identity"
	case TokenType_KW_And :
		return "TokenType_KW_And "
	case TokenType_KW_Or:
		return " TokenType_KW_Or"
	case TokenType_KW_Not:
		return " TokenType_KW_Not"
	case TokenType_KW_Eq:
		return " TokenType_KW_Eq"
	case TokenType_KW_Ne:
		return " TokenType_KW_Ne"
	case TokenType_KW_Gt:
		return " TokenType_KW_Gt"
	case TokenType_KW_Ge:
		return " TokenType_KW_Ge"
	case TokenType_KW_Lt:
		return "TokenType_KW_Lt"
	case TokenType_KW_Le:
		return "TokenType_KW_Le"
	case TokenType_KW_Like:
		return "TokenType_KW_Like"
	default:
		return "TODO undefined token type"
	}
}
