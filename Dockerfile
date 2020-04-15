# ベースとなるDockerイメージ指定
FROM golang:latest

#ディレクトリ作成
WORKDIR /go/src/

#環境変数設定
# ENV 
# ENV 

#ローカルのデータをコンテナのディレクトリにコピー
COPY . /go/src/

#buildの実行
Run go build

#コンテナ実行時のデフォルトを設定
CMD ./wcafe-api-controller
