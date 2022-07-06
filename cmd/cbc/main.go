package main

import (
	"github.com/pigfall/gosdk/flags"
	"github.com/pigfall/curdboy/pkgs/log"
	cbc "github.com/pigfall/curdboy/curdboyc"
)

func main() {
	// {
	schemaDirPathFlag := flags.NewParamString("schemaDirPath","","ent schema dir path",flags.ParamStringNotEmpty())
	targetDirPath := flags.NewParamString("targetDirPath","","which diretory to save th generated files",flags.ParamStringNotEmpty())
	entTargetDirPath := flags.NewParamString("entTargetDirPath","","ent generated direcotry path",flags.ParamStringNotEmpty())
	// }


	params := []flags.Param{schemaDirPathFlag,targetDirPath,entTargetDirPath}
	flags.FlagParams(params...)
	err := flags.ParseAndValidate(params)
	if err != nil{
		log.Fatalf("invalid arguments: %v",err)
	}

	cfg := cbc.NewConfig(schemaDirPathFlag.ValueAfterParsed,targetDirPath.ValueAfterParsed,entTargetDirPath.ValueAfterParsed)
	err = cbc.NewCURDGraphGenerator(cfg).Generate()
	if err != nil{
		log.Fatalf("generate curd graph failed: %v",err)
	}
}
