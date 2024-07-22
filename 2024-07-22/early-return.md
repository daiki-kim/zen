# 異常判定から書く早期リターンについて

今日はコードレビューで指摘してもらった早期リターンを書くメリットを簡単にまとめます！

## 早期リターンとは

複雑な条件式などで、先に例外を返す処理を書くことで処理全体のネストが深くなるのを防いだり、コードを読みやすくしたりできる書き方のことです。

## 早期リターンの具体例

Go言語で以下のようにユーザーデータを検証する関数があるとしましょう。以下のように上から正常判定を繰り返していると、いつの間にか深い谷が出来上がってしまいます。。

```go
// ユーザーデータを検証する関数
func validateUser(name string, age int, email string) (bool, error) {
    if name != "" { // 名前が入力されているか確認
        if age > 0 { // 年齢が0より大きいか確認
            if email != "" { // メールアドレスが入力されているか確認
                return true, nil
            } else {
                return false, errors.New("email is required")
            }
        } else {
            return false, errors.New("age must be greater than 0")
        }
    } else {
        return false, errors.New("name is required")
    }
}
```

これを早期リターンを使って異常判定から書くと以下のようになります。

```go
// ユーザーデータを検証する関数
func validateUser(name string, age int, email string) (bool, error) {
    if name == "" { // 名前が入力されていなければエラーを返す
        return false, errors.New("name is required")
    }
    if age < 0 { // 年齢がマイナスならエラーを返す
        return false, errors.New("age cannot be negative")
    }
    if email == "" { // メールアドレスが入力されていなければエラーを返す
        return false, errors.New("email is required")
    }
    return true, nil
}
```

極端な例を使いましたがかなりスッキリしました。今回の場合は正常判定を異常判定から書くという常套手段を用いました。

大事なことは、**無駄に条件を重ねるくらいならいらないものを弾いていこうという考え方**です。

## まとめ

今回は正常判定でネストが起こるコードを、異常判定から書くことでスッキリ書ける方法をまとめました。

早期リターンの考え方はこれ以外にもいろいろあって、以下のリンクが参考になりました。

https://zenn.dev/media_engine/articles/early_return

ただ個人的にはGoだとnil判定する慣習から、あまりネストしない傾向にある気もします。

今勉強中のluaというスクリプト言語だと気づかないうちにネストになったりしちゃって、改めて綺麗なコードの書き方を学び直さないとなあと感じております。。。