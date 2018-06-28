package analysis

import (
	"io/ioutil"
	"net/http"
	"bytes"
	"encoding/json"
)

const (
	ACCESS_TOKEN        = "https://api.weixin.qq.com/cgi-bin/token"
	DAILY_SUMMARY_TREND = "https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend"
	DAILY_VISIT_TREND   = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend"
	WEEKLY_VISIT_TREND  = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend"
	MONTHLY_VISIT_TREND = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend"
	VISIT_DISTRIBUTION  = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution"
	DAILY_RETAIN_INFO   = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo"
	WEEKLY_RETAIN_INFO  = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo"
	MONTHLY_RETAIN_INFO = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo"
	VISIT_PAGE          = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage"
	USER_PORTRAIT       = "https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type BeginAndEndDate struct {
	BeginDate string
	EndDate   string
}

type Analysis struct {
	Appid       string
	Secret      string
	AccessToken string
	ExpiresIn   int
}

func NewAnalysis(appid, secret string) Analysis {
	return Analysis{
		Appid:       appid,
		Secret:      secret,
		AccessToken: "",
		ExpiresIn:   0,
	}
}

func (this Analysis) getAccessToken() string {
	//发送请求
	resp, err := http.Get(ACCESS_TOKEN + "?grant_type=client_credential&appid=" +
		this.Appid + "&secret=" + this.Secret)
	defer resp.Body.Close()
	if err != nil || resp.StatusCode != http.StatusOK {
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
	if err != nil {
		return ""
	}
	if bytes.Contains(body, []byte("access_token")) {
		token := AccessToken{}
		err = json.Unmarshal(body, &token)
		if err != nil {
			return ""
		}
		this.AccessToken = token.AccessToken
		this.ExpiresIn = token.ExpiresIn
		return token.AccessToken
	}
	return ""
}

// 概况趋势（天）
func (this Analysis) GetdDailySummaryTrend(beginDate, endDate string) string {
	return this.sendRequest(DAILY_SUMMARY_TREND, BeginAndEndDate{
		beginDate,
		endDate,
	})
}

// 访问趋势（日趋势）
func (this Analysis) GetDailyVisitTrend(beginDate, endDate string) string {
	return this.sendRequest(DAILY_VISIT_TREND, BeginAndEndDate{
		beginDate,
		endDate,
	})
}

// 访问趋势（周趋势）
func (this Analysis) GetWeeklyVisitTrend(beginDate, endDate string) string {
	return this.sendRequest(WEEKLY_VISIT_TREND, BeginAndEndDate{
		beginDate,
		endDate,
	})
}

// 访问趋势（月趋势）
func (this Analysis) GetMonthlyVisitTrend(beginDate, endDate string) string {
	return this.sendRequest(MONTHLY_VISIT_TREND, BeginAndEndDate{
		beginDate,
		endDate,
	})
}

// 访问分布
func (this Analysis) GetVisitDistribution(beginDate, endDate string) string {
	return this.sendRequest(VISIT_DISTRIBUTION, BeginAndEndDate{
		beginDate,
		endDate,
	})
}

// 访问留存（日留存）
func (this Analysis) GetDailyRetainInfo(beginDate, endDate string) string {
	return this.sendRequest(DAILY_RETAIN_INFO, BeginAndEndDate{
		beginDate,
		endDate,
	})
}

// 访问留存（周留存）
func (this Analysis) GetWeeklyRetainInfo(beginDate, endDate string) string {
	return this.sendRequest(WEEKLY_RETAIN_INFO, BeginAndEndDate{
		beginDate,
		endDate,
	})
}

// 访问留存（月留存）
func (this Analysis) GetMonthlyRetainInfo(beginDate, endDate string) string {
	return this.sendRequest(MONTHLY_RETAIN_INFO, BeginAndEndDate{
		beginDate,
		endDate,
	})
}

// 访问页面
func (this Analysis) GetVisitPage(beginDate, endDate string) string {
	return this.sendRequest(VISIT_PAGE, BeginAndEndDate{
		beginDate,
		endDate,
	})
}

// 用户画像
func (this Analysis) GetUserPortrait(beginDate, endDate string) string {
	return this.sendRequest(USER_PORTRAIT, BeginAndEndDate{
		beginDate,
		endDate,
	})
}

func (this Analysis) sendRequest(uri string, param BeginAndEndDate) string {
	data := "{\"begin_date\":\"" + param.BeginDate +
		"\",\"end_date\":\"" + param.EndDate +
		"\"}"

	req, err := http.NewRequest("POST", this.getRequestUri(uri), bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		return ""
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func (this Analysis) getRequestUri(Uri string) string {
	var token string
	if this.AccessToken != "" {
		token = this.AccessToken
	} else {
		token = this.getAccessToken()
	}
	return Uri + "?access_token=" + token
}
