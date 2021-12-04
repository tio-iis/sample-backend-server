//今は気にしなくていい
package main

//利用するライブラリを読み込んでいる
//JavaScriptの <src> によるライブラリ読み込みのGoバージョンである
import (
	"fmt"
	"net/http"
)

// go run main.go を実行してmain()から処理が始まる
func main() {
	fmt.Println("start")

	//第一引数にパスを指定する。
	//第二引数に処理を指定する。
	//curl http://localhost:8080/ で指定した処理が実行される
	http.HandleFunc("/", hanlder)

	//curl http://localhost:8080/get で指定した処理が実行される
	http.HandleFunc("/get", countHanlder)

	//サーバを立ち上げている
	//ここで処理が止まる
	//curl http://localhost:8080 でアクセスできるサーバが起動する
	http.ListenAndServe(":8080", nil)

	//ここは実行されません。
	fmt.Println("finish")
}

//各関数内で共有することができる変数
//サーバが起動している間は値が保持される
//サーバを停止すると初期値である 0 にリセットされてしまう
var count int

//関数
func hanlder(w http.ResponseWriter, r *http.Request) {
	//変数の値を +1 している
	count++
	fmt.Fprintln(w, count)

	//helloを表示する
	//fmt.Fprintln(w, "hello")
}

func countHanlder(w http.ResponseWriter, r *http.Request) {
	//変数の値を表示している
	fmt.Fprintln(w, count)
}
