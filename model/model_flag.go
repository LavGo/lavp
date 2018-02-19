package model

import (
	"fmt"
)

type InputFlag struct {
	FilePath string   //-f
	Method Method  // -m
	ParallelAmount int //-p
	Info bool // -i
	Url string // -U
	ContextType string // -C
}

func (self *InputFlag)ConvertMethod(val string)Method{
	switch val {
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	default:
		return GET

	}
}
func (self *InputFlag)Print()  {
	fmt.Printf("url : %v\n",self.Url)
	fmt.Printf("file_path : %v\n",self.FilePath)
	fmt.Printf("method : %v\n",self.Method)
	fmt.Printf("parallel_amount : %v\n",self.ParallelAmount)
	fmt.Printf("information_flag : %v\n",self.Info)
}

type Method int
const (
	POST Method = iota
	GET
	PUT
	DELETE
)

func (self *Method)ToString() string {
	switch *self {
	case POST:
		return "POSt"
	case GET:
		return "GET"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELET"
	default:
		return ""
	}
}
