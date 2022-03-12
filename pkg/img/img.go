package img

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

//func GetLogo() (*bytes.Buffer, error) {
func GetLogo() ([]byte, error) {
	//buffer := new(bytes.Buffer)

	//m := image.NewRGBA(image.Rect(16, 16, 32, 32))
	//clr := color.RGBA{}
	//draw.Draw(m, m.Bounds(), &image.Uniform{C: clr}, image.ZP, draw.Src)

	data, err := ioutil.ReadFile("data/logo.png")
	if err != nil {
		logrus.Println(err)
	}

	//var img image.Image = m
	//var img image.Image = data
	//if err := jpeg.Encode(buffer, img, nil); err != nil {
	//if err := jpeg.Encode(buffer, data, nil); err != nil {
	//	return nil, err
	//}

	return data, nil
	//return data, nil
}
