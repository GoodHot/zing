package util

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var fileUtil = &Files{}

// FileUtil 获取工具类
func FileUtil() *Files {
	return fileUtil
}

// Files 文件操作类
type Files struct{}

// ReadFile 读取文件
func (Files) ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

// AbsPath 获取文件绝对路径
func (f Files) AbsPath(path string) string {
	file, _ := os.Getwd()
	if strings.Index(path, "/") != 0 {
		return f.ConvertSeparator(file + "/" + path)
	}
	return f.ConvertSeparator(file + path)
}

// Exist 判断是否存在
func (f Files) Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// ConvertSeparator 系统分隔符转换
func (f Files) ConvertSeparator(path string) string {
	if os.PathSeparator == '\\' {
		return strings.Replace(path, "/", "\\", -1)
	}
	return strings.Replace(path, "\\", "/", -1)
}

// CopyFile 复制文件
func (f Files) CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// FileList 读取文件夹里的所有文件名
func (f Files) FileList(folder string, absPath bool) []string {
	var result []string
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if !file.IsDir() {
			if absPath {
				result = append(result, folder+file.Name())
			} else {
				result = append(result, file.Name())
			}
		}
	}
	return result
}
