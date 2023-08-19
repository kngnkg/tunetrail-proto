# TuneTrail

## 使用技術

| 技術       | 詳細                         |
| :--------- | :--------------------------- |
| Next.js    | フロントエンドフレームワーク |
| TypeScript |                              |


バックエンド
| 技術       | 詳細              |
| :--------- | :---------------- |
| Go 1.20.3  | APIサーバー       |
| Gin        | Webフレームワーク |
| PostgreSQL | RDB               |

インフラ
| 技術          | 詳細                                                                     |
| :------------ | :----------------------------------------------------------------------- |
| ECS (Fargate) | コンテナの管理                                                           |
| Route 53      | 独自ドメインを登録する目的                                               |
| S3            | 画像の保存                                                               |
| Docker        | リリース時のコンテナイメージを軽量にする目的でマルチステージビルドを採用 |
| Terraform     | インフラのコード化                                                       |

## ER図

```mermaid
erDiagram

users {
    UUID id PK
    string user_name "ユーザーネーム"
    string name "名前"
    string icon_url "アイコンのURL"
    string bio "プロフィール"
    bool is_deleted "削除済みかどうか"
    timestamp created_at
    timestamp updated_at
}

follows {
    UUID user_id PK,FK "フォローしたユーザーのID"
    UUID followee_id PK,FK "フォローされたユーザーのID"
    timestamp created_at
    timestamp updated_at
}

posts {
    int id PK
    UUID user_id "投稿したユーザーのID"
    string body "本文"
    timestamp created_at
    timestamp updated_at
}


users ||--o{ posts : "1人のユーザーは0以上の投稿を持つ"
users ||--o{ follows : "1人のユーザーは0以上のフォロイーを持つ"
follows }o--|| users : "1人のユーザーは0以上のフォロワーを持つ"

```

## 環境構築

### リバースプロキシの公開鍵と秘密鍵を作成

[mkcert](https://github.com/FiloSottile/mkcert)を使用します。

```console
mkcert -install
```

```console
cd reverse-proxy-for-dev && \
mkcert -cert-file ./localhost.pem -key-file ./localhost-key.pem localhost "host.docker.internal" "127.0.0.1"
```

### 各コンテナを起動

```console
docker compose up
```

### REST API

VSCodeのコマンドパレットで**Dev Containers: Open Folder in Container**を実行し、`restapi`ディレクトリを選択します。

### webapp

VSCodeのコマンドパレットで**Dev Containers: Open Folder in Container**を実行し、`webapp`ディレクトリを選択します。
