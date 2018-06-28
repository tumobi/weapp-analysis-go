package main

import (
	"fmt"
	Analysis "github.com/tumobi/weapp-analysis-go"
)

func main()  {
	appid := ""
	secret := ""
	analysis := Analysis.NewAnalysis(appid, secret)
	result := analysis.GetDailyVisitTrend("20180625","20180625")
	fmt.Println(result)
}
