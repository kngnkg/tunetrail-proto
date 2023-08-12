# リリース用のビルドを行うコンテナイメージを作成するステージ

FROM node:19.8 as deploy-builder

WORKDIR /app

COPY ./ ./

RUN npm install

RUN npm run build

CMD [ "npm", "start" ]

# ---------------------------------------------------
# リリース用のコンテナイメージを作成するステージ

FROM node:19.8-slim as deploy

ENV NODE_ENV=production

WORKDIR /app

COPY --from=deploy-builder /app/next.config.js ./
COPY --from=deploy-builder /app/public ./public
COPY --from=deploy-builder /app/.next/static ./.next/static

COPY --from=deploy-builder /app/.next/standalone ./

CMD ["node", "server.js"]

# ---------------------------------------------------
# 開発用のコンテナイメージを作成するステージ

FROM node:19.8 as dev

WORKDIR /workspace/webapp

COPY ./ ./

# 証明書をインストール
COPY ../reverse-proxy-for-dev/localhost.pem /usr/local/share/ca-certificates/localhost.crt
RUN update-ca-certificates

# 開発用のツールをインストール
RUN apt-get update && apt-get install -y \
    jq \
    iputils-ping \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN npm install -g npm@latest

CMD ["npm", "run", "dev"]
