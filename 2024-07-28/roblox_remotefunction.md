# Roblox: クライアント-サーバー間の通信方法③(RemoteFunction)

前回に引き続き、Robloxでクライアント側のLocalScriptとサーバー側のScriptで情報をやり取りする際の方法についてご紹介していこうと思います！

全3回のアウトラインは以下のようになります。

1. [Client-Serverモデル: そもそものRobloxにおけるクライアントとサーバー間のルール](https://zenn.dev/kimd/articles/bbfe15872ab56a)
2. [RemoteEvent: クライアントとサーバー間の一方向の通信について](https://zenn.dev/kimd/articles/c7646e1a039293)
3. **RemoteFunction: クライアントとサーバー間の双方向の通信方法について←今回のテーマ**

前回はクライアントとサーバー間で一方向通信するためのイベント`RemoteEvent`を紹介をしました。

第3回の今回は、クライアントとサーバー間で双方向通信するためのイベント`RemoteFunction`を紹介していきます。

なお今回も前回に引き続き、以下のドキュメントの内容を元にまとめていこうと思います。

https://create.roblox.com/docs/scripting/events/remote

## `RemoteFunction`

第1回のおさらいですが、[`RemoteFunction`](https://create.roblox.com/docs/reference/engine/classes/RemoteFunction)とは、サーバーとクライアントの双方向の通信に使用され、送信側は応答を待機するイベントのことでした。

`RemoteFunction`を使用するときも`RemoteEvent`と同様、以下のようにReplicatedStorageに`RemoteFunction`クラスを用意し、両者このクラスを介して通信を実現します。

![](https://storage.googleapis.com/zenn-user-upload/a56d9fd3fe81-20240727.png)

ドキュメントでは双方向の通信として以下の2通りの例が挙げられています。ただし後者の実装には注意事項があります。（後述）

- [client→Server→Client](https://create.roblox.com/docs/scripting/events/remote#client-server-client)

- [Server→Client→Server](https://create.roblox.com/docs/scripting/events/remote#server-client-server)

### Client→Server→Client

クライアントからサーバーに通信して応答を受ける通信は、`RemoteFunction`の [`InvokeServer()`](https://create.roblox.com/docs/reference/engine/classes/RemoteFunction#InvokeServer)メソッドを呼び出すと、LocalScriptを使用してサーバー上の関数を呼び出すことができます。 リモートイベントとは異なり`RemoteFunction`を呼び出す LocalScriptは、コールバックが返されるまで待機します。[`InvokeServer()`](https://create.roblox.com/docs/reference/engine/classes/RemoteFunction#InvokeServer)に渡す引数は、サーバー側のScriptで定義する`RemoteFunction`の[`OnServerInvoke`](https://create.roblox.com/docs/reference/engine/classes/RemoteFunction#OnServerInvoke)コールバックに渡されます。この引数に渡せる値には特定の制限がありますが今回は割愛します。[^1]

[^1]: https://create.roblox.com/docs/scripting/events/remote#argument-limitations

- Client: `RemoteFunction:InvokeServer(args)`
- Server: `RemoteFunction.OnServerInvoke = function(player, args)`

ドキュメントの、クライアントから色とサイズを引数に入れてサーバー側のPartを新規作成する関数を実行するよう要求し、サーバーは作成したPartをクライアント側に返すといったサンプルコードを見てみましょう。

トリガーとなるクライアント側のLocalScriptは以下のように書きます。

- Client

```lua
local ReplicatedStorage = game:GetService("ReplicatedStorage")

-- リモートファンクションインスタンスを取得
local remoteFunction = ReplicatedStorage:FindFirstChildOfClass("RemoteFunction")

-- colorとpositionをサーバー側のOnInvokeに設定されている関数に渡してその返り値を待つ
local newPart = remoteFunction:InvokeServer(Color3.fromRGB(255, 0, 0), Vector3.new(0, 25, -20))

-- サーバーから帰ってきたnewPartを表示
print("The server created the requested part:", newPart)
```

続いてこれを受けるサーバー側のScriptは以下のように書きます。

- Server

```lua
local ReplicatedStorage = game:GetService("ReplicatedStorage")

-- リモートファンクションインスタンスを取得
local remoteFunction = ReplicatedStorage:FindFirstChildOfClass("RemoteFunction")

-- コールバック関数に設定する関数
local function createPart(player, partColor, partPosition)
	print(player.Name .. " requested a new part")
	local newPart = Instance.new("Part")
	newPart.Color = partColor
	newPart.Position = partPosition
	newPart.Parent = workspace
	return newPart -- newPartをクライアントに返す
end

-- createPartをリモートファンクションのコールバック関数に設定
remoteFunction.OnServerInvoke = createPart
```

### Server→Client→Server

続いてサーバーからクライアントに通信して応答を受ける通信ですが、ドキュメント内では以下のように注意書きがされています。

> You can use a Script to call a function on the client by calling the InvokeClient() method on a RemoteFunction, but it has serious risks as follows:
>
> - If the client throws an error, the server throws the error too.
> - If the client disconnects while it's being invoked, InvokeClient() throws an error.
> - If the client doesn't return a value, the server yields forever.
>
> For actions that don't require two-way communications, such as updating a GUI, use a RemoteEvent and communicate from server to client.

簡単に説明すると、以下の3点の危険性から、**ServerからClientへの通信は基本的には応答を待たず、`RemoteEvent`での一方向の通信をするのが望ましい**とのことです。

- クライアントがエラーを返すと、サーバーもエラーになり予期せぬサーバー停止に繋がりかねない
- クライアントの呼び出し中にクライアントが切断されると、InvokeClient()はエラーを返すので、この場合もサーバー側でエラーになってしまう
- クライアントが値を返さない場合、サーバーは永遠に応答を待ち続けてしまい、次の処理が実行されなくなる

これらの理由から、ドキュメントには実装のサンプルが用意されていませんでした。

Server→Client→Serverの双方向通信を実装する方法自体は、先ほどのClient→Server→Clientの逆で、`RemoteFunction`の [`InvokeClient()`](https://create.roblox.com/docs/reference/engine/classes/RemoteFunction#InvokeClient)メソッドを呼び出すと、Scriptを使用してクライアント上の関数を呼び出すことができます。そして、`RemoteFunction`を呼び出す Scriptは、コールバックが返されるまで待機します。`InvokeClient()`に渡す引数は、クライアント側のLocalScriptで定義する`RemoteFunction`の[`OnClientInvoke`](https://create.roblox.com/docs/reference/engine/classes/RemoteFunction#OnClientInvoke)コールバックに渡されます。

## まとめ

クライアントとサーバー間で双方向通信をするための`RemoteFunction`を紹介しました。

今回で3回に渡って紹介したクライアント-サーバー間の通信方法の記事は以上です。

`RemoteEvent`と`RemoteFunction`、それぞれのケースの特徴ををよく理解して実装して行くことが大切だと思いました。