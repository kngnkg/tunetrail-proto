# リリース用のビルドを行うコンテナイメージを作成するステージ
FROM golang:1.20.3-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags "-w -s" -o app

# ---------------------------------------------------
# リリース用のコンテナイメージを作成するステージ

FROM debian:bullseye-slim as deploy

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ---------------------------------------------------
# ローカル環境でのテスト用のコンテナイメージを作成するステージ
FROM golang:1.20.3-bullseye as test

WORKDIR /

# RIEの設定
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/local/bin/aws-lambda-rie
RUN chmod 755 /usr/local/bin/aws-lambda-rie

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags "-w -s" -o app

ENTRYPOINT [ "/usr/local/bin/aws-lambda-rie" ]
CMD ["./app"]

# ---------------------------------------------------
# 開発時に利用するコンテナイメージを作成するステージ

FROM golang:1.20.3 as dev

WORKDIR /

# RIEの設定
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/local/bin/aws-lambda-rie
RUN chmod 755 /usr/local/bin/aws-lambda-rie

# ワーキングディレクトリを変更
WORKDIR /workspace

# goの各種ツールをインストール
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/stamblerre/gocode@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/k0kubun/sqldef/cmd/psqldef@latest

ENTRYPOINT ["sleep", "infinity"]
