# ---------------------------------------------------
# リリース用のコンテナイメージを作成するステージ

FROM node:19.8 as deploy

WORKDIR /usr/src/app

COPY ./ ./

RUN npm install

RUN npm run build

CMD [ "npm", "start" ]

# ---------------------------------------------------
# 開発用のコンテナイメージを作成するステージ

FROM node:19.8 as dev

WORKDIR /workspace

COPY ./ ./

# 開発用のツールをインストール
RUN apt-get update && apt-get install -y \
        jq \
        iputils-ping \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN npm install -g npm@latest && npm install create-next-app

ENTRYPOINT ["sleep", "infinity"]
