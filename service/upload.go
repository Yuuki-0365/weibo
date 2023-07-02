package service

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"log"
	"mime/multipart"
	"os"
	"weibo/conf"
)

type UploadService struct {
}

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
		Zone:          &storage.ZoneHuadongZheJiang2,
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

func UploadNoteFileToLocalStatic(file multipart.File, userId string, noteId string, count string) (filePath string, err error) {
	// ./static/imgs/note/user123/note123/
	basePath := "." + conf.NotePath + "user" + userId + "/" + "note" + noteId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	NoteFilePath := basePath + count + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(NoteFilePath, content, 0777)
	if err != nil {
		return "", err
	}
	return basePath, nil
}

func DirExistOrNot(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
