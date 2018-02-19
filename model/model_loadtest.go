package model

import (
	"fmt"
	"time"
)

type HttpResult struct {
	CostTime int
	Err error
}

type LavpResult struct {
	URL string
	TPS uint64
	ParallelAmount int
	AvgTime int //ms
	TimeoutAmount int
	ProcessTime time.Duration //time
}

func (self *LavpResult)Print()  {
	fmt.Println("\r")
	fmt.Printf("URL : %v\n",self.URL)
	fmt.Printf("TPS(QPS) : %v\n",self.TPS)
	fmt.Printf("Parallel Amount : %v\n",self.ParallelAmount)
	fmt.Printf("Avg Cost Time(ms) : %v\n",self.AvgTime)
	fmt.Printf("Timeout Amount ： %v\n",self.TimeoutAmount)
	fmt.Printf("Lavp Process Time(s) ：%v\n",self.ProcessTime.Seconds())
}

