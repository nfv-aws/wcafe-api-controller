# wcafe-api-controller
## 動かし方
1. `go`の作業ディレクトリ `/home/ec2-user/environment/go/src/github.com/hogehoge/`配下で`git clone`する。
2. `hogehoge/wcafe-api-controller/app/server/`で`go run main.go`を叩く。
3. `localhost:8080/stores`にアクセスするとレスポンスが返ってくる。（別Terminalで`curl localhost:8080/stores`を叩く）
備考. `wcafe-api-controller/app/server/go/api_stores_service.go`の`GET`と`POST`部分を編集するとレスポンス（value）を変えられる。
