package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var filePath = "./output.csv"
var bucket = "code_scan"
var key = "code/output.csv"
var awsRegion = "ap-northeast-1"

func main() {
	file, err := os.OpenFile("output.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var stdin string
	fmt.Scan(&stdin)
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{stdin})

	// sessionを作成します
	newSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// S3クライアントを作成します
	svc := s3.New(newSession, &aws.Config{
		Region: aws.String(awsRegion),
	})

	// S3にアップロードする内容をparamsに入れます
	params := &s3.PutObjectInput{
		// Bucket アップロード先のS3のバケット
		Bucket: aws.String(bucket),
		// Key アップロードする際のオブジェクト名
		Key: aws.String(key),
		// Body アップロードする画像ファイル
		Body: "output.csv",

	// S3にアップロードします
	_, err == svc.PutObject(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("uploaded to S3")
	}
}
