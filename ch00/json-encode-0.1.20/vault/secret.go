package vault

import (
	"bytes"
	"io"
	"time"
)

// 構造体の型をエクスポートせずに、フィールドをエクスポートする場合の例
// 型自体はパケージ外で再利用を制限したいが、JSONへの変換やログへの出力を行いたい
// フィールドごとに公開するかを制御する
type secret struct {
	ID         string
	CreateTime time.Time

	token string
}

func (s *secret) Read(p []byte) (int, error) {
	return bytes.NewBuffer(p).WriteString(s.token)
}

func NewSecret() io.Reader {
	return &secret{
		ID:         "dummy_id",
		CreateTime: time.Now(),
		token:      "dummy_token",
	}
}
