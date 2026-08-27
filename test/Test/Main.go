package Test_Main

import (
	"gopurs/output/gopurs_runtime"
	"time"
	"fmt"
)

var myDateInternal = gopurs_runtime.Box(Date{Ms: float64(time.Now().UnixNano()) / 1e6})

var MyDate = myDateInternal

func init() {
	fmt.Println("TagOf MyDate:", Foreign_TagOf(myDateInternal).StrVal())
}
