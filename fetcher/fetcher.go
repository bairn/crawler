package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
		       fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)

	return all, err
}


func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes , err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}