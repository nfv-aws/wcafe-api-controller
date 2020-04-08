# Golang Rest Sample

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

db/db.goにAWSのDB情報書く

```
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql
go get github.com/google/uuid
```

```
go run main.go
```

## 動作確認

```
curl localhost:8080/pets
```

## UnitTest

```
cd server
go test

PASS
ok      github.com/nfv-aws/wcafe-api-controller/server  0.206s
```

## 参考

https://qiita.com/Asuforce/items/0bde8cabb30ac094fcb4