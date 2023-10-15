package logic

import (
	"testing"
)

func TestGetTotalFile(t *testing.T) {
	fl, _ := GetTotalFile("../lexcion")
	t.Log(fl)
}
