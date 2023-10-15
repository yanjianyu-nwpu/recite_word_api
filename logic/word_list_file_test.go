package logic

import "testing"

func TestReadText(t *testing.T) {

	b := ReadTextFile(1)
	t.Log(b)

}
