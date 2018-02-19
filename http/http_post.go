package http

import (
	"github.com/LavGo/lavp/model"
	"time"
	"net/http"
	"github.com/hprose/hprose-golang/io"
)

type HttpPOST struct {

}

func (self *HttpPOST) Run(url string,ch chan *model.HttpResult,cType string,body []byte){
	timeStart:=time.Now()
	_,err:=http.Post(url,cType,io.NewByteReader(body))
	//fmt.Println(resp)
	timeEnd:=time.Now()
	var timeDua int=0
	if timeDua=timeEnd.Nanosecond()-timeStart.Nanosecond();timeDua<0{
		timeDua=0
	}
	ch<- &model.HttpResult{CostTime:timeDua/(1000*1000),Err:err}
}


