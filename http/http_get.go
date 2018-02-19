package http

import (
	"time"
	"github.com/LavGo/lavp/model"
	"net/http"
)

type HttpGet struct {
	
}

func (self *HttpGet) Run(url string,ch chan *model.HttpResult,cType string,body []byte){
	timeStart:=time.Now()
	_,err:=http.Get(url)
	//fmt.Println(url,resp,err)
	timeEnd:=time.Now()
	var timeDua int=0
	if timeDua=timeEnd.Nanosecond()-timeStart.Nanosecond();timeDua<0{
		timeDua=0
	}
	ch<- &model.HttpResult{CostTime:timeDua/(1000*1000),Err:err}
}
