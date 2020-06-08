# ベースとなるDockerイメージ指定
FROM golang:latest

#作業ディレクトリ作成
WORKDIR /go/src/

#ホストのデータを作業ディレクトリにコピー
COPY . /go/src/

#go buildの実行
RUN go build

#コンテナ生成時の命令を指定
CMD /go/src/wcafe-api-controller
