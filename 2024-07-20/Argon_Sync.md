# Argonを使ってRoblox StudioとVS Codeを同期する

今日はRoblox StudioというRobloxゲームの開発プラットフォームをVS Codeと同期できる便利なプラグインArgonの簡単なセットアップ方法をまとめ、実際にRoblox StudioとVS Codeを同期する手順を説明していこうと思います！

https://create.roblox.com/

## Argon

[公式ドキュメント](https://argon.wiki/docs/intro)に記載の通り、Roblox Studioを外部エディタやバージョン管理ツールなどと同期できるプラグインです。

> What is Argon?
It's a feature-full tool that allows use of external editors, version control systems, automated workflows and help manage both small and huge projects. It's made for professional users as well as for beginners!


Roblox Studio上のエディタではできることが限られており、バージョン管理も基本的にできません。そこでこのArgonを使って、Roblox Studioと自身のエディタを同期すれば、コーディングが快適になり、またバージョン管理もできるメリットがあります。

また、既存のエディタで使用したいという開発者にとっても非常に便利なプラグインになっています。

## Argonのインストール

今回はVS Codeで同期する方法を紹介します。

ArgonはRoblox StudioとVS Codeの両方にありますが、インストールはVS Codeからのみで大丈夫です。VS Codeからインストールすれば自動でRoblox Studio側にも追加されるためです。
手順も簡単で、VS CodeのエクステンションからArgonを検索して以下のものをインストールするだけです。

![](https://storage.googleapis.com/zenn-user-upload/843727a168d7-20240720.png)

VS Codeでインストールが完了すれば、Roblox Studio側で以下のようにArgonがインストールされているかと思います。

![](https://storage.googleapis.com/zenn-user-upload/ac276672f8b3-20240720.png)

## VS Code側での初期設定

次に、VS Code側で同期するディレクトリや設定ファイルを用意します。

手順は公式ドキュメントの内容で進めます。

https://argon.wiki/docs/getting-started/new-project#project-initialization

1. 任意のプロジェクトディレクトリ（私の場合は'Argon'）を作成してVS Codeで開きます。

2. `cmd`(winは `cntl`)+`shft`+`P`でVS Codeのメニューを開き、以下の'Argon: Open Menu`を選択します。

![](https://storage.googleapis.com/zenn-user-upload/e5d268a6f226-20240720.png)

3. 'Init'を選択します。

![](https://storage.googleapis.com/zenn-user-upload/ca53d0b64703-20240720.png)

4. その後のproject name、project template、project optionsはデフォルトの設定で進めます。

![](https://storage.googleapis.com/zenn-user-upload/7d0797b018a2-20240720.png)
*project name*

![](https://storage.googleapis.com/zenn-user-upload/1079754f2960-20240720.png)
*project template*

![](https://storage.googleapis.com/zenn-user-upload/2c651174f507-20240720.png)
*project options*

5. 完了するとプロジェクトディレクトリに以下のディレクトリと`default.project.json`などが作成されます。また、`default.project.json`に以下のような内容が記載されているのも確認できます。

![](https://storage.googleapis.com/zenn-user-upload/32dffe46ef5a-20240720.png)

なお、`default.project.json`ではRoblox StudioとVS Codeのディレクトリのマッピングなどを設定しています。詳細はドキュメントから確認できます。

https://argon.wiki/api/project#project-file


## Roblox StudioとVS Codeを同期する

1. VS Code側で、`cmd`(winは `cntl`)+`shft`+`P`でVS Codeのメニューを開き、'Argon: Open Menu`を開いて以下の'Serveを選択します。

![](https://storage.googleapis.com/zenn-user-upload/7e2aedfccc99-20240720.png)

2. 設定ファイルの選択に移るので、先ほど作成された`default.project.json`を選択します。

![](https://storage.googleapis.com/zenn-user-upload/33b13f2ef3bd-20240720.png)

3. 'serve options'はそのままの設定で進みます。

![](https://storage.googleapis.com/zenn-user-upload/eeae8b3eea9e-20240720.png)

これでVS Code側がServerとしてlocalhost:8000でServeした状態になります。

4. 次はRoblox Studio側でConnectします。Roblox StudioのPLUGINタブからArgonを選択すると以下のようなメッセージが表示されるので、'Connect'を選択します。

![](https://storage.googleapis.com/zenn-user-upload/29a58e363ee6-20240720.png)

5. 以下のように'Synced'と表示されれば接続成功です。

![](https://storage.googleapis.com/zenn-user-upload/52e4054312a7-20240720.png)

## 同期確認

これでRoblox StudioとVS Codeはリアルタイムで同期されたので、双方向の動機状態を確認してみましょう。

現時点でRoblox StudioとVS Codeのいずれの`ServerScriptService`にもファイルがないことを確認してください。

![](https://storage.googleapis.com/zenn-user-upload/b8bf5b33bfc7-20240720.png)
*Roblox Studio*

![](https://storage.googleapis.com/zenn-user-upload/3a611bff788c-20240720.png)
*VS Code*

### Roblox Studio → VS Codeの同期を確認

1. Roblox Studio側の`ServerScriptService`横にある+マークをおし、'Script'を選択してください。以下のようなスクリプトファイルが作成されると思います。

![](https://storage.googleapis.com/zenn-user-upload/8e72aa9ab2e4-20240720.png)

2. VS Code側で`ServerScriptService`を確認してください。以下のように`Script.server.luau`というファイル名でRoblox Studioのスクリプトと同じ内容のファイルが作成されているのが確認できます。

![](https://storage.googleapis.com/zenn-user-upload/42a2b5c5ea7f-20240720.png)

:::message
Robloxのスクリプトはluauというlua言語の拡張言語を使用しているため、エディタでは拡張子が.luauになっています。
:::

### VS Code → Roblox Studioの同期を確認

1. VS Code側で、先ほど作成された`Script.server.luau`を以下のように編集して保存します。

![](https://storage.googleapis.com/zenn-user-upload/126b43d63e7f-20240720.png)

2. Roblox Studio側で`ServerScriptService/Script`を確認し、以下のように変更内容が同期されているのが確認できます。

![](https://storage.googleapis.com/zenn-user-upload/37725b2ce348-20240720.png)

以上で、Argonを使ってRoblox StudioとVS Codeが同期できたので、スクリプトはVS Codeでコーディングできるようになりました。

## まとめ
Argonというプラグインを使ってRoblox StudioとVS Codeを同期していきました。

ドキュメントを参考に慣れると`project.json`をカスタマイズしてより高度な同期設定などもできるので、ぜひご活用ください！