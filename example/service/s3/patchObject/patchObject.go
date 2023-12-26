//go:build example
// +build example

package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type CustomReader struct {
	fp      *os.File
	size    int64
	read    int64
	signMap map[int64]struct{}
	mux     sync.Mutex
}

func (r *CustomReader) Read(p []byte) (int, error) {
	return r.fp.Read(p)
}

func (r *CustomReader) ReadAt(p []byte, off int64) (int, error) {
	n, err := r.fp.ReadAt(p, off)
	if err != nil {
		return n, err
	}

	r.mux.Lock()
	// Ignore the first signature call
	if _, ok := r.signMap[off]; ok {
		// Got the length have read( or means has uploaded), and you can construct your message
		r.read += int64(n)
		fmt.Printf("\rtotal read:%d    progress:%d%%", r.read, int(float32(r.read*100)/float32(r.size)))
	} else {
		r.signMap[off] = struct{}{}
	}
	r.mux.Unlock()
	return n, err
}

func (r *CustomReader) Seek(offset int64, whence int) (int64, error) {
	return r.fp.Seek(offset, whence)
}

func main() {
	if len(os.Args) < 4 {
		log.Println("USAGE ERROR: AWS_REGION=us-west-2 go run -tags example putObjWithProcess.go <bucket> <key for object> <local file name>")
		return
	}

	bucket := os.Args[1]
	key := os.Args[2]
	filename := os.Args[3]

	myCustomResolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		if service == endpoints.S3ServiceID {
			return endpoints.ResolvedEndpoint{
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		}

		return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("ru-central1"),
		EndpointResolver: endpoints.ResolverFunc(myCustomResolver),
	}))

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file %v, %v", filename, err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("failed to stat file %v, %v", filename, err)
		return
	}

	reader := &CustomReader{
		fp:      file,
		size:    fileInfo.Size(),
		signMap: map[int64]struct{}{},
	}

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024
		u.LeavePartsOnError = true
	})

	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   reader,
	})
	if err != nil {
		log.Fatalf("failed to put file %v, %v", filename, err)
		return
	}

	fmt.Println()
	log.Println(output.Location)
}
