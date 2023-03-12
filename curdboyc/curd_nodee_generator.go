package curdboyc

import (
	"entgo.io/ent/entc/gen"

	"fmt"
	"os"
	"path"
	"strings"
	tpl "text/template"

	ent "github.com/pigfall/ent_utils"
)

type CURDNodeGenerator struct {
	TargetNode  *ent.Type
	GlobalGraph *CURDGraphGenerator
}

func NewCURDNodeGenerator(targetNode *ent.Type, graph *CURDGraphGenerator) *CURDNodeGenerator {
	return &CURDNodeGenerator{
		TargetNode:  targetNode,
		GlobalGraph: graph,
	}
}

func (this *CURDNodeGenerator) Generate() error {
	fileToSave := path.Join(this.GlobalGraph.TargetDirPath(), fmt.Sprintf("curd_%s.go", strings.ToLower(this.TargetNode.Name())))

	fd, err := os.Create(fileToSave)
	if err != nil {
		return fmt.Errorf("Failed to craete file %s: %w", fileToSave, err)
	}

	defer fd.Close()

	tplFuncs := tpl.FuncMap{
		"buildGenPredicateParam": func(field *gen.Field, op string) *GenPredicateparamInTpl {
			return &GenPredicateparamInTpl{
				Field:                field,
				Op:                   op,
				NodeName:             this.TargetNode.Name(),
				NodePredicatePkgName: this.EntNodePredicatePkgName(),
			}
		},
		"ToFirstCharacterUpper": func(input string) string {
			runes := []rune(input)
			first := strings.ToUpper(string(runes[0]))
			return first + string(input[1:])
		},
		"error": func(format string, args ...interface{}) (struct{}, error) {
			return struct{}{}, fmt.Errorf(format, args...)
		},
		"buildMap": func(kv ...interface{}) map[string]interface{} {
			m := make(map[string]interface{})
			for i := 0; i < len(kv); i += 2 {
				m[kv[i].(string)] = kv[i+1]
			}
			return m
		},
		"setMap": func(m map[string]interface{}, k string, v any) map[string]interface{} {
			m[k] = v
			return m
		},
		"capitalFirstLetter": func(input string) string {
			capAll := strings.ToUpper(input)
			tmp := []rune(input)
			tmp[0] = []rune(capAll)[0]
			return string(tmp)
		},
		"toLower": func(input string)string{
			return strings.ToLower(input)
		},
	}

	tplIns, err := tpl.New("curd_node.tmpl").Funcs(tplFuncs).ParseFS(templates, "tpls/curd_node.tmpl", "tpls/_helper.tmpl")
	if err != nil {
		return fmt.Errorf("Failed to parse template of CURD_NODE: %w", err)
	}

	err = tplIns.Execute(fd, this)
	if err != nil {
		return fmt.Errorf("Failed to execute template CURD_NODE: %w", err)
	}

	return nil
}

// { တ generated name
func (this *CURDNodeGenerator) GeneratedQueryFuncName() string {
	return fmt.Sprintf("%sQuery", this.TargetNode.Name())
}

func (this *CURDNodeGenerator) GenerateCreateFuncName() string {
	return fmt.Sprintf("%sCreate", this.TargetNode.Name())
}

func (this *CURDNodeGenerator) GeneratedCountFuncName() string {
	return fmt.Sprintf("%sCount", this.TargetNode.Name())
}

func (this *CURDNodeGenerator) GeneratedUpdateFuncName() string {
	return fmt.Sprintf("%sUpdate", this.TargetNode.Name())
}

func (this *CURDNodeGenerator) GeneratedDeleteFuncName() string {
	return fmt.Sprintf("%sDelete", this.TargetNode.Name())
}

// }

func (this *CURDNodeGenerator) Imports() []string {
	return []string{
		"context",
		this.GlobalGraph.EntPkgPath(),
		this.GlobalGraph.EntPredicatePkgPath(),
		this.GlobalGraph.FilterParserPkgPath(),
		this.EntNodePredicatePkgPath(),
	}
}

func (this *CURDNodeGenerator) EntNodePredicatePkgPath() string {
	return path.Join(this.GlobalGraph.EntPkgPath(), strings.ToLower(this.TargetNode.Name()))
}

func (this *CURDNodeGenerator) EntNodePredicatePkgName() string {
	return path.Base(this.EntNodePredicatePkgPath())
}

type GenPredicateparamInTpl struct {
	Field                *gen.Field
	Op                   string
	NodeName             string
	NodePredicatePkgName string
}
