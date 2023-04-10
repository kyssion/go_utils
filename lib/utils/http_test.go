package util

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/bitly/go-simplejson"

	. "code.byted.org/gopkg/mockito"
	. "github.com/smartystreets/goconvey/convey"
)

var testNoBody = TestNoBody{}

type TestNoBody struct{}

func (TestNoBody) Read([]byte) (int, error)         { return 0, io.EOF }
func (TestNoBody) Close() error                     { return nil }
func (TestNoBody) WriteTo(io.Writer) (int64, error) { return 0, nil }

func Test_post(t *testing.T) {
	ctx := context.Background()
	PatchConvey("post 测试 1", t, func() {
		Convey("json 序列化异常", func() {
			Mock(json.Marshal).IncludeCurrentGoRoutine().Return([]byte{}, errors.New("err")).Build()
			_, err := Post(ctx, "", map[string]string{"data": "sdfsdf"}, nil)
			So(err, ShouldResemble, errors.New("err"))
		})
		Convey("请求异常", func() {
			Mock(json.Marshal).IncludeCurrentGoRoutine().Return([]byte{}, nil).Build()
			Mock(http.NewRequestWithContext).IncludeCurrentGoRoutine().Return(nil, errors.New("err")).Build()
			_, err := Post(ctx, "", map[string]string{"data": "sdfsdf"}, nil)
			So(err, ShouldResemble, errors.New("err"))
		})
	})
}

func Test_GetResponse(t *testing.T) {
	PatchConvey("post 测试 2", t, func() {
		Convey("io 获取异常", func() {
			Mock(ioutil.ReadAll).IncludeCurrentGoRoutine().Return([]byte{}, errors.New("err")).Build()
			_, err := GetResponseBody(nil, &http.Response{
				Body: testNoBody,
			})
			So(err, ShouldResemble, errors.New("err"))
		})
		Convey("正常", func() {
			Mock(ioutil.ReadAll).IncludeCurrentGoRoutine().Return([]byte{}, nil).Build()
			_, err := GetResponseBody(nil, &http.Response{
				Body: testNoBody,
			})
			So(err, ShouldResemble, nil)
		})
	})
}

func Test_doRequest(t *testing.T) {
	PatchConvey("post 测试 3", t, func() {
		Convey("req 获取异常", func() {
			Mock(doBaseRequest).IncludeCurrentGoRoutine().Return(&http.Response{}, errors.New("error")).Build()
			_, err := doRequest(nil, &http.Request{}, nil)
			So(err, ShouldResemble, errors.New("error"))
		})
		Convey("io 获取异常", func() {
			Mock(doBaseRequest).IncludeCurrentGoRoutine().Return(&http.Response{
				Body: testNoBody,
			}, nil).Build()
			Mock(ioutil.ReadAll).IncludeCurrentGoRoutine().Return([]byte{}, errors.New("error")).Build()
			_, err := doRequest(nil, &http.Request{}, nil)
			So(err, ShouldResemble, errors.New("error"))
		})
		Convey("json 序列化异常", func() {
			Mock(doBaseRequest).IncludeCurrentGoRoutine().Return(&http.Response{
				Body: testNoBody,
			}, nil).Build()
			Mock(ioutil.ReadAll).IncludeCurrentGoRoutine().Return([]byte{}, nil).Build()
			Mock(simplejson.NewJson).IncludeCurrentGoRoutine().Return(&simplejson.Json{}, errors.New("error")).Build()

			_, err := doRequest(nil, &http.Request{}, nil)
			So(err, ShouldResemble, errors.New("error"))
		})
		Convey("正常", func() {
			Mock(doBaseRequest).IncludeCurrentGoRoutine().Return(&http.Response{
				Body: testNoBody,
			}, nil).Build()
			Mock(ioutil.ReadAll).IncludeCurrentGoRoutine().Return([]byte{}, nil).Build()
			Mock(simplejson.NewJson).IncludeCurrentGoRoutine().Return(&simplejson.Json{}, nil).Build()

			_, err := doRequest(nil, &http.Request{}, nil)
			So(err, ShouldResemble, nil)
		})
	})
}
