package model

import "time"

type VKpi struct {
	TglEntri    time.Time
	TglStatus   time.Time
	TglTempo    time.Time
	KodeProduk  string
	Tujuan      string
	Status      string
	WaktuRespon string
	Kpi         int
}
