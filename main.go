package main

import (
	flag2 "github.com/LavGo/lavp/flag"
	"github.com/LavGo/lavp/model"
	"github.com/LavGo/lavp/http"
	"io/ioutil"
)

var httpMethod http.Http
func main()  {
	inputFlag:=flag2.Parse()
	switch inputFlag.Method {
	case model.GET:
		httpMethod=new(http.HttpGet)
	case model.POST:
		httpMethod=new(http.HttpPOST)
	default:
		httpMethod=new(http.HttpNone)
	}
	body,_:=ioutil.ReadFile(inputFlag.FilePath)
	res:=http.Run(httpMethod,inputFlag.Url,inputFlag.ParallelAmount,inputFlag.ContextType,body)
	res.Print()
}