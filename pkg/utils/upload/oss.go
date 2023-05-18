package upload

import (
	"context"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"

	conf "github.com/CocaineCong/gin-mall/config"
)

// UploadToQiNiu 封装上传图片到七牛云然后返回状态和图片的url，单张
func UploadToQiNiu(file multipart.File, fileSize int64) (path string, err error) {
	qConfig := conf.Config.Oss
	var AccessKey = qConfig.AccessKeyId
	var SerectKey = qConfig.AccessKeySecret
	var Bucket = qConfig.BucketName
	var ImgUrl = qConfig.QiNiuServer
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
