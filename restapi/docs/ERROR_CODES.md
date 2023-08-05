# Error Codes

このドキュメントでは、APIが返す可能性があるエラーコードとその意味について説明します。

## 汎用的なエラー

| Error Code | HTTP status code | Message                | Description                                   |
|------------|------------------|------------------------|-----------------------------------------------|
| `4000`       | `400 Bad Request`              | Bad request            | その他の不正なリクエストの場合に返されます。 |
| `4001`       | `400 Bad Request`              | Invalid parameter      | ユーザーからの入力が不正な場合に返されます。具体的には、必要なパラメータが欠けている、またはパラメータの形式が正しくない場合に発生します。 |

## 認証に関連するエラー

| Error Code | HTTP status code | Message                | Description                                   |
|------------|------------------|------------------------|-----------------------------------------------|
| `4101`       | `400 Bad Request`              | Invalid confirmation code | 確認コードが一致しない場合に返されます。 |
| `4102`       | `400 Bad Request`              | Confirmation code expired | 確認コードが期限切れの場合に返されます。 |
| `4103`       | `409 Conflict`              | Email already confirmed | メールアドレスが既に認証済みの場合に返されます。 |
| `4104`       | `400 Bad Request`              | Wrong email or password | メールアドレスまたはパスワードが一致しない場合に返されます。 |
| `4105`       | `400 Bad Request`              | Token expired | トークンが期限切れの場合に返されます。 |

## ユーザーに関連するエラー

| Error Code | HTTP status code | Message                | Description                                   |
|------------|------------------|------------------------|-----------------------------------------------|
| `4201`       | `404 Not Found`              | User not found         | ユーザーが見つからない場合に返されます。 |
| `4202`       | `409 Conflict`              | UserName already entry | ユーザーネームが既に登録されている場合に返されます。 |
| `4203`       | `409 Conflict`              | Email already entry    | メールアドレスが既に登録されている場合に返されます。 |

## その他サーバーエラー

| Error Code | HTTP status code | Message                | Description                                   |
|------------|------------------|------------------------|-----------------------------------------------|
| `5000`       | `500 Internal Server Error`              | Unknown server error   | サーバー内部で不明なエラーが発生した場合に返されます。 |
