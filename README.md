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

users ||--o{ posts : "1人のユーザーは0以上の投稿を持つ"
users ||--o{ follows : "1人のユーザーは0以上のフォローを持つ"
posts ||--o| reply_relations : "リプライのツリー構造"
posts ||--o{ likes : "1つのポストは0以上のいいねを持つ"

users {
    UUID id PK
    string user_name
    string name
    string icon_url
    string bio
    timestamp created_at
    timestamp updated_at
}
deleted_users {
    UUID id PK
    string user_name
    string name
    string icon_url
    string bio
    timestamp created_at
    timestamp updated_at
    timestamp deleted_at
}
follows {
    UUID user_id PK,FK
    UUID followee_id PK,FK
    timestamp created_at
    timestamp updated_at
}
posts {
    UUID id PK
    UUID user_id
    string body
    timestamp created_at
    timestamp updated_at
}
reply_relations {
    UUID post_id PK,FK
    UUID parent_id PK,FK
    timestamp created_at
}
likes {
    UUID post_id PK,FK
    UUID user_id PK,FK
    timestamp created_at
    timestamp updated_at
}
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
