# RobloxでPlayerにステータスを追加する

Roblox Studioでプレイヤーにカスタマイズしたステータス（スタミナ、攻撃力、防御力など）を用意する方法を紹介します。

## プレイヤーにステータスを追加する手順

以下がプレイヤーに追加のステータスを実装するための手順です。

1. 追加するステータスとその初期値のテーブルを用意
2. プレイヤーに追加ステータスを登録する処理を実装
3. プレイヤーが新規でゲームに参加した時に初期値とともにステータスを登録

## ディレクトリ構成

今回は以下の場所にステータスを追加するLocalScriptを用意しましょう。

```text
workspace
├── Stamina (IntValue, Value = 100)
.
.
.
├── StarterPlayer
    ├── StarterCaracterScripts
    └── StarterPlayerScripts
        └── LocalScript // ここに用意
.
.
.
```

ちなみに'StarterPlayer'ディレクトリにはプレイヤー個々に対してクライアント側で実行したいスクリプトを置きます。
ここに置いたファイルはここで実行されるのではなく、**ゲーム開始時に各プレイヤー側にコピーされて実行される**ためです。

https://create.roblox.com/docs/reference/engine/classes/StarterPlayer

## 参加したプレイヤーにステータスを追加する

以上を踏まえて、'LocalScript'に、参加したプレイヤーに対してスタミナ、攻撃力、防御力のステータスを追加するコードを、先ほどの手順と照らし合わせながら以下のように実装してみました。

```lua
local Players = game:GetService("Players")

-- ステータスを追加する関数
local function addStatsToPlayer(player)
    -- ①各ステータスの名前と初期値を用意
    local stats = {
        {Name = "Stamina", Value = 100},
        {Name = "Attack", Value = 100},
        {Name = "Deffence", Value = 100},
    }

    -- ②プレイヤーに追加ステータスを登録する処理をforループで実装
    for _, stat in ipairs(stats) do
        local statValue = Instance.new("IntValue")  -- IntValueインスタンスを生成
        statValue.Name = stat.Name  -- ステータス名を設定
        statValue.Value = stat.Value  -- 初期値を設定
        statValue.Parent = player  -- プレイヤーにステータスを追加
        print("Added stat:", stat.Name, "to player:", player.Name)  -- デバッグ用ログ
    end
end

-- ③プレイヤーがゲームに参加したときにステータスを追加する処理
Players.PlayerAdded:Connect(function(player)
    addStatsToPlayer(player)
end)
```

ちなみに、この`IntValue`によって、参加プレイヤーには以下のように追加のステータスが実装されます。

```text
Player1
├── Stamina (IntValue, Value = 100)
├── Attack (IntValue, Value = 100)
└── Deffence (IntValue, Value = 100)
```

![](https://storage.googleapis.com/zenn-user-upload/103594023350-20240723.png)

プレイヤーのプロパティが増えていくイメージでいいかと思います。

## まとめ

以上のように新規プレイヤーに追加のカスタマイズしたステータスを実装する方法を紹介しました。

今回のコードはかなり自由度も拡張性も高い機能ですので、ぜひ真似していろいろ実装してみてください！