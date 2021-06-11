package usecase

import (
	"cloud.google.com/go/storage"
)

//VideoService : Service layer to incorporate all the apis related to Videos
type VideoService struct{}

//UploadToStorage
func (vs *VideoService) UploadVideo(c gin.Context) (statusCode int, success bool, message string) {
	statusCode := http.StatusOK
	success := true
	message := "video uploaded successfully."
	video, err := c.FormFile("video")
	if err == nil {
		wg.Add(1)
		go func(video *multipart.FileHeader) {
			fileExtension := filepath.Ext(video.Filename)
			videoFileName = uuid.New().String() + fileExtension
			file, err := video.Open()
			defer file.Close()
			if err == nil {
				videoBytes, err := ioutil.ReadAll(file)
				if err == nil {
					uploaded := vs.UploadToStorage(videoBytes, videoFileName, fileExtension, "video")
					if !uploaded {
						statusCode = http.StatusInternalServerError
						success = false
						message := "Uploading of video has failed."
					}
				} else {
					statusCode = http.StatusInternalServerError
					success = false
					message := "Uploading of video has failed."
				}
			} else {
				statusCode = http.StatusInternalServerError
				success = false
				message := "Uploading of video has failed."
			}
			wg.Done()
		}(video)
	}
	return statusCode, success, message
}

//UploadToS3 upload byte array to s3
func (vs *VideoService) UploadToStorage(fileBytes []byte, fileName, ext, fileType string) (res bool) {
	func configureStorage(bucketID string) (*storage.BucketHandle, error) {
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			return nil, err
		}
		return client.Bucket(bucketID), nil
	}
}
