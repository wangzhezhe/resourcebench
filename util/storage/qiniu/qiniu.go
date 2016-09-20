package qiniu

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

type QiniuManager struct {
	ACCESS_KEY string
	SECRET_KEY string
	BUCKETNAME string
	DOMAIN     string
	Client     *kodo.Client
}

type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func NewQiniuClient(accesskey string, secretkey string, bucketname string, domain string) (*QiniuManager, error) {

	if accesskey == "" || secretkey == "" || bucketname == "" || domain == "" {
		return nil, errors.New("failed to create the qiniuclient, ACCESS_KEY SECRET_KEY BUCKETNAME DOMAIN should be included in env")
	}
	//assign the key before create the client

	conf.ACCESS_KEY = accesskey
	conf.SECRET_KEY = secretkey
	kodoClient := kodo.New(0, nil)
	qiniuClient := &QiniuManager{
		ACCESS_KEY: accesskey,
		SECRET_KEY: secretkey,
		BUCKETNAME: bucketname,
		DOMAIN:     domain,
		Client:     kodoClient,
	}

	return qiniuClient, nil

}

//the file name should be like tempath/name/blobsum
//refer to http://developer.qiniu.com/code/v7/sdk/go.html
func (q *QiniuManager) UpLoad(filePath string, uploadKey string) error {
	fmt.Printf("the manager %+v\n", q)

	policy := &kodo.PutPolicy{
		Scope:   q.BUCKETNAME,
		Expires: 3600,
		//the key saved in qiniu (could be sepreated by /)
		SaveKey: uploadKey,
		//refer to : http://www.iana.org/assignments/media-types/application/octet-stream
	}

	token := q.Client.MakeUptoken(policy)

	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet

	res := uploader.PutFile(context.Background(), &ret, token, uploadKey, filePath, nil)

	if res != nil {
		fmt.Println("io.Put failed:", res)
		return errors.New("failed for uploading: " + res.Error())
	}
	fmt.Printf("upload %s successfully\n", uploadKey)
	return nil
}

func (q *QiniuManager) DownLoadPublic(publicUrl string) error {
	reqest, err := http.NewRequest("GET", publicUrl, nil)
	if err != nil {
		return err
	}

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	_, err = client.Do(reqest)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//check the existance of the key
//the key is manifestid directly
//the unit of fsize is byte
func (q *QiniuManager) CheckExistance(key string) (bool, int) {
	p := q.Client.Bucket(q.BUCKETNAME)
	entry, err := p.Stat(nil, key)
	if err != nil {
		return false, -1
	}
	//debug
	fmt.Println("get entry info in qiniu: ", entry)
	return true, int(entry.Fsize)

}

//down load the file with key into the file path on local machine
func (q *QiniuManager) PrivateURL(key string) string {
	domain := os.Getenv("DOMAIN")
	baseUrl := kodo.MakeBaseUrl(domain, key)
	policy := kodo.GetPolicy{}

	privateUrl := q.Client.MakePrivateUrl(baseUrl, &policy)

	return privateUrl
}
