package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

func main() {

	bucketName := os.Getenv("S3_BUCKET")

	var buf []byte
	buf, err := ioutil.ReadAll(os.Stdin)
	checkErr(err)

	auth, err := aws.EnvAuth()
	checkErr(err)

	// Open Bucket
	s := s3.New(auth, aws.APSoutheast2)
	bucket := s.Bucket(bucketName)

	t := time.Now()
	fileName := formatTime(t)

	err = bucket.Put(fileName, buf, "text/plain", s3.BucketOwnerFull)
	checkErr(err)
	fmt.Printf("Successfully received email and saved in S3 as %s\n", fileName)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		// exit with code 75 for Postfix to know to bounce mail
		os.Exit(75)
	}
}

func formatTime(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d--%02d-%02d-%02d-%03d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/100000)
}
