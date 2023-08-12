# ---------------------------------------------------
# ローカル実行用のコンテナイメージを作成するステージ

FROM golang:1.20.3-alpine as run-local

ENV GO111MODULE on

WORKDIR /workspace/reverse-proxy-for-dev

CMD ["go", "run", "main.go"]

# ---------------------------------------------------
# 開発時に利用するコンテナイメージを作成するステージ

FROM golang:1.20.3 as dev

ENV GO111MODULE on

WORKDIR /workspace/reverse-proxy-for-dev

# goの各種ツールをインストール
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/stamblerre/gocode@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/k0kubun/sqldef/cmd/psqldef@latest && \
    go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
