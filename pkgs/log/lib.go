package log


import(
	"os"
"github.com/pigfall/gosdk/output"
)


func Fatal(args ...any){
	output.ErrWithRedColor(args...)
	os.Exit(1)
}

func Fatalf(format string,args ...any){
	output.ErrfWithRedColor(format,args...)
	os.Exit(1)
}

func Error(args ...any){
	output.ErrWithRedColor(args...)
}

func Errorf(format string,args ...any){
	output.ErrfWithRedColor(format,args...)
}

