/*
 * データベースのスキーマを定義する
 */

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/*
 * ユーザー
 */
CREATE TABLE users (
    id UUID,
    user_name VARCHAR(100) UNIQUE,
    name VARCHAR(100),
    icon_url VARCHAR(100),
    bio VARCHAR(1000),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

/*
 * フォロー
 */
CREATE TABLE follows (
    user_id UUID NOT NULL,
    followee_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT follows_pkey PRIMARY KEY (user_id, followee_id),
    CONSTRAINT follows_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT follows_followee_id_fkey FOREIGN KEY (followee_id)
        REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX follows_user_id_followee_id_idx ON follows (user_id, followee_id);

/*
 * 投稿
 */
CREATE TABLE posts (
    id UUID DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    body VARCHAR(1000),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT posts_pkey PRIMARY KEY (id),
    CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX posts_user_id_idx ON posts (user_id);

/*
 * リプライのツリー構造を表すテーブル
 * 削除された投稿についても、リプライのツリー構造だけは保持する
 */
CREATE TABLE reply_relations (
    post_id UUID NOT NULL, -- 削除された投稿のIDも保持したいので参照制約は設けない
    parent_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT reply_relations_pkey PRIMARY KEY (post_id, parent_id)
);

/*
 * いいね
 */
CREATE TABLE likes (
    post_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT likes_pkey PRIMARY KEY (post_id, user_id),
    CONSTRAINT likes_post_id_fkey FOREIGN KEY (post_id)
        REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT likes_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES users(id) ON DELETE CASCADE
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
