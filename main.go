package main
import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Hobby     string `json:"hobby"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	person := new(Person)

	// URLパラメータの値を構造体に格納
	person.FirstName = r.FormValue("first")
	person.LastName = r.FormValue("last")
	person.Hobby = r.FormValue("hobby")

	// json に変換
	data, err := json.Marshal(person)

	// エラーがあった場合はエラーを表示
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(data))
}

// http://localhost:9000/ToJson?first=名前&last=苗字&hobby=趣味 にアクセスするとJsonデータを返します。
func main() {
	http.HandleFunc("/ToJson", handler)
    http.ListenAndServe(":9000", nil)
}