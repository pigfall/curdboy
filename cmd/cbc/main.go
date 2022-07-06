package main

import (
	"github.com/pigfall/gosdk/flags"
	"github.com/pigfall/curdboy/pkgs/log"
	cbc "github.com/pigfall/curdboy/curdboyc"
)

func main() {
	schemaDirPathFlag := flags.NewParamString("schemaDirPath","","ent schema dir path",flags.ParamStringNotEmpty())

	params := []flags.Param{schemaDirPathFlag}
	flags.FlagParams(params...)
	err := flags.ParseAndValidate(params)
	if err != nil{
		log.Fatalf("invalid arguments: %v",err)
	}

	err = cbc.NewCURDGraphGenerator().Generate()
	if err != nil{
		log.Fatalf("generate curd graph failed: %v",err)
	}
}
