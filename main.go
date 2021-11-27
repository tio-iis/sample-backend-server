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

	//http.HandleFunc("/", hanlder)
	//http.HandleFunc("/get", countHanlder)

	//サーバを立ち上げている
	http.ListenAndServe(":8080", nil)

	//サーバを停止すると、ここが実行される
	fmt.Println("shutdown")
}

var count int

func hanlder(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintln(w, count)
}

func countHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, count)
}
