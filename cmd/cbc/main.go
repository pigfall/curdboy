package main

import (
	cbc "github.com/pigfall/curdboy/curdboyc"
	"github.com/pigfall/curdboy/pkgs/log"
	"github.com/pigfall/gosdk/flags"
	"os"
)

func main() {
	// {
	schemaDirPathFlag := flags.NewParamString("schemaDirPath", "", "ent schema dir path", flags.ParamStringNotEmpty())
	targetDirPath := flags.NewParamString("targetDirPath", "", "which diretory to save th generated files", flags.ParamStringNotEmpty())
	entTargetDirPath := flags.NewParamString("entTargetDirPath", "", "ent generated direcotry path", flags.ParamStringNotEmpty())
	// }

	params := []flags.Param{schemaDirPathFlag, targetDirPath, entTargetDirPath}
	flags.FlagParams(params...)
	err := flags.ParseAndValidate(params)
	if err != nil {
		log.Fatalf("invalid arguments: %v", err)
	}

	cfg := cbc.NewConfig(schemaDirPathFlag.ValueAfterParsed, targetDirPath.ValueAfterParsed, entTargetDirPath.ValueAfterParsed)
	generator, err := cbc.LoadCURDGraphGenerator(cfg)
	if err != nil {
		log.Fatalf("generate curd graph failed: %v", err)
		os.Exit(1)
	}
	err = generator.Generate()
	if err != nil {
		log.Fatalf("generate curd graph failed: %v", err)
		os.Exit(1)
	}

}
