package img

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/nfnt/resize"
	"golang.org/x/image/webp"
)

/*
传入图片url，返回图片的颜色
1. 下载图片并转换为字节数组
2. 计算图片的颜色
3. 返回图片的颜色
*/

func Img2color(url string) (string, error) {
	img, err := downloadImg(url)
	if err != nil {
		return "", err
	}
	color, err := extractMainColor(img)
	if err != nil {
		return "", err
	}
	return color, nil
}

// 下载图片并转换为字节数组
func downloadImg(url string) (image.Image, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var img image.Image
	contentType := resp.Header.Get("Content-Type")
	if contentType == "image/webp" {
		img, err = webp.Decode(resp.Body)
	} else {
		img, err = imaging.Decode(resp.Body)
	}
	if err != nil {
		return nil, err
	}
	return img, nil
}

// 计算图片的颜色
func extractMainColor(img image.Image) (string, error) {
	img = resize.Resize(50, 0, img, resize.Lanczos3)
	bounds := img.Bounds()
	var r, g, b uint32
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r0, g0, b0, _ := c.RGBA()
			r += r0
			g += g0
			b += b0
		}
	}

	totalPixels := uint32(bounds.Dx() * bounds.Dy())
	averageR := r / totalPixels
	averageG := g / totalPixels
	averageB := b / totalPixels

	mainColor := colorful.Color{R: float64(averageR) / 0xFFFF, G: float64(averageG) / 0xFFFF, B: float64(averageB) / 0xFFFF}

	return mainColor.Hex(), nil
}
