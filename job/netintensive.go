package job

import (
	"log"

	"github.com/resourcebench/util/storage/qiniu"
)

func NetBenchmark(interval int, QiniuAccessKey string, QiniuSecretKey string) {

	//init qiniu client
	QiniuBucketName := "registrytest"
	QiniuDomain := "o9kxond0d.bkt.clouddn.com"

	qiniuManager, err := qiniu.NewQiniuClient(QiniuAccessKey, QiniuSecretKey, QiniuBucketName, QiniuDomain)
	if err != nil {
		log.Println(err)
	}

	log.Println("start io intensive testing")

	//ticker := time.NewTicker(time.Millisecond * time.Duration(interval))
	for {
		UploadLoad(qiniuManager)
		go DownloadLoad(qiniuManager)
	}

	//for t := range ticker.C {
	//log.Println("Tick at", t)
	//create the basic load
	//after := time.After(time.Millisecond * 1000 * 5)
	//go UploadLoad(qiniuManager)
	//go DownloadLoad(qiniuManager)
	//}

}

func UploadLoad(manager *qiniu.QiniuManager) {

	err := manager.UpLoad("/opt/testupload.tar.gz", "test_upload")
	if err != nil {
		log.Println(err)
	}

}

func DownloadLoad(manager *qiniu.QiniuManager) {
	DownloadUrl := "http://o9kxond0d.bkt.clouddn.com/test_upload"
	err := manager.DownLoadPublic(DownloadUrl)
	if err != nil {
		log.Println(err)
	}
}
