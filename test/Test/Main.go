package Test_Main

import (
	"gopurs/output/Data.JSDate"
	"gopurs/output/Foreign"
	"gopurs/output/gopurs_runtime"
	"time"
	"fmt"
)

var MyDate = gopurs_runtime.Box(Data_JSDate.Date{Ms: float64(time.Now().UnixNano()) / 1e6})

func init() {
	fmt.Println("TagOf MyDate:", Foreign.TagOf(MyDate).StrVal())
}
