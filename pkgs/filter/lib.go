package filter

import(
		"fmt"
)

func ParseFilter(source string)(Expr,error){
	tks,err :=  NewScanner([]rune(source)).Scan()
	if err != nil{
		return nil,fmt.Errorf("Scan Failed %w",err)
	}
	expr,err :=  NewParser(tks).Parse()
	if err != nil{
		return nil,fmt.Errorf("Parse failed from [ %s ], Reason: %w",source,err)
	}
	return expr,nil
}
