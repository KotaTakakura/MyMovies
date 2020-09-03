# MyPIPE
## 設計方針(ディレクトリ構造)
goのソースコードは、 /go/src/MyPIPEに作成していく
main.goにルーティングを作成していく。
### /go/src/MyPIPE/Controllers
リクエストデータを最初に受け取る部分

### /go/src/MyPIPE/Services
リクエストデータを使い、具体的な処理を行う部分

### /go/src/MyPIPE/Repository
データベースからデータを取得・追加・変更・削除を行なう（ユーザーを追加・検索など）

### /go/src/MyPIPE/Entity
データベースのデータと一対一になる構造体。例えば、名前・年齢を持つユーザーならば、
type User struct {  
  Name string  
  Age int  
}  
