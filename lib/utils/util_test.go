package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJsonMarshal(t *testing.T) {
	Convey("Test JSONMarshal", t, func() {
		So(JSONMarshal(123), ShouldEqual, "123")
		So(JSONMarshal(map[float64]int{}), ShouldEqual, `{"err_msg": "not json format"}`)
		So(JSONMarshal(nil), ShouldEqual, "")
	})

}

func TestTimeHourStringToUnix(t *testing.T) {
	Convey("Test TimeHourStringToUnix", t, func() {
		_, err := TimeHourStringToUnix("2021-07-18 20:19:22")
		So(err, ShouldBeNil)
		_, err = TimeHourStringToUnix("2021-07-18 20:19:22fffffff")
		So(err, ShouldNotBeNil)
	})
}

func TestInIntList(t *testing.T) {
	Convey("Test TestInIntList", t, func() {
		ans := InIntList([]int{1, 2, 3}, 3)
		So(ans, ShouldEqual, true)
		ans = InIntList([]int{1, 2, 3}, 4)
		So(ans, ShouldEqual, false)
	})
}

func TestTimeStringToUnix(t *testing.T) {

	Convey("Test TimeStringToUnix", t, func() {
		time, err := TimeStringToUnix("2021-07-18")
		So(err, ShouldBeNil)
		wantTime, err := TimeAddOneDayStringToUnix("2021-07-17")
		So(time, ShouldEqual, wantTime) //2021-07-18时间戳
		time, err = TimeStringToUnix("2021-17-18")
		So(err, ShouldNotBeNil)
		So(time, ShouldEqual, 0)
	})
}

func TestTimeAddOneDayStringToUnix(t *testing.T) {

	Convey("Test TimeStringToUnix", t, func() {
		time, err := TimeAddOneDayStringToUnix("2021-07-18")
		So(err, ShouldBeNil)
		wanted, err1 := TimeStringToUnix("2021-07-19")
		So(err1, ShouldBeNil)
		So(time, ShouldEqual, wanted) //2021-07-19时间戳
		time, err = TimeAddOneDayStringToUnix("2021-17-18")
		So(err, ShouldNotBeNil)
		So(time, ShouldEqual, 0)
	})
}

func TestSliceToSet(t *testing.T) {
	Convey("Test TimeSliceToSet", t, func() {
		ids := []int64{1, 2, 3}
		ids2 := map[string]string{
			"1": "1",
		}
		set := SliceToMap(ids)
		set2 := SliceToMap(ids2)
		So(set, ShouldNotBeNil)
		So(set2, ShouldNotBeNil)
		So(len(set2), ShouldEqual, 0)
		So(len(set), ShouldEqual, 3)
		So(set[1], ShouldNotBeNil)
	})
}

func TestInStringList(t *testing.T) {
	Convey("Test InStringList", t, func() {
		strs := []string{"1", "2", "3"}

		So(InStringList(strs, "1"), ShouldBeTrue)
		So(InStringList(strs, "2"), ShouldBeTrue)
		So(InStringList(strs, "3"), ShouldBeTrue)
		So(InStringList(strs, "4"), ShouldBeFalse)
	})
}
