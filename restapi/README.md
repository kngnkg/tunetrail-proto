# API

## エンドポイント

| HTTPメソッド  | パス              | 概要                                   |
| :----------- | :--------------- | :------------------------------------- |
| GET          | `/health`        | 疎通を確認する                         |
| POST         | `/auth/register` | ユーザーを登録する                     |
| PUT          | `/auth/confirm`  | メールアドレスを検証する                     |
| POST         | `/auth/signin`   | サインインする                     |
| POST         | `/auth/refresh`  | アクセストークンをリフレッシュする           |
| GET          | `/user/hoge`     | ユーザー名が`hoge`のユーザーを取得する |
| PUT          | `/user`          | ユーザーを更新する                     |
| DELETE       | `/user/hoge`     | ユーザー名が`hoge`のユーザーを削除する |

## エラーレスポンス

エラーレスポンスは以下のようなJSON形式で返されます。

```json
{
    "code": 4203,
    "developerMessage": "Email already entry",
    "userMessage": "登録できないメールアドレスです。"
}
```

詳細は[ERROR_CODES.md](docs/ERROR_CODES.md)を参照してください。

## 環境構築
### Go Modulesの初期化

```
go mod init github.com/kngnkg/tunetrail/api
```
