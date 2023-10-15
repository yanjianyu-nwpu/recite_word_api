package logic

import "os"

// GetTotalFile 获取filename
func GetTotalFile(path string) ([]string, error) {

	filepathName, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	res := make([]string, 0)
	for _, f := range filepathName {
		res = append(res, f.Name())
	}

	return res, nil
}
