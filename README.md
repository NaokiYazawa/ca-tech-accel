## 環境構築

```sh
$ go mod tidy
$ go run main.go

# 別ターミナル
$ curl http://localhost:8080/user/list
```
## Todo
- `/ping` で `pong` を返すようにしよう
- ユーザ名の更新処理を追加しよう
- 排他制御に対応しよう *
- mysql を導入しよう ([database/sql](https://pkg.go.dev/database/sql) / [jmoiron/sqlx](https://github.com/jmoiron/sqlx) / [go-gorm/gorm](https://github.com/go-gorm/gorm))
- [echo](https://echo.labstack.com/)をサーバー構築に使ってみる
- docker で 開発環境を作ってみよう (multi stage build)
- アーキテクチャを変更する (DDD, MVC, クリーンアーキテクチャ, レイヤードアーキテクチャ)
- テストを書いてみる (TDT, モック)
- linter を導入する (golangci-lint) *
- 継続的にコードの品質を検証する (github actions)


- パス設計を考察する (REST) *
- graceful shutdown に対応する **
- Grpc に置き換える **
- デプロイしてみる (CD)　*
  - k8s について学ぶ **
