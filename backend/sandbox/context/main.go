package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func main() {
	doSomeThingParallel(100)
}

func doSomeThingParallel(workerNum int) error {
	// 必要なコンテキストを生成する
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	// 正常完了時にコンテキストのリソースを解放
	defer cancel() // (4)
	// 複数のゴルーチンからエラーメッセージを
	// 集約するためにチャネルを用意する
	errCh := make(chan error, workerNum)
	// workerNum分の並行処理を行う
	wg := sync.WaitGroup{}
	for i := 0; i < workerNum; i++ { // (1)
		i := i
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			// エラーが発生すれば、キャンセル処理を行い、
			// エラーメッセージを送信する
			if err := doSomeThingWithContext(
				cancelCtx, num); err != nil { // (2)
				cancel()
				errCh <- err
			}
			return
		}(i)
	}
	// 並行処理の終了を待つ
	wg.Wait()

	// ↗左下の続き
	// エラーチャネルに入ったメッセージを取り出す
	close(errCh)
	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}
	// エラーが発生していれば、最初のエラーを返す
	if len(errs) > 0 {
		return errs[0]
	}
	// 正常終了
	return nil
}

// コンテキストを利用した何らかの処理をする関数
func doSomeThingWithContext(ctx context.Context, num int) error {
	// 処理に入る前に、コンテキストの死活を確認する
	select {
	case <-ctx.Done(): // (3)
		return ctx.Err()
	// コンテキストがまだキャンセルされていなければ、
	// そのまま処理に進む
	default:
	}
	if num > 50 {
		return errors.New("error")
	}
	fmt.Println(num)
	return nil
}

// プライベートタイプを宣言
type requestIDKey struct{} // (2)

// GetRequestID 外部からrequestIDを取得するための関数
func GetRequestID(ctx context.Context) (int, bool) {
	// 値取得と型のキャストを行い、値が存在しないか、キャストできない場合は 0, falseが返される
	r, ok := ctx.Value(requestIDKey{}).(int)
	if ok {
		return r, true
	}
	return 0, false
}

// WithRequestID 外部からRequestIDを保存するための関数
func WithRequestID(ctx context.Context, reqID int) context.Context {
	// パッケージ内で宣言したキーで値を保存
	return context.WithValue(ctx, requestIDKey{}, reqID)
}
