package file

import (
	"os"
	"path/filepath"
)

func IsExist(path string) (bool, os.FileInfo) {
	info, err := os.Stat(path)
	return err == nil || os.IsExist(err), info
}

func IsDir(path string) (bool, os.FileInfo) {
	flag, info := IsExist(path)
	return flag && info.IsDir(), info
}

func IsFile(path string) (bool, os.FileInfo) {
	flag, info := IsExist(path)
	return flag && !info.IsDir(), info
}

func Rename(from, to string) error {
	return os.Rename(from, to)
}

func Remove(path string) error {
	return os.Remove(path)
}

func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

func EnsureDir(path string) (bool, error) {
	if isExist, _ := IsExist(path); isExist {
		return true, nil
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return false, err
	}
	return true, nil
}

func EnsureFile(filePath string) (bool, error) {
	if isExist, _ := IsExist(filePath); isExist {
		return true, nil
	}

	dir, filename := filepath.Split(filePath)
	if _, err := EnsureDir(dir); err != nil {
		return false, err
	}

	file, err := os.Create(filename)
	if err != nil {
		return false, err
	}
	file.Close()

	return true, nil
}
