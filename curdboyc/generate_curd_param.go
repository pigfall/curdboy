package curdboyc

import(
	"path"
	"os"
	"fmt"
	tpl "text/template"

	"github.com/pigfall/curdboy/curdboyc/tpls"
)

type CURDParamGenerator struct{
	GlobalGraph *CURDGraphGenerator
}

func NewCURDParamGenerator(graph *CURDGraphGenerator)*CURDParamGenerator{
	return &CURDParamGenerator{
		GlobalGraph: graph,
	}
}

func (this *CURDParamGenerator) Generate()error{
	tplIns,err := tpl.New("").Parse(tpls.CURD_PARAM)
	if err != nil{
		return fmt.Errorf("Failed to parse the template of CURD_PARAM: %s",err.Error())
	}

	filepathToSave := path.Join(this.GlobalGraph.TargetDirPath(),"param.go")

	fileHandler,err :=os.Create(filepathToSave)
	if err != nil{
		return fmt.Errorf("Failed to create file %s",filepathToSave)
	}
	defer fileHandler.Close()

	err = tplIns.Execute(fileHandler,this)
	if err != nil{
		return fmt.Errorf("Failed to execute template CURD_PARAM: %w",err)
	}

	return nil
}
