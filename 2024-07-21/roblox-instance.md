# Roblox StudioのInstanceとは？

Robloxにおいて、ありとあらゆるところで出現する`Instance`について簡単にまとめます。

## 結論：`Instance`=全オブジェクト

そうです。Robloxでは全ての **オブジェクトが`Instance`**qです。

luauでは、置物や障害物などの物体を表す`Part`やそれらを制御するための`Script`、またGUIの要素など全てを総称して`Instance`型としちゃっています。

[公式ドキュメント](https://create.roblox.com/docs/reference/engine/datatypes/Instance)でもシンプルに以下の記載のみです。

>The Instance data type holds the constructor for Instance objects.

Robloxではディレクトリの階層の関係を親子関係で表しますが、その親も子も、彼らのステータスや特技も何もかも`Instance`といったイメージです。

## 具体例

`Instance`の代表例である`Part`と`Script`の具体例を挙げます。といっても、これら自身が`Instance`なのでオブジェクト自体の解説コードになってしまいますが...

```lua
-- sample of 'Part'

local part = Instance.new("Part")
part.Name = "MyPart"  -- オブジェクトの名前を設定
part.Position = Vector3.new(0, 10, 0)  -- 位置を設定
part.Parent = workspace  -- 親オブジェクトを設定（workspaceに追加）
```

```lua
-- sample of 'Script'

local script = Instance.new("Script")
script.Name = "MyScript"
script.Source = "print('Hello, world!')"  -- スクリプトのソースコードを設定
script.Parent = workspace  -- 親オブジェクトを設定（workspaceに追加）
```

このように、すべてが`Instance`なので、`Part`やら`Script`やらを区別せず同じように扱えたりメソッドが使えたりといったメリットがあります。

ですがその反面、すべて同じであるがゆえに、不必要なプロパティが付与されたり特定の要件に対するカスタマイズ性にかけたりするデメリットもあります。

ですので、設計する際などは特に注意が入りそうですね。。

## まとめ

**`Instance`は全てのオブジェクト**です。これだけ理解できれば問題ないです！
