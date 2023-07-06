/*
 * データベースのスキーマを定義する
 */

/*
 * ユーザー
 */
CREATE TABLE users (
    id serial PRIMARY KEY,
    user_name VARCHAR(100) UNIQUE,
    name VARCHAR(100),
    password VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    icon_url VARCHAR(100),
    bio VARCHAR(1000),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

/*
 * 投稿
 */
CREATE TABLE posts (
    id serial PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    body VARCHAR(1000),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

/*
 * 投稿に添付される画像
 * 1つの投稿に複数の画像を添付できる
 */
CREATE TABLE post_images (
    id serial PRIMARY KEY,
    post_id INTEGER NOT NULL REFERENCES posts(id),
    image_url VARCHAR(100)
);

/*
 * リプライ
 * 投稿を継承している
 * ある投稿に対して、別の投稿がリプライされたことを表す
 */
CREATE TABLE replies (
    post_id INTEGER PRIMARY KEY REFERENCES posts(id),
    dest_post_id INTEGER NOT NULL REFERENCES posts(id)
);

/*
 * リプライにメンションされたユーザー
 */
CREATE TABLE reply_destination_users (
    post_id INTEGER NOT NULL REFERENCES posts(id),
    dest_user_id INTEGER NOT NULL REFERENCES posts(id),
    PRIMARY KEY (post_id, dest_user_id)
);

/*
 * いいね
 */
CREATE TABLE likes (
    user_id INTEGER NOT NULL REFERENCES users(id),
    post_id INTEGER NOT NULL REFERENCES posts(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id, post_id)
);

/*
 * タグ
 */
CREATE TABLE tags (
    id serial PRIMARY KEY,
    name VARCHAR(100) UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

/*
 * 投稿とタグの関連付け
 */
CREATE TABLE post_tag (
    post_id INTEGER NOT NULL REFERENCES posts(id),
    tag_id INTEGER NOT NULL REFERENCES tags(id),
    PRIMARY KEY (post_id, tag_id)
);
