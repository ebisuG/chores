// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	tasks := []string{
// 		"cmake ..",
// 		"cmake . --build Release",
// 		"cpack",
// 	}
// 	for _, task := range tasks {
// 		go func() {
// 			// goroutineが起動するときにはループが回りきって
// 			// 全部のtaskが最後のタスクになってしまう
// 			//実行すると、tasksの順序が変わって表示された
// 			fmt.Println(task)
// 		}()
// 	}
// 	time.Sleep(time.Second)
// }

// channelを使った書き換え版
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("start sub()")
// 	// 終了を受け取るためのチャネル
// 	done := make(chan bool)
// 	go func() {
// 		fmt.Println("sub() is finished")
// 		// 終了を通知
// 		done <- true
// 	}()
// 	// 終了を待つ、データ自体は不要なので捨てる
// 	<-done
// 	fmt.Println("all tasks are finished")
// }

// contextを利用した書き換え版
// package main

// import (
// 	"context"
// 	"fmt"
// )

// func main() {
// 	fmt.Println("start sub()")
// 	// 終了を受け取るための終了関数付きコンテキスト
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func() {
// 		fmt.Println("sub() is finished")
// 		// 終了を通知
// 		cancel()
// 	}()
// 	// 終了を待つ
// 	<-ctx.Done()
// 	fmt.Println("all tasks are finished")
// }

// select文の利用
package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "from channel 2"
	}()

	//動く
	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <-chan1:
	// 		fmt.Println("Received", msg1)
	// 	case msg2 := <-chan2:
	// 		fmt.Println("Received", msg2)
	// 	}
	// }

	//動かない
	// for msg1 := range chan1 {
	// 	if len(chan1) != 0 {
	// 		fmt.Println(msg1)
	// 	}
	// }
	// for msg2 := range chan2 {
	// 	if len(chan2) != 0 {
	// 		fmt.Println(msg2)
	// 	}
	// }

	//動くが、複数のmsg1,2が来る場合は対応できない
	for msg1 := range chan1 {
		fmt.Println(msg1)
		if len(chan1) == 0 {
			close(chan1)
		}
	}
	for msg2 := range chan2 {
		fmt.Println(msg2)
		if len(chan2) == 0 {
			close(chan2)
		}
	}
}
