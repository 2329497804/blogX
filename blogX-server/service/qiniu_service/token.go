package qiniu_service

import (
	"context"
	"time"

	"github.com/2329497804/blogX-server/global"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
)

func GenerateToken(ctx context.Context) (token string, err error) {
	q := global.Config.QiNiu

	mac := credentials.NewCredentials(q.AccessKey, q.SecretKey)
	
	bucket := q.Bucket
	expiry := time.Now().Add(time.Duration(q.Expire) * time.Second)
	putPolicy, err := uptoken.NewPutPolicy(bucket, expiry)
	if err != nil {
		return
	}

	return uptoken.NewSigner(putPolicy, mac).GetUpToken(ctx)
}
