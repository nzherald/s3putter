package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

var bucketName string

func init() {
	flag.StringVar(&bucketName, "bucket", "", "Bucket Name")
}

func main() {

	flag.Parse()

	stats, err := os.Stdin.Stat()
	checkErr(err)

	var buf []byte
	buf, err = ioutil.ReadAll(os.Stdin)
	checkErr(err)

	auth, err := aws.EnvAuth()
	checkErr(err)

	if stats.Size() > 0 {
		// Open Bucket
		s := s3.New(auth, aws.APSoutheast2)
		bucket := s.Bucket(bucketName)

		t := time.Now()
		fileName := formatTime(t)

		err = bucket.Put(fileName, buf, "text/plain", s3.BucketOwnerFull)
		checkErr(err)
		fmt.Printf("Created %s", fileName)

	} else {
		fmt.Println("Nothing on STDIN")
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func formatTime(t time.Time) string {
	return fmt.Sprintf("%d%02d%02d-%02d-%02d-%02d-%03d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/100000)
}
