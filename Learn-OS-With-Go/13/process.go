package main

import (
	"fmt"
	"os"
	"syscall"
)

// プロセスで見られる情報
func main() {
	fmt.Printf("プロセスID: %d\n", os.Getpid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())

	//Windowsには、プロセスグループとセッショングループに対応する情報がない。Getgidは-1を返す。
	// sid, _ := syscall.GET(os.Getpid())
	fmt.Fprintf(os.Stderr, "グループID: %d", syscall.Getgid())
	// fmt.Fprintf(os.Stderr, "グループID: %d セッションID: %d\n", syscall.Getgid(), sid)
}

//プロセスは誰かのユーザー権限で動作する。
//ユーザーはグループに所属する。
//ユーザー、グループには権限があり、ファイルシステムの読み・書き・実行の権限の制限に利用される

//WindowsはセキュリティIDで権限管理を行っていて、Go言語だと実装されていない

//すべてのプロセスは、コマンドライン引数/環境変数/終了コードの3つのデータを必ず持っている
//終了コードは親プロセスに返される
