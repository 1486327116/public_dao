package dao

import (
	"fmt"
	"public_dao"
	"time"
)

type Bar struct {
	ID        int64     `gorm:"column:id" json:"id" form:"id"`
	BarName   string    `gorm:"column:bar_name" json:"bar_name" form:"bar_name"`
	Createdat time.Time `gorm:"column:createdat" json:"createdat" form:"createdat"`
	Updatedat time.Time `gorm:"column:updatedat" json:"updatedat" form:"updatedat"`
}

func (Bar) TableName() string {
	return "Bars"
}

func (p *Bar) NewOne() {
	record := public_dao.DB.NewRecord(Bar{
		BarName: "this public_dao test",
	})

	fmt.Println("success new one bar,",record)
}
