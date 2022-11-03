package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	path := "./Pic Download/"
	pathExit := IsExist(path)
	if !pathExit {
		os.Mkdir("Pic Download", os.ModePerm)
	}

	//開始爬蟲
	fmt.Println("start main")
	c := colly.NewCollector()
	count := 0

	// 要爬蟲的相關 KEY 值 //此為 LY平台登入頁面 的title
	c.OnHTML(".mb-box li img[src]", func(e *colly.HTMLElement) {
		var picurl string
		fmt.Println("目標值:", e.Text)
		fmt.Println("Pic imf:", e.Attr("src"))
		picurl = e.Attr("src")
		getImg(picurl)
		count++
	})

	//div(標頭類型)直接輸入即可
	//class: .form-control, 後面直接加div -->抓到整個範圍的div
	//id: #logintitle

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	// 此為 爬蟲目標網站 網站
	c.Visit("https://zh.pngtree.com/free-vectors?source_id=310&chnl=ggas&srid=764883278&gpid=57293140505&asid=423823300736&ntwk=g&tgkw=kwd-286971347&mchk=%E5%9C%96%E5%BA%AB&mcht=b&pylc=9040379&dvic=c&gclid=Cj0KCQjwkt6aBhDKARIsAAyeLJ1bU9NabQ7KjghGeeBGvWAZDHUiFyuIp6Jigh174AiOGl9AiEgUeHoaAraCEALw_wcB")

	fmt.Println(count)
}

func getImg(url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = "Pic Download/" + path[len(path)-1]
	}
	//"../output" 在go檔案的上一頁
	fmt.Println(name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
	// 或者
	//return err == nil || !os.IsNotExist(err)
	// 或者
	//return !os.IsNotExist(err)
}

// 爬蟲相關 套件 : go get -u github.com/gocolly/colly/v2
