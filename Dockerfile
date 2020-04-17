# ベースとなるDockerイメージ指定
FROM golang:latest

#ディレクトリ作成
WORKDIR /go/src/

## 引数設定
ARG DB_ENDPOINT
ARG DB_PASS

#環境変数設定
ENV WCAFE_DB_PASSWORD=$DB_PASS
ENV WCAFE_DB_ENDPOINT=$DB_ENDPOINT

#ローカルのデータをコンテナのディレクトリにコピー
COPY . /go/src/

#buildの実行
RUN go build

#コンテナ実行時のデフォルトを設定
CMD ./wcafe-api-controller
