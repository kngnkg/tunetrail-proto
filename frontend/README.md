# Frontend

## 環境構築

開発コンテナに入ったら`create-next-app`でプロジェクトを作成する。既に存在するDocker関連のファイルとコンフリクトするため、一旦`tmp`に作成してから移動する。

```
npx create-next-app tmp --ts
```

```
✔ Would you like to use ESLint with this project? … Yes
✔ Would you like to use Tailwind CSS with this project? … No
✔ Would you like to use `src/` directory with this project? … Yes
✔ Would you like to use experimental `app/` directory with this project? … No
✔ What import alias would you like configured? … @/*
...
```
展開したファイル/ディレクトリを`/frontend`直下に移動する
```
mv tmp/* . && mv tmp/.* . && rm -r tmp
```

不要なファイルを削除する
```
rm .gitignore
```

ESLint/Prettierの設定 (ESLintは既にインストールされている)

```
npm install --save-dev prettier
```

`.prettierrc`を作成

lintを実行

```
npm run lint
```

`.eslintrc.json`が作成される。
