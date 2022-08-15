package concurrency

import (
	"context"
	"reflect"
	"sync"

	"code.byted.org/ad/gromore/lib/consts"
	"code.byted.org/ad/gromore/lib/utils/enhancelog"
)

func BatchHandleObjects(ctx context.Context, object interface{},
	handler func(object interface{}, errCh chan<- *consts.GMErr, wg *sync.WaitGroup)) *consts.GMErr {
	wg := sync.WaitGroup{}
	objects := make([]interface{}, 0, reflect.ValueOf(object).Len())
	for i := 0; i < reflect.ValueOf(object).Len(); i++ {
		objects = append(objects, reflect.ValueOf(object).Index(i).Interface())
	}
	errCh := make(chan *consts.GMErr, len(objects)) // 初始化一个带缓存的channel，用来存放错误码

	for _, obj := range objects {
		wg.Add(1)
		go handler(obj, errCh, &wg)
	}

	go func() {
		// 开启一个goroutine，等待上面全部执行完毕。然后关掉接收错误码的channel
		wg.Wait()
		close(errCh)
	}()

	// 一直阻塞并不断读取channel中的消息，直到channel被关闭
	for gmErr := range errCh {
		if gmErr != nil {
			enhancelog.CtxWarn(ctx, "Batch execute handler failed: %v", gmErr)
			return gmErr
		}
	}
	return nil
}
