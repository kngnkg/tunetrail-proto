# API

## エンドポイント

| HTTPメソッド | パス             | 概要                                   |
| :----------- | :--------------- | :------------------------------------- |
| GET          | `/health`        | 疎通を確認する                         |
| POST         | `/user` | ユーザーを登録する                     |
| GET          | `/user/hoge`     | ユーザー名が`hoge`のユーザーを取得する |
| PUT          | `/user`   | ユーザーを更新する                     |
| DELETE       | `/user/hoge`     | ユーザー名が`hoge`のユーザーを削除する |

## 環境構築
### Go Modulesの初期化

```
go mod init github.com/kngnkg/tunetrail/api
```
