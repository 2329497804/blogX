package qiniu_service

import (
	"context"
	"fmt"
	"io"
	"uuid"

	"github.com/2329497804/blogX-server/global"
	"github.com/2329497804/blogX-server/utils/files"
	"github.com/2329497804/blogX-server/utils/hashs"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
)

func SendFile(filePath string, ctx context.Context) (url string, err error) {
	q := global.Config.QiNiu

	mac := credentials.NewCredentials(q.AccessKey, q.SecretKey)

	hashString, err := hashs.FileMD5(filePath)
	if err != nil {
		return
	}

	suffix, err := files.ImageSuffixJudge(filePath)
	if err != nil {
		return
	}
	fileName := fmt.Sprintf("%s.%s", hashString, suffix)
	key := fmt.Sprintf("%s/%s", q.Prefix, fileName)

	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{Credentials: mac})
	err = uploadManager.UploadFile(ctx, filePath,
		&uploader.ObjectOptions{
			BucketName: q.Bucket,
			ObjectName: &key,
			FileName:   fileName,
		}, nil)
	if err != nil {
		return
	}
	url = fmt.Sprintf("%s/%s", q.URI, key)
	return
}

func SendReader(reader io.Reader, ctx context.Context) (url string, err error) {
	q := global.Config.QiNiu

	mac := credentials.NewCredentials(q.AccessKey, q.SecretKey)

	uid := uuid.New().String()
	fileName := fmt.Sprintf("%s.png", uid)
	key := fmt.Sprintf("%s/%s", q.Prefix, fileName)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{Credentials: mac})
	err = uploadManager.UploadReader(ctx, reader,
		&uploader.ObjectOptions{
			BucketName: q.Bucket,
			ObjectName: &key,
			FileName:   fileName,
		}, nil)
	if err != nil {
		return
	}
	url = fmt.Sprintf("%s/%s", q.URI, key)
	return
}