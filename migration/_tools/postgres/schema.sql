/*
 * データベースのスキーマを定義する
 */

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/*
 * ユーザー
 */
CREATE TABLE users (
    id UUID PRIMARY KEY,
    user_name VARCHAR(100) UNIQUE,
    name VARCHAR(100),
    icon_url VARCHAR(100),
    bio VARCHAR(1000),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

/*
 * フォロー
 */
CREATE TABLE follows (
    user_id UUID NOT NULL REFERENCES users(id),
    followee_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id, followee_id)
);

/*
 * 投稿
 */
CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    body VARCHAR(1000),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

/*
 * リプライのツリー構造を表すテーブル
 * ある投稿に対して、別の投稿がリプライされたことを表す
 * 削除された投稿に対するリプライはアプリケーション側で制御する
 */
CREATE TABLE reply_relations (
    post_id UUID NOT NULL,
    parent_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY (post_id, parent_id)
);

/*
 * いいね
 */
CREATE TABLE likes (
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id, post_id)
);

-- /*
--  * リプライにメンションされたユーザー
--  */
-- CREATE TABLE reply_destination_users (
--     post_id UUID NOT NULL REFERENCES posts(id),
--     dest_user_id UUID NOT NULL REFERENCES posts(id),
--     PRIMARY KEY (post_id, dest_user_id)
-- );

-- /*
--  * 投稿に添付される画像
--  * 1つの投稿に複数の画像を添付できる
--  */
-- CREATE TABLE post_images (
--     id serial PRIMARY KEY,
--     post_id UUID NOT NULL REFERENCES posts(id),
--     image_url VARCHAR(100)
-- );

-- /*
--  * タグ
--  */
-- CREATE TABLE tags (
--     id serial PRIMARY KEY,
--     name VARCHAR(100) UNIQUE,
--     created_at TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP NOT NULL
-- );

-- /*
--  * 投稿とタグの関連付け
--  */
-- CREATE TABLE post_tag (
--     post_id UUID NOT NULL REFERENCES posts(id),
--     tag_id UUID NOT NULL REFERENCES tags(id),
--     PRIMARY KEY (post_id, tag_id)
-- );
