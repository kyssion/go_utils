package util

import (
	"testing"
	"time"

	"code.byted.org/ad/gromore/lib/consts"
	. "code.byted.org/gopkg/mockito"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCustomTimerMock_GetNow(t *testing.T) {

	getNow := func() time.Time {
		return CustomTimerIns.GetNow()
	}

	PatchConvey("通过时钟接口mock当前时间", t, func() {
		timeMock := NewCustomTimerMock("2022-01-19 11:11:05")
		CustomTimerIns = timeMock
		now := getNow()
		So(now.Format(consts.Month2MinuteFormat), ShouldEqual, "01_19_11:11")

		CustomTimerIns = NewCustomTimer()
	})
}

func TestGetYestoday(t *testing.T) {
	PatchConvey("获取昨日的日期", t, func() {
		CustomTimerIns = NewCustomTimerMock("2022-03-08")
		So(GetYestoday(), ShouldEqual, "2022-03-07")
	})
}

func TestGetSevenDaysAgo(t *testing.T) {
	PatchConvey("获取前七日的日期", t, func() {
		CustomTimerIns = NewCustomTimerMock("2022-03-08")
		So(GetSevenDaysAgo(), ShouldEqual, "2022-03-01")
	})
}

func TestGetStrFromTimestamp(t *testing.T) {

	Convey("正常输入", t, func() {
		So(GetStrFromTimestamp(GetPtrInt64(1646296704), DateTemplateOfYearMonthDay), ShouldEqual, "2022-03-03")
	})

	Convey("异常输入->nil", t, func() {
		So(GetStrFromTimestamp(nil, DateTemplateOfYearMonthDay), ShouldEqual, "")
	})

	Convey("异常输入->0", t, func() {
		So(GetStrFromTimestamp(GetPtrInt64(0), DateTemplateOfYearMonthDay), ShouldEqual, "1970-01-01")
	})
}

func TestValidDate(t *testing.T) {
	PatchConvey("校验日期是否符合指定的格式", t, func() {
		Convey("符合", func() {
			So(ValidDate("2006-01-02", "2022-03-21"), ShouldEqual, true)
		})

		Convey("不符合", func() {
			So(ValidDate("2006-01-02", "20220321"), ShouldEqual, false)
		})
	})
}
