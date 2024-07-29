# Roblox: Destroy()とnil宣言

Partを消す時など、以下のように`Destroy()`を使うと思います。

```lua
local part = Instance.new("Part")
part.Name = "newPart"
part.Parent = workpace

-- partを使った処理を実装

part:Destroy()

print(part.Name) -- "newPart"
```

これでpartが削除されてると思っていたのですが、これで`print(part.Name)`を実行すると、"newPart"と表示されてしまいます。

## `Destroy()`の仕様

インスタンスに対して`Destroy()`を実行すると、**そのインスタンスのParentをnilにして、そのインスタンスの全てのchildに対しても`Destroy()`を実行**します。

ただし、このインスタンス自身を破棄するわけではないため、先ほどのように削除後にもプロパティが表示されてしまっていたわけです。

## `Destroy()`+nilを代入する

不要になったPartを完全に削除するには、以下のようにその変数をnil宣言する必要があります。

```lua
local part = Instance.new("Part")
part.Name = "newPart"
part.Parent = workpace

-- partを使った処理を実装

part:Destroy()
part = nil -- nilを代入

print(part.Name) -- 表示されない
```

## nilを代入するメリット

`Destroy()`+nil代入のメリットは以下の通りです。

- **メモリリークの防止**:
Destroy()を呼び出しても、変数自体は依然としてインスタンスへの参照を保持しています。nilを代入することで、この参照を完全に解放し、ガベージコレクションを促進します。

- **誤用の防止**:
Destroy()後も変数は元のインスタンスのプロパティにアクセスできてしまいます。nilを代入することで、誤ってDestroy済みのインスタンスを使用するのを防ぎます。

- **コードの明確化**:
nilを代入することで、そのインスタンスがもう使用されないことを明示的に示すことができます。これはコードの可読性と保守性を向上させます。
エラーの早期検出:
Destroy()後にnilを代入しておくと、誤ってそのインスタンスを使用しようとした場合に即座にエラーが発生し、バグの早期発見につながります。

- **一貫性のある動作の保証**:
Destroy()の動作が環境によって異なる可能性があるため、nilを代入することで一貫した動作を保証できます。


## まとめ

今回は、`Destroy()`の仕様とnil代入の必要性についてまとめました。

作成したPartを削除するときは`Destroy()`だけではなくnilを代入するところまで実践しましょう。

https://devforum.roblox.com/t/destroy-and-nil-or-just-nil/2019750/22

https://create.roblox.com/docs/reference/engine/classes/Instance#Destroy