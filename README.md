# wcafe Rest Sample

## これなに

ginのRestサーバーお試し版

## リポジトリクローン

```
cd $GOPATH/src/github.com
mkdir nfv-aws
cd nfv-aws
git clone git@github.com:nfv-aws/wcafe-api-controller.git
```

## 使い方

### パッケージインストール

```
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql
go get github.com/google/uuid
```

### DB初期設定

bashrcとかに以下を追記

```
vi ~/.bashrc

export WCAFE_DB_PASSWORD=password
export WCAFE_DB_ENDPOINT=endpoint

source ~/.bashrc
```


ユーザーやDB名は以下でも編集可能

```
vi config/config.toml
```

## 動作確認

```
go run main.go

curl localhost:8080/api/v1/pets
```

## UnitTest

### Controller層(DBはMock利用)

```
go test -v ./controller/...

PASS
ok      github.com/nfv-aws/wcafe-api-controller/server  0.206s
```

### Service層(DBないとうごきません)

```
go test -v ./server/...

PASS
ok      github.com/nfv-aws/wcafe-api-controller/server  0.206s
```

## コンテナ上で動作確認
イメージ作成
```
docker build --build-arg DB_PASS=$(echo $WCAFE_DB_PASSWORD) --build-arg DB_ENDPOINT=$(echo $WCAFE_DB_ENDPOINT) -t wcafe .
```

コンテナの生成と実行
```
docker run -d -p 8080:8080 wcafe
```
確認
```
curl localhost:8080/api/v1/pets
```

## 参考

https://qiita.com/Asuforce/items/0bde8cabb30ac094fcb4
https://qiita.com/hiroyky/items/4a9be463e752d5c0c41c

## Tips

### Mockの作り方

```
mockgen -source service/pets_service.go -destination mocks/pets_service.go -package mocks
mockgen -source service/stores_service.go -destination mocks/stores_service.go -package mocks
mockgen -source service/users_service.go -destination mocks/users_service.go -package mocks
```
