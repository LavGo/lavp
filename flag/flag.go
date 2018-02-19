package flag

import (
	"github.com/LavGo/lavp/model"
	"flag"
)

func Parse()*model.InputFlag{
	fm:=&model.InputFlag{}
	flag.StringVar(&fm.FilePath,"f","./data.txt","The request data")
	var methodStr string
	flag.StringVar(&methodStr,"m","GET","Method of http protocol")
	flag.IntVar(&fm.ParallelAmount,"p",10,"The Amount of Parallel process.")
    flag.BoolVar(&fm.Info,"i",false,"OutPut the information of sys")
    flag.StringVar(&fm.Url,"U","","The address of lavp load test")
    flag.StringVar(&fm.ContextType,"C","","The content-type of http")
	flag.Parse()
	fm.Method=fm.ConvertMethod(methodStr)
	if fm.Info{
		fm.Print()
	}
	return fm

}



