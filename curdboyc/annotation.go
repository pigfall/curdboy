package curdboyc

import(
		"entgo.io/ent/entc/gen"
	"github.com/mitchellh/mapstructure"
)

type Annotation struct{
	List []AnnotationUnit
}

type AnnotationUnit struct{
	Left string// left alias
	Right string //right alias
	FromName string // inverse edge' name eg: edge.From(${FROM_NAME},Type)
}

func NewAnnotationUnit(left,right,fromName string)AnnotationUnit{
	return AnnotationUnit{
		Left:left,
		Right:right,
		FromName:fromName,
	}
}

func (this Annotation) Name()string{
	return "curdboy_edge_alias"
}

func GetAnnotation(edge *gen.Edge)(*Annotation,error){

			an,ok := edge.Annotations[(Annotation{}).Name()]
			if !ok{
				return nil,nil
			}
			anIns := &Annotation{}
			err := mapstructure.Decode(an,anIns)
			if err != nil{
				panic(err)
			}
			return anIns,nil
}
