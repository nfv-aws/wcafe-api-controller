# ベースとなるDockerイメージ指定
FROM golang:latest

#ディレクトリ作成
WORKDIR /go/src/

#環境変数設定
# ENV [WCAFE_DATABASE_PASSWORD]
# ENV [WCAFE_DATABASE_ENDPOINT]

#ローカルのデータをコンテナのディレクトリにコピー
COPY . /go/src/

#buildの実行
RUN go build

#コンテナ実行時のデフォルトを設定
CMD ./wcafe-api-controller
