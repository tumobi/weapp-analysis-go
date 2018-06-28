# weapp-analysis-go

微信小程序数据分析 GO SDK

### 环境要求
Go >= 1.9.x

### 安装
```
go get github.com/tumobi/weapp-analysis-go
```

### 使用
```go
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

```

### 接口列表
```go
// 概况趋势（天）
dailySummaryTrend := analysis.GetdDailySummaryTrend(beginDate, endDate)

// 访问趋势（日趋势）
dailyVisitTrend := analysis.GetDailyVisitTrend(beginDate, endDate)

// 访问趋势（周趋势）
weeklyVisitTrend := analysis.GetWeeklyVisitTrend(beginDate, endDate)

// 访问趋势（月趋势）
monthlyVisitTrend := analysis.GetMonthlyVisitTrend(beginDate, endDate)

// 访问分布
visitDistribution := analysis.GetVisitDistribution(beginDate, endDate)

// 访问留存（日留存）
dailyRetainInfo := analysis.GetDailyRetainInfo(beginDate, endDate)

// 访问留存（周留存）
weeklyRetainInfo := analysis.GetWeeklyRetainInfo(beginDate, endDate)

// 访问留存（月留存）
monthlyRetainInfo := analysis.GetMonthlyRetainInfo(beginDate, endDate)

// 访问页面
visitPage := analysis.GetVisitPage(beginDate, endDate)

// 用户画像
userPortrait := analysis.GetUserPortrait(beginDate, endDate)

```

### 详细接口文档
[数据 · 小程序](https://developers.weixin.qq.com/miniprogram/dev/api/analysis.html)
