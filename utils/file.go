package utils

import (
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// 创建临时文件
func CreateFile(bt []byte) string {

	//✅:将字节读取到文件

	tmpFile, err := ioutil.TempFile(os.TempDir(), "simple-.*.png")
	log.Print(tmpFile.Name())
	if err != nil {
		log.Print("创建临时文件失败", err)
	}
	// write to file
	_, err = tmpFile.Write(bt)

	if err != nil {
		log.Print("写入文件失败", err)
	}
	defer func() {
		_ = os.Remove(tmpFile.Name())
		_ = tmpFile.Close()
	}()
	file, err := os.Open(tmpFile.Name())
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	// 图片压缩时使用

	sourcebuffer, _ := ioutil.ReadAll(file)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer)

	return sourcestring
}

func photo(tmpFile *os.File) []byte {
	// ✅:图片压缩

	// decode jpeg into image.Image
	var file *os.File
	file, err := os.Open(tmpFile.Name())
	if err != nil {
		log.Print(err)
	}
	defer file.Close()
	var img image.Image
	img, err = jpeg.Decode(file)
	if err != nil {
		//log.Print("图片转换错误:", err.Error())
		file, err = os.Open(tmpFile.Name())
		img, err = jpeg.Decode(file)
	}

	m := resize.Resize(800, 0, img, resize.Lanczos3)
	tmp, err := ioutil.TempFile(os.TempDir(), "transfer-.*.png")
	defer func() {
		_ = os.Remove(tmp.Name())
		_ = tmp.Close()
	}()
	if err != nil {
		log.Print("创建临时文件失败", err)
	}

	// write new image to file
	_ = png.Encode(tmp, m)

	ff, _ := os.Open(tmp.Name())
	defer ff.Close()

	buf, _ := ioutil.ReadAll(ff)
	return buf
}
