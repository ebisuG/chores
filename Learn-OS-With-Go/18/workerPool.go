package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 計算: 元金均等
func calc(id, price int, interestRate float64, year int) {
	months := year * 12
	interest := 0
	for i := 0; i < months; i++ {
		balance := price * (months - i) / months
		interest += int(float64(balance) * interestRate / 12)
	}
	fmt.Printf("year=%d total=%d interest=%d id=%d\n", year, price+interest, interest, id)
}

// ワーカー
func worker(id, price int, interestRate float64, years chan int, wg *sync.WaitGroup) {
	// タスクがなくなってタスクのチャネルがcloseされるまで無限ループ
	for year := range years {
		calc(id, price, interestRate, year)
		wg.Done()
	}
}

func main() {
	// 借入額
	price := 40000000
	// 利子 1.1%固定
	interestRate := 0.011
	// タスク（各年の計算）はchanに格納
	years := make(chan int, 35)
	// タスクがなくなったらチャネルをクローズし、すべてのワーカーが終了する
	//ちなみに、クローズしなくてもこの場合は動くが、各workerでchannelのループを回しているため、クローズしたほうが良い
	defer close(years)
	// ファイルのクローズとは違い、オープンしっぱなしでも問題なさそう
	// https://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open

	for i := 1; i < 36; i++ {
		years <- i
	}
	var wg sync.WaitGroup
	wg.Add(35)
	// CPUコア数分のgoroutine起動
	//複数のgoroutineを起動して、一つのyearチャネルから読み込み、計算を行う
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, price, interestRate, years, &wg)
	}
	wg.Wait()
}
