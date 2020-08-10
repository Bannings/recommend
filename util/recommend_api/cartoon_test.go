package recommend_api

import (
	"fmt"
	"github.com/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
	"recommend/golbal"
	"strings"
	"testing"
)

func TestFetchComicInfos(t *testing.T) {
	_, err := golbal.LoadConfig("../dev.json")
	if err != nil {
		t.Fatal(err)
	}
	e := []string{"89567548", "male"}
	a := strings.Join(e, "_")
	t.Log(a)
	tmpComicInfo, err := QueryComicInfo("77")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tmpComicInfo)
}

func DecodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}

func TestRead(t *testing.T) {

	url := "http://m.37zw.net/7/7579/5329349.html"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Fatal(err)

	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	title := doc.Find("title").Text()
	title, _ = DecodeToGBK(title)
	//next := doc.Find(".bottem2 a").Eq(3)
	next := doc.Find("#pb_next")
	content := doc.Find("#nr1")
	con := content.Text()
	con, _ = DecodeToGBK(con)
	herf, _ := next.Attr("href")
	fmt.Println(title)
	fmt.Println(con)
	fmt.Println(herf)
}
