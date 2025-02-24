// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// 指定した外部プログラムを実行し、消費した時間を計測するプログラム
// func main() {
// 	if len(os.Args) == 1 {
// 		return
// 	}
// 	cmd := exec.Command(os.Args[1], os.Args[2:]...)
// 	err := cmd.Run()
// 	if err != nil {
// 		panic(err)
// 	}
// 	state := cmd.ProcessState
// 	fmt.Printf("%s\n", state.String())
// 	fmt.Printf("  Pid: %d\n", state.Pid())
// 	fmt.Printf("  System: %v\n", state.SystemTime())
// 	fmt.Printf("  User: %v\n", state.UserTime())
// }

package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	count := exec.Command("./count")
	//count,exeの標準出力を取得
	stdout, _ := count.StdoutPipe()
	go func() {
		//取得した標準出力からscannerを取得
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf("(stdout) %s\n", scanner.Text())
		}
	}()
	err := count.Run()
	if err != nil {
		panic(err)
	}
}
