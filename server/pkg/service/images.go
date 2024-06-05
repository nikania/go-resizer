package service

import "os"

type ImagesService struct{}

func NewImagesService() *ImagesService {
	return &ImagesService{}
}

func (i *ImagesService) Resize(file *os.File, width, height int, saveRatio bool) (*os.File, error) {
	return nil, nil
}
func (i *ImagesService) Crop(file *os.File, x, y, width, height int) (*os.File, error) {
	return nil, nil
}
func (i *ImagesService) Convert(file *os.File, format string) (*os.File, error) {
	return nil, nil
}
func (i *ImagesService) Compress(file *os.File, quality int) (*os.File, error) {
	return nil, nil
}
