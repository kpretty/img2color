package img

import (
	"fmt"
	"testing"
)

func TestImg2Color(t *testing.T) {
	webpUrl := "https://npm.elemecdn.com/anzhiyu-blog@1.1.6/img/post/banner/%E7%A5%9E%E9%87%8C.webp"
	pngUrl := "https://uhope.fun/upload/logo.png"
	webpColor, err := Img2color(webpUrl)
	if err != nil {
		t.Error("webpUrl error", err)
	}
	fmt.Println("webpColor: ", webpColor)
	pngColor, err := Img2color(pngUrl)
	if err != nil {
		t.Error("pngUrl error", err)
	}
	fmt.Println("pngColor: ", pngColor)
}
