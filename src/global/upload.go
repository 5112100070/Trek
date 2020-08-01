package global

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func UploadCompanyImage(companyImage multipart.File, filename string) (string, error) {
	var result string

	var imageDir string
	if GetEnv() == "production" {
		imageDir = "/var/www/trek/img/partner"
	} else {
		imageDir = "files/WEB-INF/attr/img/partner"
	}

	filename = strings.ToLower(filename)
	filename = regexResImageNameChar.ReplaceAllString(filename, "-")
	filename = strings.Replace(filename, " ", "-", -1)

	// file name for response
	result = fmt.Sprintf("%s.jpg", filename)

	// image saved on this path
	filename = fmt.Sprintf("%s/%s.jpg", imageDir, filename)

	errCreateFile := createFileOnServer(filename, companyImage)
	if errCreateFile != nil {
		Error.Println("func UploadCompanyImage error create new file with name: %v, error: %v", filename, errCreateFile)
		return result, errCreateFile
	}

	return result, nil
}

func UploadProfileImage(profileImage multipart.File, filename string) (string, error) {
	var result string

	var imageDir string
	if GetEnv() == "production" {
		imageDir = "/var/www/trek/img/user"
	} else {
		imageDir = "files/WEB-INF/attr/img/user"
	}

	filename = strings.ToLower(filename)
	filename = regexResImageNameChar.ReplaceAllString(filename, "-")
	filename = strings.Replace(filename, " ", "-", -1)

	// file name for response
	result = fmt.Sprintf("%s.jpg", filename)

	// image saved on this path
	filename = fmt.Sprintf("%s/%s.jpg", imageDir, filename)

	errCreateFile := createFileOnServer(filename, profileImage)
	if errCreateFile != nil {
		Error.Println("func UploadProfileImage error create new file with name: %v, error: %v", filename, errCreateFile)
		return result, errCreateFile
	}

	return result, nil
}

func createFileOnServer(filename string, payload multipart.File) error {
	dst, err := os.Create(filename)
	if err != nil {
		Error.Println("func createFileOnServer error create new file with name: %v, error: %v", err)
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, payload)
	if err != nil {
		Error.Println("func createFileOnServer error write company image to server: ", err)
		return err
	}

	Error.Println("Success write new file")
	return nil
}
