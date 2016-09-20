package qiniu

import (
	"log"
	"path"
	"testing"

	blobconf "github.com/DaoCloud/mirror-blob-server/blob_server/conf"
)

func TestGetUpLoad(t *testing.T) {

	blobconf.QiniuAccessKey = ""
	blobconf.QiniuSecretKey = ""
	blobconf.QiniuBucketName = "registrytest"
	//blobconf.QiniuDomain = "o9kxond0d.bkt.clouddn.com"
	//blobconf.QiniuDomain = "o9kxond0d.upload.qiniu.com"

	qiniuManager, err := NewQiniuClient()
	if err != nil {
		t.Error(err)
	}
	dirPath := "."
	fileName := "testupload"
	filePath := path.Join(dirPath, fileName)
	err = qiniuManager.UpLoad(filePath, fileName)
	if err != nil {
		t.Error(err)

	}
	testKey := "testblob.txt"
	privateURL := qiniuManager.PrivateURL(testKey)
	log.Println("the private url: ", privateURL)

	//test existance
	isexistance, fsize := qiniuManager.CheckExistance("testkey")
	log.Println("the size: ", fsize)
	//expect true
	t.Log(isexistance)

	isexistance, fsize = qiniuManager.CheckExistance("abc")
	//expect false
	log.Println("the size: ", fsize)
	t.Log(isexistance)
}
