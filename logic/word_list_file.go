package logic

import (
	"fmt"
	"io/ioutil"
	"os"
)

// GetFileNameByID 返回
func GetFileNameBy(id int) string {
	return fmt.Sprintf("word_list_%d.json", id)
}

// GetFilePath 返回文件位置
func GetFilePath(id int) string {
	return "../lexcion/" + GetFileNameBy(id)
}

/*
	存储样式txt 文件为 英文 abaondon 放弃;丢弃这样的样式
*/
// ReadTextFile 读取文件
func ReadTextFile(id int) (w *WordList) {
	fileP := GetFilePath(id)

	f, err := os.OpenFile(fileP, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return nil
	}

	objByte, err := ioutil.ReadAll(f)

	if err != nil {
		return nil
	}
	w = &WordList{}
	err = w.UnmarshalStr(objByte)
	if err != nil {
		return nil
	}
	return
}

func WriteTextFile(w *WordList) {

}

//func HasFile(id int) bool {

//}
