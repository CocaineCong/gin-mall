package upload

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"strconv"

	"mall/conf"
	util "mall/pkg/utils/log"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// UploadProductToLocalStatic 上传到本地文件中
func UploadProductToLocalStatic(file multipart.File, bossId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(bossId))
	basePath := "." + conf.ProductPhotoPath + "boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := fmt.Sprintf("%s%s.jpg", basePath, productName)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}
	return fmt.Sprintf("boss%s/%s.jpg", bId, productName), err
}

// UploadAvatarToLocalStatic 上传头像
func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := fmt.Sprintf("%s%s.jpg", basePath, userName)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}
	return fmt.Sprintf("user%s/%s.jpg", bId, userName), err
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
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// UploadToQiNiu 封装上传图片到七牛云然后返回状态和图片的url，单张
func UploadToQiNiu(file multipart.File, fileSize int64) (path string, err error) {
	var AccessKey = conf.AccessKey
	var SerectKey = conf.SerectKey
	var Bucket = conf.Bucket
	var ImgUrl = conf.QiniuServer
	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}
	url := ImgUrl + ret.Key
	return url, nil
}
