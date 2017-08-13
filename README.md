# Timoney [ティモニー] 

## 時は金なり

会計情報を管理するためのツールです。  
GolangのフレームワークであるRevelを利用しています。

## 前提

以下がインストールされていること

- Golang

## セットアップ

- Revelのコマンドを取得

```sh
$ go get -u github.com/revel/cmd/revel
```

- depの取得

```sh
$ go get -u github.com/golang/dep/cmd/dep
```

- depの実行

```sh
$ dep ensure
```

## 実行

```sh
$ revel run timoney
```
