# Deviation-value

## Command

このアプリを動かすためのコマンド集です。

#### Dockerイメージを作成後、アプリケーション起動

```makefile
make
```

#### アプリケーション起動
```makefile
make up
```

#### Dockerイメージ作成

```makefile
make build
```

またサーバ(Go)側とDB側の別々に実行することもで可能です。
```makefile
Go : make go
DB : make posgre
```

