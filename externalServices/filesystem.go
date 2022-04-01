package externalServices

import (
	"app/config"
	"app/systemService"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type ImageUpload struct {
	Allowed  string
	FileType string
	Path     string
}

func (img *ImageUpload) Upload(r *http.Request) (string, error) {

	filesystem := config.G_STORAGE.Filesystem

	fmt.Println("Filesystem is ", filesystem)

	switch filesystem {
	case "local":
		return localUpload(r, img)
	case "s3":
		return s3Upload(r, img)
	default:
		return localUpload(r, img)
	}
}

func s3Upload(r *http.Request, img *ImageUpload) (string, error) {

	s3Config := &aws.Config{
		Region:      aws.String(config.G_STORAGE.S3Region),
		Credentials: credentials.NewStaticCredentials(config.G_STORAGE.S3KeyId, config.G_STORAGE.S3SecretKey, ""),
	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(s3Config))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	rFile, _, err := r.FormFile(img.Allowed)
	if err != nil {
		fmt.Println("Error retriving file")
		return "", err
	}
	defer rFile.Close()

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.G_STORAGE.S3BucketName),
		Key:    aws.String("images/" + img.Path + systemService.GetRandomString(10) + ".jpg"),
		Body:   rFile,
	})
	if err != nil {
		fmt.Println("failed to upload file", err)
		return "", err
	}

	// fmt.Printf("file uploaded to, %s\n", aws.StringValue(result.Location))
	fmt.Println("file uploaded to bucket. Location = ", result.Location)
	return result.Location, nil
}

func localUpload(r *http.Request, image *ImageUpload) (string, error) {

	file, handler, err := r.FormFile("img")
	if err != nil {
		fmt.Println("Error retriving file")
		return "", err
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header.Get("Content-Type"))

	ftype := strings.Split(handler.Filename, ".")

	fmt.Printf("file type: %+v", ftype[len(ftype)-1])

	os.MkdirAll("storage/images", os.ModePerm)

	tempFile, err := ioutil.TempFile("storage/"+image.Path, "*.jpg")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Sprintln("Successfully Uploaded File")

	return "", nil
}
