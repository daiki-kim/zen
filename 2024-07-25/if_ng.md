# if文は正常系ではなく異常系で考えてみよう

if文の条件分岐は、"~がtrueなら..."という書き方をベースで習うため、私も今までこの書き方ベースで書いていました。

ですが研修では、「異常系で早期リターンすること」や「異常系はスルーしないように」と連日フィードバック頂いたので、今回は"~がfalseならエラーを出力してreturn"という異常系で条件分岐を考えるメリットをまとめようと思います。

## 原則：「異常系はスルーしてはいけない」

### 具体例

早速、正常系で条件判定している以下のGoの簡単なコードを例に考えましょう。

```Go
package main

import (
    "fmt"
    "io/ioutil"
)

// ファイルを読み込む関数
func readFile(filename string) (string, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func main() {
    filename := "example.txt"
    content, err := readFile(filename)
    
    if err == nil {
        fmt.Println("File content:", content) // 正常系
    } else {
        fmt.Println("No content to display")  // 異常系
    }
}
```

次に、これを異常系で条件判定した場合のコードを見てみましょう。

```Go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func readFile(filename string) (string, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func main() {
    filename := "example.txt"
    content, err := readFile(filename)
    
    if err != nil {
        log.Fatalf("Failed to read file: %v", err) // 異常系
    }
    
    fmt.Println("File content:", content) // 正常系
}
```

今回の例だと、エラーが発生した際に前者は"No content to display"と表示だけして後の処理を継続してしまいます。そのため**異常があるまま処理が進んで後に大きなトラブルにつながりかねません**。

### 異常系をスルーしてはいけない一般的な理由

他には以下のような理由もあるため、異常系の条件分岐は基本スルーしない方が良さそうです。

- システムの安定性と信頼性の低下
- データの整合性と品質の低下
- デバッグと保守の困難化
- セキュリティリスクの増大
- ユーザー体験の低下
- システム全体への影響

要は、**異常が起こったものを後回しにすると痛い目見る**よねってことです、自分もユーザーも。

## まとめ

今回はif文は異常系から書こうというテーマを扱いました。

独学でプログラミングを学んだ私にとって、この考え方はフィードバックを受けてとても納得しました。

もちろん、必ずしも毎回これでいいと思考停止で使っていいわけではないです。

大事なのは、今実装しているコードに問題が発生したときにちゃんとプログラムが停止して、原因を即座に突き止められるかという考え方だと思います。

この辺は経験によるところも大きいと思うので、実装経験を積んでいろんな考え方を学んでいこうと思います。
