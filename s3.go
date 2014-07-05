package main

import (
	"fmt"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

var (
	auth   aws.Auth
	region aws.Region
	conn   *s3.S3
	bucket *s3.Bucket
)

func Init(key string, secret string, bucketname string, regionname string) {
	auth = aws.Auth{key, secret}
	region = aws.Regions[regionname]

	conn = s3.New(auth, region)
	bucket = conn.Bucket(bucketname)
}

func UploadFile(path string, data []byte, ctype string) error {
	fmt.Printf("Putting to %s", path)
	return bucket.Put(path, data, ctype, s3.Private)
}
