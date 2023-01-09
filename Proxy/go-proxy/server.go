package goproxy

import "fmt"

/*
Proxy と Service 共通の API を示す ServiceInterface
Service と Proxy はともに ServiceInterface を実装する

# Client から見ると区別がつかない

そのため Proxy は
- キャッシュ
- 認証
- ログ
- 遅延ロード
など Service の処理を一部肩代わりしたり、処理を追加したりする
*/
type ServerInterface interface {
	HttpRequest(string) string
}

type Server struct {
}

func (s *Server) HttpRequest(message string) string {
	return fmt.Sprintf("Hello! %s", message)
}
