# Deviation-value

## Deviation valueとは？？
■開発詳細

模試などで提供される，自分の点数，偏差値，平均点を入力して，設定した偏差値を取るためには何点取ればよいのかと，設定した点数に対応した偏差値を表示してくれるWebアプリを作成しました．
   
■機能について
   
２つ機能があります。
   
１つ目は、模試などで提供される，自分の点数，偏差値，平均点を入力して，設定した偏差値を取るためには何点取ればよいのか，あるいは、設定した点数に対応した偏差値のどちらかを選択して表示
   
２つ目は、表示された偏差値と都道府県を選択すると、選択した偏差値と同じ高校名を表示


## Dirctory

ディレクトリの役割について簡単に紹介します。

### templates

webページを形成する雛形のhtmlファイルがあります。

### assets

基本的にフロントエンド関係のファイル(html, css, js)が入っています。


### controller

Mこのディレクトリではクライアントから受け取ったリクエストを適切なメソッドへ渡しています。


### db

DB (PostgreSQL) への接続や、データを抜き出すような処理を行っています。

### util

汎用的な処理（ポート番号取得等）をまとめています。

### docker

Dockerイメージを作成する際に必要となるファイル等をまとめています。

### Godeps, vendor

Herokuへデプロするために用意したディレクトリになります。

Dockerで動作できるようになったので、なくそうと考えています。

## Command

このアプリを動かすためのコマンド集です。

#### Dockerイメージを作成後、アプリケーション起動

```makefile
make
```

#### アプリケーション起動

このコマンドを最初に実行すると、Docker Hubからイメージをpullして実行されます
```makefile
make up
```

#### Dockerイメージ(Go, DB)作成

```makefile
make build
```

Docker Hubからpullしてサーバ(Go)側とDB側を別々に実行するには

```makefile
Go : make go
DB : make posgre
```


またサーバ(Go)側とDB側の別々にDockerイメージを作成することができます。
```makefile
Go : make build_go
DB : make build_posgre
```

