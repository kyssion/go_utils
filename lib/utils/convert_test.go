package util

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	. "code.byted.org/gopkg/mockito"
	. "github.com/smartystreets/goconvey/convey"

	"code.byted.org/ad/gromore/model"
)

func TestGetNumberValue(t *testing.T) {
	type args struct {
		value interface{}
	}

	type enum1 int64
	const enumOne = 1

	var num1 = 1
	var num2 int8 = 2
	var num3 int16 = 2
	var num4 int32 = 2
	var num5 int64 = 2
	var num6 = "3"

	tests := []struct {
		name  string
		args  args
		wantV int64
	}{
		{
			name: "测试int类型",
			args: args{
				num1,
			},
			wantV: 1,
		},
		{
			name: "测试int8类型",
			args: args{
				num2,
			},
			wantV: 2,
		},
		{
			name: "测试int16类型",
			args: args{
				num3,
			},
			wantV: 2,
		},
		{
			name: "测试int32类型",
			args: args{
				num4,
			},
			wantV: 2,
		},
		{
			name: "测试int64类型",
			args: args{
				num5,
			},
			wantV: 2,
		},
		{
			name: "测试非数字类型",
			args: args{
				num6,
			},
			wantV: MinNumber,
		},
		{
			name: "测试 -1",
			args: args{
				-1,
			},
			wantV: -1,
		},
		{
			name: "测试枚举类型",
			args: args{
				enumOne,
			},
			wantV: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := GetNumberValue(tt.args.value); gotV != tt.wantV {
				t.Errorf("GetNumberValue() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestConvertStruct(t *testing.T) {
	type args struct {
		originValue  interface{}
		targetValue  interface{}
		ignoreFields []string
	}

	type a struct {
		Name *string
		Age  int32
		Num  int
	}

	type b struct {
		Name string
		Age  int
		Num  *int32
	}

	name := "test"

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "测试转换结构体，忽略指定字段",
			args: args{
				originValue: &a{
					Name: &name,
					Age:  1,
				},
				targetValue:  &b{},
				ignoreFields: []string{"Name"},
			},
			want: &b{
				Age: 1,
			},
		},
		{
			name: "测试反向转换结构体",
			args: args{
				originValue: &b{
					Name: name,
					Age:  1,
				},
				targetValue: &a{},
			},
			want: &a{
				Name: &name,
				Age:  1,
			},
		},
		{
			name: "测试空指针",
			args: args{
				originValue: &b{
					Age: 1,
				},
				targetValue: &a{},
			},
			want: &a{
				Age: 1,
			},
		},
		{
			name: "非指针数值转指针数值",
			args: args{
				originValue: &a{
					Num: 1,
				},
				targetValue: &b{},
			},
			want: &b{
				Num: GetPtrInt32(1),
			},
		},
		{
			name: "指针数值转非指针数值",
			args: args{
				originValue: &b{
					Num: GetPtrInt32(1),
				},
				targetValue: &a{},
			},
			want: &a{
				Num: 1,
			},
		}, {
			name: "array-> string 测试",
			args: args{
				originValue: &model.AdSlotInfo{
					BiddingType: 3,
					MultilevelPrices: []*model.MultilevelPriceItem{
						{
							Tag:             "1",
							Price:           "2",
							MultilevelPrice: "3",
						},
					},
				},
				targetValue: &model.WaterfallCodeInfo{},
			},
			want: &model.WaterfallCodeInfo{
				BiddingType:      3,
				MultilevelPrices: "[{\"tag\":\"1\",\"price\":\"2\",\"multilevel_price\":\"3\"}]",
			},
		}, {
			name: "string-> array 测试",
			args: args{
				originValue: &model.WaterfallCodeInfo{
					BiddingType:      3,
					MultilevelPrices: "[{\"tag\":\"1\",\"price\":\"2\",\"multilevel_price\":\"3\"}]",
				},
				targetValue: &model.AdSlotInfo{},
			},
			want: &model.AdSlotInfo{
				BiddingType: 3,
				MultilevelPrices: []*model.MultilevelPriceItem{
					{
						Tag:             "1",
						Price:           "2",
						MultilevelPrice: "3",
					},
				},
			},
		},
		{
			name: "nil 测试",
			args: args{
				originValue: nil,
				targetValue: &model.AdSlotInfo{},
			},
			want: &model.AdSlotInfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStruct(context.Background(), tt.args.originValue, tt.args.targetValue, tt.args.ignoreFields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapFromStructWithFields(t *testing.T) {
	ctx := context.Background()

	type test11 struct {
		Age  int
		Name string
		ID   int
	}

	type args struct {
		ctx    context.Context
		value  interface{}
		fields []string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "测试结构体转换成map",
			args: args{
				ctx: ctx,
				value: &test11{
					Age:  1,
					Name: "111",
					ID:   123,
				},
				fields: []string{
					"age", "name", "id",
				},
			},
			want: map[string]interface{}{
				"age":  1,
				"name": "111",
				"id":   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMapFromStructWithFields(tt.args.ctx, tt.args.value, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMapFromStructWithFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssignWithUpdateValue(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx         context.Context
		oldValue    interface{}
		updateValue interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "测试用新值覆盖旧值",
			args: args{
				ctx: ctx,
				oldValue: &model.AdSlotInfo{
					ID:       123,
					SortType: 1,
					Network:  3,
				},
				updateValue: &model.AdSlotInfo{
					SortType: 2,
					Network:  1,
				},
			},
			want: &model.AdSlotInfo{
				SortType: 2,
				Network:  1,
				ID:       123,
			},
		},
		{
			name: "要更新的字段找不到",
			args: args{
				ctx: ctx,
				oldValue: &model.AdSlotInfo{
					ID:       123,
					SortType: 1,
					Network:  3,
				},
				updateValue: &model.WaterfallCodeInfo{
					ID:       2222,
					SortType: 3333,
					Network:  4444,
					Ctime:    time.Now(),
				},
			},
			want: &model.AdSlotInfo{
				ID:       2222,
				SortType: 3333,
				Network:  4444,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			println(tt.want.(*model.AdSlotInfo).MultilevelPrices)
			if got := AssignWithUpdateValue(tt.args.ctx, tt.args.oldValue, tt.args.updateValue, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssignWithUpdateValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStructFromMap(t *testing.T) {
	ctx := context.Background()

	PatchConvey("将map转换成指定的结构体对象", t, func() {
		Convey("转换成功，key的命名格式为下划线", func() {
			targetValue := GetStructFromMap(ctx, map[string]interface{}{
				"network":      1,
				"bidding_type": 1,
			}, &model.AdSlotInfo{}, UNDERLINE).(*model.AdSlotInfo)
			So(targetValue, ShouldResemble, &model.AdSlotInfo{
				Network:     1,
				BiddingType: 1,
			})
		})

		Convey("转换成功，转成指针", func() {
			targetValue := GetStructFromMap(ctx, map[string]interface{}{
				"network": 1,
				"price":   float64(12),
			}, &model.AdSlotInfo{}, UNDERLINE).(*model.AdSlotInfo)
			So(targetValue, ShouldResemble, &model.AdSlotInfo{
				Network: 1,
				Price:   GetPtrFloat64(12),
			})
		})

		Convey("转换成功，key的命名格式为驼峰", func() {
			targetValue := GetStructFromMap(ctx, map[string]interface{}{
				"Network":     1,
				"BiddingType": 1,
			}, &model.AdSlotInfo{}, CAMEL).(*model.AdSlotInfo)
			So(targetValue, ShouldResemble, &model.AdSlotInfo{
				Network:     1,
				BiddingType: 1,
			})
		})
	})
}

func TestDeepCopyByConvert(t *testing.T) {
	type Test1 struct {
		Price *float64 `json:"Price,omitempty"`
	}
	ctx := context.Background()

	PatchConvey("测试gob", t, func() {
		test1 := &Test1{
			Price: thrift.Float64Ptr(0),
		}
		test2 := &Test1{}
		DeepCopyByConvert(ctx, test1, test2)
		So(test2, ShouldResemble, &Test1{
			Price: thrift.Float64Ptr(0),
		})
	})
}

func BenchmarkDeepCopyByConvert(b *testing.B) {
	type Test1 struct {
		Price           *float64 `json:"Price,omitempty"`
		BiddingType     int      `json:"req_bidding_type"`
		Network         int      `json:"adn_name"`
		Priority        int      `json:"priority"`
		RitID           string   `json:"adn_slot_id"`
		CodeName        string   `json:"-"`
		SortType        int      `json:"sort_type"`
		ShowSort        int      `json:"show_sort"`
		LoadSort        int      `json:"load_sort"`
		Status          int      `json:"-"`
		Currency        string   `json:"-"`
		SettlementPrice *float64 `json:"-"`
		ID              int64    `json:"-"`
		OriginType      int      `json:"origin_type"`
	}
	ctx := context.Background()
	test1 := &Test1{
		Price:       thrift.Float64Ptr(0),
		BiddingType: 1,
		Status:      2,
		ID:          1322,
		OriginType:  1,
		Priority:    0,
		CodeName:    "11",
	}
	test2 := &Test1{}
	DeepCopyByConvert(ctx, test1, test2)
}

func TestDeepCopyMap(t *testing.T) {
	PatchConvey("深拷贝map", t, func() {
		Convey("成功", func() {
			src := map[string]interface{}{
				"a": 11,
			}
			dst := make(map[string]interface{})
			err := DeepCopyMap(context.Background(), src, &dst)
			So(err, ShouldBeNil)
			So(len(dst), ShouldEqual, 1)
		})

		Convey("传入的目标值不是引用类型，报错", func() {
			src := map[string]interface{}{
				"a": 11,
			}
			dst := make(map[string]interface{})
			defer func() {
				err := recover()
				So(err, ShouldEqual, "dst should be reference, dst: map")
			}()
			err := DeepCopyMap(context.Background(), src, dst)
			So(err, ShouldBeNil)
		})
	})
}
