//Goではselect/channelで、複数のgoroutineの同期がとれる
//他言語のコードをgoで再実装したい場合などは、syncパッケージを利用する

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// var id int
var id int64

// atomicを利用し、不可分操作として実装する
// DBのトランザクションみたいなもん？
func generateId(mutex *sync.Mutex) int64 {
	// Lock()/Unlock()をペアで呼び出してロックする
	mutex.Lock()
	defer mutex.Unlock()
	id++
	// return id
	return atomic.AddInt64(&id, 1)
}

//このmainでは、mainが終了すると、goroutineも終了してしまう
// func main() {
// 	// sync.Mutex構造体の変数宣言
// 	// 次の宣言をしてもポインタ型になるだけで正常に動作します
// 	// mutex := new(sync.Mutex)
// 	var mutex sync.Mutex

// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			fmt.Printf("id: %d\n", generateId(&mutex))
// 		}()
// 	}
// }

// wg.Wait()を利用して、mainの終了をgoroutineの終了まで待つ版
func main() {
	var wg sync.WaitGroup
	var maxCount int = 10
	wg.Add(maxCount)

	var mutex sync.Mutex

	for i := 0; i < maxCount; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
			wg.Done()
		}()
	}
	wg.Wait()
}

//
