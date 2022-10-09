package utils

import (
	"fmt"
	"os"
)

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateMultiDir 调用os.MkdirAll递归创建文件夹
func CreateMultiDir(filePath string) error {
	_, err := os.Stat(filePath) //os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
	}
	return nil
}
