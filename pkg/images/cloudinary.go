package images

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type Cloudinary struct {
	// client cloudinary
	Cloud   *cloudinary.Cloudinary
	IsError error
}

func NewCloudinary(cloud, apiKey, apiSecret string) Cloudinary {
	c, err := cloudinary.NewFromParams(cloud, apiKey, apiSecret)

	return Cloudinary{
		Cloud:   c,
		IsError: err,
	}
}

func (c Cloudinary) Upload(ctx context.Context, file interface{}, path string) (string, error) {
	filename := uuid.NewString()

	res, err := c.Cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: path + "/" + filename,
		Eager:    "q_auto:low",
	})

	if err != nil {
		return "", err
	}

	if len(res.Eager) > 0 {
		return res.Eager[0].SecureURL, nil
	}

	url := res.SecureURL

	return url, nil
}
