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

## 事前準備
* RDSでMysql5.7.22のDBを作成してエンドポイントが分かる状態にしておく
* SQSで標準キューのQueueを2つ作成してエンドポイントが分かる状態にしておく

### パッケージインストール

```
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql
go get github.com/google/uuid
github.com/aws/aws-sdk-go/aws
github.com/aws/aws-sdk-go/aws/session
github.com/aws/aws-sdk-go/service/sqs
github.com/guregu/dynamo
```

### 環境設定

DB設定とAWSのSQS操作用の設定を追加

```
vi ~/.bashrc

export WCAFE_DB_PASSWORD=password
export WCAFE_DB_ENDPOINT=endpoint
export WCAFE_SQS_REGION=region
export WCAFE_SQS_PETS_QUEUE_URL=queue_url_1
export WCAFE_SQS_STORES_QUEUE_URL=queue_url_2
export WCAFE_SQS_USERS_QUEUE_URL=queue_url_3
export WCAFE_LOG_FILE_PATH=file_path  
export WCAFE_DYNAMODB_REGION=region
export WCAFE_CONDUCTOR_IP=ip
export WCAFE_CONDUCTOR_PORT=port
source ~/.bashrc
```


ユーザーやDB名、キューのURLは以下でも編集可能

```
vi config/config.toml
```
**file_pathは末尾の`/`まで記載すること**  

### DynamoDBの準備
DynamoDBにて、以下のテーブルを用意する

```
テーブル名：clerks_name
プライマリキー：name_id
キー：name
```
```
テーブル名：supplies
プライマリキー：id (string)
ソートキー：name (string)
GSI-プライマリーキー：type (string)
GSI-プライマリーキー：price (int)
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
docker-compose build
```
コンテナの生成と実行
```
docker-compose up -d 
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
mockgen -source service/clerks_service.go -destination mocks/clerks_service.go -package mocks
mockgen -source service/supplies_service.go -destination mocks/supplies_service.go -package mocks
```
