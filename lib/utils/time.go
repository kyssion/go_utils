package util

import (
	"time"
)

// 时间格式
const (
	DateFormat              = "2006-01-02"
	HourDateFormat          = "2006-01-02 15"
	TimeFormatISO           = "2006-01-02 15:04:05"
	HivePartitionDateFormat = "20060102"
	Month2MinuteFormat      = "01_02_15:04"
)
const DateTemplateOfYearMonthDay = "2006-01-02"

func GetUnixMill() int64 {
	return time.Now().UnixNano() / 1e6
}

var CustomTimerIns ICustomTimer

func init() {
	CustomTimerIns = NewCustomTimer()
}

type (
	ICustomTimer interface {
		GetNow() time.Time
	}
	CustomTimer struct {
	}
	CustomTimerMock struct {
		now time.Time
	}
)

func NewCustomTimer() ICustomTimer {
	return &CustomTimer{}
}

func (c *CustomTimer) GetNow() time.Time {
	return time.Now()
}

func GetToday() string {
	return CustomTimerIns.GetNow().Format(DateFormat)
}

func GetYestoday() string {
	return CustomTimerIns.GetNow().AddDate(0, 0, -1).Format(DateFormat)
}

func GetSevenDaysAgo() string {
	return CustomTimerIns.GetNow().AddDate(0, 0, -7).Format(DateFormat)
}

func NewCustomTimerMock(timeStr string) ICustomTimer {
	return &CustomTimerMock{now: *FormatTimeStr(timeStr)}
}

func (c *CustomTimerMock) GetNow() time.Time {
	return c.now
}

func GetStrFromTimestamp(timestamp *int64, timeFormatTpl string) string {
	if timestamp == nil {
		return ""
	}
	currentTime := time.Unix(*timestamp, 0)
	return currentTime.Format(timeFormatTpl)
}

func ValidDate(format string, date string) bool {
	_, err := time.ParseInLocation(format, date, time.Local)
	return err == nil
}
