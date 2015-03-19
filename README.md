S3Putter
========

Simple Golang app to read from the STDIN and turn that into a file inside an S3 bucket.

Naming scheme:

`YYYY-MM-DD--HH-MM-SS-NNN` (where NNN is nanosecond to three places)

Filename will be a formatted datetime, down to the nanosecond to avoid naming collision.

Currently hardcoded to use `aws.APSoutheast2` as the bucket location, but shouldn't be hard to change.

Require environment variables:

```
AWS_ACCESS_KEY_ID=xxx
AWS_SECRET_ACCESS_KEY=xxx
```

either exported or injected somehow, or on the command line.

Example usage:

`cat samples/1161 | AWS_ACCESS_KEY_ID=<your access key id> AWS_SECRET_ACCESS_KEY=<your secret key> go run main.go -bucket=<your bucket name>`


This is designed to work with an Alaveteli installation to separate mail parsing out from the main Rails app installation.
