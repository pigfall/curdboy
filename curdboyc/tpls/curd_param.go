package tpls

const CURD_PARAM = `// {{.GlobalGraph.GeneratedPrelude}}
package {{.GlobalGraph.GeneratedPkgName}}

type {{.GlobalGraph.Generated_QueryRequestStructName}} struct{
	Filter string
	PageIndex int
	PageSize int
}

`
