# ---------------------------------------------------
# リリース用のコンテナイメージを作成するステージ

FROM node:19.8 as deploy

WORKDIR /usr/src/app

RUN apt-get update

COPY ./ ./

RUN npm install

ARG NEXT_PUBLIC_API_ROOT
ENV NEXT_PUBLIC_API_ROOT=$NEXT_PUBLIC_API_ROOT

RUN npm run build

CMD [ "npm", "start" ]

# ---------------------------------------------------
# 開発用のコンテナイメージを作成するステージ

FROM node:19.8 as dev

WORKDIR /workspace

RUN apt-get update

COPY ./ ./

RUN npm install -g npm@latest && npm install create-next-app

# 開発に使用するコマンドをインストール
RUN apt -y install jq && \
    apt-get -y install iputils-ping net-tools

ENTRYPOINT ["sleep", "infinity"]
