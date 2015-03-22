package models

import (
	"github.com/astaxie/beego"
	"time"
)

type cpippi struct {
	Id    int64
	Qgcpi float64
	Hncpi float64
	Qgppi float64
	Hnppi float64
	Date  time
}

func (cp *cpippi) GetAll() {

}
