package curdboyc

import(

"entgo.io/ent/entc/gen"

tpl "text/template"
	"strings"
 	ent "github.com/pigfall/ent_utils"
	"fmt"
	"os"
	"path"
)


type CURDNodeGenerator struct{
	TargetNode *ent.Type
	GlobalGraph *CURDGraphGenerator
}

func NewCURDNodeGenerator(targetNode *ent.Type,graph *CURDGraphGenerator)*CURDNodeGenerator{
	return &CURDNodeGenerator{
		TargetNode: targetNode,
		GlobalGraph: graph,
	}
}

func (this *CURDNodeGenerator) Generate()error{
	fileToSave := path.Join(this.GlobalGraph.TargetDirPath(),fmt.Sprintf("curd_%s.go",strings.ToLower(this.TargetNode.Name())))

	fd,err := os.Create(fileToSave)
	if err != nil{
		return fmt.Errorf("Failed to craete file %s: %w",fileToSave,err)
	}

	defer fd.Close()

	tplFuncs :=tpl.FuncMap{
		"buildGenPredicateParam":func(field *gen.Field,op string)*GenPredicateparamInTpl {
			return &GenPredicateparamInTpl{
				Field: field,
				Op:op,
				NodeName: this.TargetNode.Name(),
				NodePredicatePkgName: this.EntNodePredicatePkgName(),
			}
		},
	}

	tplIns,err := tpl.New("curd_node.tmpl").Funcs(tplFuncs).ParseFS(templates,"tpls/curd_node.tmpl")
	if err != nil{
		return fmt.Errorf("Failed to parse template of CURD_NODE: %w",err)
	}

	err = tplIns.Execute(fd,this)
	if err != nil{
		return fmt.Errorf("Failed to execute template CURD_NODE: %w",err)
	}

	return nil
}

func (this *CURDNodeGenerator)GeneratedQueryFuncName() string{
	return fmt.Sprintf("%sQuery",this.TargetNode.Name())
}


func (this *CURDNodeGenerator) Imports() []string{
	return []string{
		"context",
		this.GlobalGraph.EntPkgPath(),
		this.GlobalGraph.EntPredicatePkgPath(),
		this.GlobalGraph.FilterParserPkgPath(),
		this.EntNodePredicatePkgPath(),
	}
}

func (this *CURDNodeGenerator) EntNodePredicatePkgPath() string{
	return path.Join(this.GlobalGraph.EntPkgPath(),strings.ToLower(this.TargetNode.Name()))
}

func (this *CURDNodeGenerator) EntNodePredicatePkgName() string{
	return path.Base(this.EntNodePredicatePkgPath())
}

type GenPredicateparamInTpl struct{
	Field *gen.Field
	Op string
	NodeName string
	NodePredicatePkgName string
}
