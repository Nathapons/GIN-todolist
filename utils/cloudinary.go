package utils

import "github.com/cloudinary/cloudinary-go"

func SetUpClodinary() *cloudinary.Cloudinary{
	cld, err := cloudinary.NewFromParams("<your-cloud-name>", "<your-api-key>", "<your-api-secret>")
	if err != nil {
		panic(err)
	}
	return cld
}

func UploadCloudinary() {

}
