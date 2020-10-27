# MyMovies
※このReadmeは作成途中です。
## 使用言語
Go言語  
Nuxt(https://github.com/takatsu111/MyPIPE-front)

## 使用技術(実装予定も含む)
VPC  
ECS(Fargate)  
ALB  
Lambda  
MediaConvert  
SNS  
AuroraDB

## 設計方針(ディレクトリ構造)
goのソースコードは、 /go/src/MyPIPEに作成していく  
main.goにルーティングを作成していく。
### /go/src/MyPIPE/domain
ビジネスルール・仕様の関心事を取り扱う。  
User structやComment structをドメインオブジェクトとして実装。

### /go/src/MyPIPE/handler
リクエストデータを受け取り、加工したものをusecase層にデータを渡す。

### /go/src/MyPIPE/infra
永続化のための具体的な技術を実装する。

### /go/src/MyPIPE/usecase
ソフトウェアの行なう仕事の流れを表現する。
