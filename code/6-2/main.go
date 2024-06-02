package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// 管理するためのCookieStoreオブジェクトをstoreとして定義
var store = sessions.NewCookieStore([]byte("secret-key"))

const (
	SessionKey  = "session-key"
	SessionName = "session-name"
)

// セッションをチェックし、認証されたユーザーが存在する場合はそのユーザーを返す関数
// セッションが存在しないかユーザーが認証されていない場合は、403(Forbidden)を返す
func CheckSession(w http.ResponseWriter, r *http.Request) {
	// 指定された名前を使用してセッションを取得
	session, _ := store.Get(r, SessionName)

	// セッションの値からユーザーを取得。型アサーションを使用して、値が文字列であることを確認
	// アサーションが失敗した場合、okはfalse
	user, ok := session.Values[SessionKey].(string)
	if !ok {
		// ユーザーが認証されていない場合は、許可されていないエラーを返す
		http.Error(w, "許可されていない", http.StatusForbidden)
		return
	}
	// ユーザーが認証されている場合は、認証されたユーザーをコンソールに表示
	fmt.Fprintf(w, "認証されたユーザー: %v", user)
}

// login関数は、SessionKeyのセッションの値を"email"に設定し、そのセッションを保存する関数
func login(w http.ResponseWriter, r *http.Request) {
	// 指定された名前を使用してセッションを取得
	session, _ := store.Get(r, SessionName)

	// セッションの値に値を設定
	session.Values[SessionKey] = "email"

	// セッションを保存
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/check", CheckSession)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
