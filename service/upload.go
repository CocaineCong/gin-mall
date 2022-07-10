package service

import (
	"io/ioutil"
	"log"
	"mall/conf"
	"mime/multipart"
	"os"
	"strconv"
)

// UploadProductToLocalStatic 上传到本地文件中
func UploadProductToLocalStatic(file multipart.File, bossId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(bossId))
	basePath := conf.ProductPhotoPath + "boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		return "", err
	}
	return productPath, err
}

// UploadAvatarToLocalStatic 上传头像
func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + userName + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		return "", err
	}
	return productPath, err
}

// DirExistOrNot 判断文件是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	err := os.Mkdir(dirName, 755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
