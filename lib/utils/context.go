package util

import (
	"context"

	"code.byted.org/overpass/ad_pangle_aurora/kitex_gen/ad/pangle/aurora"

	"code.byted.org/ad/gromore/lib/consts"
	"code.byted.org/ad/gromore/lib/utils/enhancelog"
	"code.byted.org/ad/gromore/model"
)

// GetUserInfo 从上下文中获取用户信息
func GetUserInfo(ctx context.Context) *model.CtxUserInfo {
	if ctx.Value(consts.CtxUserInfo) == nil {
		return nil
	}
	if userInfo, ok := ctx.Value(consts.CtxUserInfo).(*model.CtxUserInfo); ok {
		return userInfo
	}
	return nil
}

func GetContextWithMockUserInfo(ctx context.Context) context.Context {
	ctxUserInfo := &model.CtxUserInfo{
		LoginBySubAccount:       true,
		SubAccountWithAdminRole: true,
		MediaID:                 459,
		SubAccountID:            459,
		SubAccountCanAccessSite: []int64{500331735, 50032888, 50030820, 50033173},
		Permissions:             []string{},
		Functions:               []string{},
	}
	ctx = context.WithValue(ctx, consts.CtxUserInfo, ctxUserInfo)
	return ctx
}

// PutToContext 将信息写入上下文
func PutToContext(ctx context.Context, key interface{}, info interface{}) context.Context {
	enhancelog.CtxInfo(ctx, "set to context key[%v], value[%v]", JSONMarshal(key), JSONMarshal(info))
	ctx = context.WithValue(ctx, key, info)
	return ctx
}

func CtxUser2AuroraUser(userInfo *model.CtxUserInfo) *aurora.UserInfo {
	return &aurora.UserInfo{
		UserId:     userInfo.MediaID,
		CoreUserId: userInfo.CoreUserID,
		SubUserId:  userInfo.SubAccountID,
		AdminId:    userInfo.AdminID,
	}
}
