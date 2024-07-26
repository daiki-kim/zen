# Roblox: 新規プレイヤー参加時に実行される処理をクライアント側で書く方法

Robloxで参加してきたプレイヤー対して実行する関数を、クライアント側で書く方法を研修で教えてもらったのでまとめます！

## `PlayerAdded`

まず、Robloxには便利なイベントやメソッドがいっぱいあります。その中でもプレイヤーが参加した時にそのプレイヤーに対して何か実行する際には`PlayerAdded`というイベントが使えます。

:::message
ドキュメントでは実行や発生を'fire'と表現していて、この表現が気に入っているので本記事もファイヤーを使っていきますね。
:::

`PlayerAdded` : **プレイヤーがゲームに参加した瞬間にファイヤー(発生)するイベント。**

>PlayerAdded
The PlayerAdded event fires when a player enters the game. This is used to fire an event when a player joins a game, such as loading the player's saved GlobalDataStore data.

https://create.roblox.com/docs/reference/engine/classes/Players#PlayerAdded

ドキュメントでは以下のように使用しています。

```lua
local Players = game:GetService("Players")

-- playerが参加した時にConnect()の引数に入れたfuncton()をファイヤー(実行)するイベント
Players.PlayerAdded:Connect(function(player)
    print(player.Name .. "joined the game!")
```

## `PlayerAdded`はServer側で使う

大前提として、Robloxではクライアント側とサーバー側でスクリプトを保存するフォルダが明確に分かれています。この話はまた次回まとめようと思いますので、今回はクライアントとサーバーで処理を分けられるくらいの認識があればOKです。

私はクライアント側で、参加したプレイヤーに対してカスタムステータスを初期化してプレイヤーの子インスタンスにする処理をこの`PlayerAdd`を使って実装していましたが、いざテストプレイするとこの`PlayerAdded`がファイヤーしてくれませんでした。

ここで注意したいのが、`PlayerAdded`イベントは、**プレイヤーが参加した瞬間にファイヤーする**ということです。

それに対してクライアント側で実行するスクリプトというのはプレイヤーが参加した後に実行されます。厳密にいうと、プレイヤーが参加するとサーバー側で保管しているクライアント側のスクリプトをそのプレイヤーのローカルにコピーしてから実行しています。

ですので、クライアント側のスクリプトで`PlayerAdded`を用意しても、そのスクリプトが作成されて実行される段階ではプレイヤーはすでにゲーム内なのでファイヤーのタイミングを逃してしまっています。

このことから`PlayerAdded`は、サーバー側でプレイヤーが参加したタイミングで実行したい処理に対して使用するのが基本になってきます。

## クライアント側では`init()`を定義して代替しよう

それでも、クライアント側でプレイヤー参加時にファイヤーする処理を書く必要が出てくると思います。

そんな時は以下のように書けばいいと上司から教わりました。

```lua
local Players = game:GetService("Players")

-- playerが参加した時にファイヤー(実行)する関数
local init()
    player = Players.LocalPlayer
    print(player.Name .. "joined the game!")
end

-- スクリプトが作成されたときに実行される
init()
```

GoやPythonを書いてる時は割とよく使う方法でしたが、Robloxになっていろんなメソッドがあるので、普通にこのように書く方法が思いつきませんでした。

## init()を使う注意点

`init()`の多様はやめましょう。開発規模が大きくなるとあちらこちらに`init()`が出てきて、どの`init()`から実行しているか意味不明になりバグの原因にもなりかねません。

各所に初期化処理を実装したい場合は、`InitStatus()`のようにグローバルにユニークな命名をしてあげて、初期化関数のみをまとめて実行するスクリプトを一つ用意するのがいいと教えて頂きました。

## まとめ

今回は便利なプレイヤー参加時にファイヤーできる便利な`PlayerAdded`の紹介とその保存先の注意点、そしてクライアント側で同様な処理を実装する方法について紹介しました。

Robloxではゲーム制作に便利な機能が多い一方で、どれをどのタイミングで使うべきかがわからなくなってしまいやすいです。

今後はそれぞれの便利な機能とその使い方なんかも小出しで紹介できればと思っております。