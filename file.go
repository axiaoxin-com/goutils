// 文件操作相关通用方法

package goutils

import (
	"io/ioutil"
)

// CopyFile 复制文件
func CopyFile(sourceFile, destinationFile string) (err error) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	return
}
