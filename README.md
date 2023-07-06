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
    int id PK
    string user_name "ユーザーネーム"
    string name "名前"
    string password "パスワード"
    string email "メールアドレス"
    string icon_url "アイコンのURL"
    string bio "プロフィール"
    timestamp created_at
    timestamp updated_at
}

posts {
    int id PK
    int user_id "投稿したユーザーのID"
    string body "本文"
    timestamp created_at
    timestamp updated_at
}

post_images {
    int id PK
    int post_id FK
    string url "画像のURL"
}

replies {
    int post_id PK,FK "継承元の投稿のID"
    int dest_post_id PK,FK "宛先の投稿のID"
}

reply_destination_users {
    int post_id PK,FK "リプライが継承している投稿のID"
    int dest_user_id "宛先のユーザーのID"
}

likes {
    int post_id PK,FK "投稿のID"
    int user_id PK,FK "いいねしたユーザーのID"
    timestamp created_at
    timestamp updated_at
}

post_tag {
    int post_id PK,FK
    int tag_id PK,FK
}

tags {
    int id PK
    string name "タグ名"
    timestamp created_at
    timestamp updated_at
}

users ||--o{ posts : "1人のユーザーは0以上の投稿を持つ"
posts ||--o{ post_images: "1つの投稿は0以上の画像を持つ"
posts ||--o{ replies : "1つの投稿は0以上のリプライを持つ"
posts ||--|| replies : "1つのリプライは1つの投稿を継承する"
posts ||--o{ likes : "1つの投稿は0以上のいいねを持つ"
posts ||--o{ post_tag : "1つの投稿は0以上の`post_tag`を持つ"
tags ||--o{ post_tag : "1つのタグは0以上の`post_tag`を持つ"
replies ||--o{ reply_destination_users: "1つのリプライは1以上の`reply_destination_users`を持つ"
users ||--o{ reply_destination_users: "1人のユーザーは0以上の`reply_destination_users`を持つ"

```

## 環境構築

externalなネットワークを作成します。

```
docker network create tunetrail-external
```

docker-composeで各コンテナを起動します。

```
docker compose up
```
### api

VSCodeのコマンドパレットで**Dev Containers: Open Folder in Container**を実行し、`api`ディレクトリを選択します。

### webapp

VSCodeのコマンドパレットで**Dev Containers: Open Folder in Container**を実行し、`webapp`ディレクトリを選択します。
