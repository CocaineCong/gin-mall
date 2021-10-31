package service

import (
	"FanOneMall/pkg/e"
	"FanOneMall/pkg/logging"
	"FanOneMall/serializer"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"mime"
	"os"
	"path/filepath"
)

// UploadAvatarService 获得上传oss token的服务
type UploadAvatarService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 创建token
func (service *UploadAvatarService) Post() serializer.Response {
	code := e.SUCCESS
	client, err := oss.New("oss-cn-beijing.aliyuncs.com" , "LTAI4G9m5vcxdmrjLDuG5Uuf", "DNmjvI8n5rHFwhtIyPNtgsYBLyLaVa")
	if err != nil {
		logging.Info(err)
		code = e.ErrorOss
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 获取存储空间。
	//bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	bucket, err := client.Bucket("gomall1" )
	if err != nil {
		logging.Info(err)
		code = e.ErrorOss
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 获取扩展名
	ext := filepath.Ext(service.Filename)
	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
	}
	//key := "upload/avatar/" + ext
	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext
	// 签名直传。
	fmt.Println("Filename")
	fmt.Println(service.Filename)
	err = bucket.PutObjectFromFile(service.Filename, "D:\\CodeProjects\\GoLandProjects\\GoSuperMark\\gin-mail\\static\\img\\avatar\\1.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		logging.Info(err)
		code = e.ErrorOss
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 查看图片
	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		logging.Info(err)
		code = e.ErrorOss
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}

//
//func (service *UploadAvatarService) Post(c *gin.Context) serializer.Response {
//	code :=e.SUCCESS
//	_ , _, err := c.Request.FormFile("avator")
//	fileFullPath := "./uploadfile/" + service.Filename
//	_ , err = os.Create(fileFullPath)
//	if err != nil {
//		logging.Error(err.Error())
//		return serializer.Response{
//			Status: 500,
//			Msg:    e.GetMsg(code),
//			Error:  err.Error(),
//		}
//	}
//	return serializer.Response{
//		Status: code,
//		Msg:    e.GetMsg(code),
//		Data: map[string]string{
//			"put": fileFullPath,
//			"get": fileFullPath,
//		},
//	}
//}
