package http

import (
	"github.com/LavGo/lavp/model"
	"time"
	"math"
	"fmt"
	"strings"
)

type Http interface{
	Run(url string,ch chan *model.HttpResult,cType string,body []byte)
}
type HttpNone struct {
}

func (self *HttpNone)Run(url string,ch chan *model.HttpResult,cType string,body []byte)()  {

}

func Run(ref Http,url string,amount int,cType string,body []byte)*model.LavpResult  {
	channels:=make([]chan *model.HttpResult,amount)
	timeStart:=time.Now()
	for idx:=0;idx<amount;idx++{
		channels[idx]=make(chan *model.HttpResult)
		go ref.Run(url,channels[idx],cType,body)
	}
	return analRes(url,channels,timeStart)
}

func analRes(url string,ress []chan *model.HttpResult,timeStart time.Time)*model.LavpResult{
	var costTimeSum int
	var timeoutCount int
	for idx:=0;idx<len(ress);idx++{
		res :=  <-ress[idx]
		if res.Err!=nil&&strings.Index(res.Err.Error(),"Timeout")>=0{
			timeoutCount+=1
			panic(res.Err)
		}else {
			costTimeSum+=res.CostTime
		}

		fmt.Printf("\r%s",Bar((idx*100*5)/len(ress),((idx+1)*100)/len(ress)))
	}
	costTimeProcess:=time.Since(timeStart)
	/*fmt.Printf("%v\n",costTimeSum)
	fmt.Printf("%v\n",uint64(costTimeSum))
	fmt.Printf("%v\n",math.Float64frombits(uint64(costTimeSum)))
	fmt.Printf("%v\n",math.Float64frombits(uint64(len(ress))))*/
	return &model.LavpResult{
		URL:url,
		TPS:math.Float64bits(math.Float64frombits(uint64(len(ress)-timeoutCount))/costTimeProcess.Seconds()),
		AvgTime:costTimeSum/(len(ress)-timeoutCount),
		ParallelAmount:len(ress),
		TimeoutAmount:timeoutCount,
		ProcessTime:time.Since(timeStart),
	}
}
func Bar(vl int, percent int) string {
	res:=fmt.Sprintf("%s%c%02d%%", strings.Repeat("█", vl/10),  ([]rune(" ▏▎▍▌▋▋▊▉█"))[vl%10],percent)
	return res
}