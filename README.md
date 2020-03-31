# wcafe-api-controller
## 動かし方
1. `go`の作業ディレクトリ で`git clone`する。
```
$ cd /home/ec2-user/environment/go/src/github.com/hogehoge/
$ git clone git@hoge
```
2. `main.go`を実行する。
```
$ cd wcafe-api-controller/app/server/
$ go run main.go
```
3. `localhost:8080/stores`にアクセスするとレスポンスが返ってくる。下記コマンドを別ターミナル
```
$ curl localhost:8080/stores
{"id":"IDIDIDID","name":"example","tag":"Tag"}
```
備考. `wcafe-api-controller/app/server/go/api_stores_service.go`の`GET`と`POST`部分を編集するとレスポンス（value）を変えられる。
