package logic

import "encoding/json"

// Word 单词内容
type Word struct {
	English    string   `json:"english"`
	Chinese    []string `json:"chinese"`
	WrongTimes int      `json:"wrong_times"`
}

// WordList 每个单元的词
type WordList struct {
	ID    int     `json:"id"`
	Words []*Word `json:"words"`
}

// MarshalStr 序列化
func (w *WordList) MarshalStr() ([]byte, error) {
	b, err := json.Marshal(w)
	return b, err
}

// U
func (w *WordList) UnmarshalStr(rb []byte) error {
	return json.Unmarshal(rb, w)
}
