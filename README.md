# Golang Rest Sample

## これなに

ginのRestサーバーお試し版

## 使い方

db/db.goにAWSのDB情報書く

```
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql
```

```
go run main.go
```

## 動作確認

```
curl localhost:8080/pets
```

## 参考

https://qiita.com/Asuforce/items/0bde8cabb30ac094fcb4