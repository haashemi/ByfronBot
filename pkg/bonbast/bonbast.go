package bonbast

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Client struct {
	sync.Mutex

	client *http.Client

	cachedAt   time.Time
	cachedData *Response

	token          string
	tokenCreatedAt time.Time
}

type Response struct {
	Aed1         string `json:"aed1"`
	Aed2         string `json:"aed2"`
	Afn1         string `json:"afn1"`
	Afn2         string `json:"afn2"`
	Amd1         string `json:"amd1"`
	Amd2         string `json:"amd2"`
	Aud1         string `json:"aud1"`
	Aud2         string `json:"aud2"`
	Azadi1       string `json:"azadi1"`
	Azadi12      string `json:"azadi12"`
	Azadi120     string `json:"azadi1_2"`
	Azadi122     string `json:"azadi1_22"`
	Azadi14      string `json:"azadi1_4"`
	Azadi142     string `json:"azadi1_42"`
	Azadi1G      string `json:"azadi1g"`
	Azadi1G2     string `json:"azadi1g2"`
	Azn1         string `json:"azn1"`
	Azn2         string `json:"azn2"`
	Bhd1         string `json:"bhd1"`
	Bhd2         string `json:"bhd2"`
	Bitcoin      string `json:"bitcoin"`
	Bourse       string `json:"bourse"`
	Cad1         string `json:"cad1"`
	Cad2         string `json:"cad2"`
	Chf1         string `json:"chf1"`
	Chf2         string `json:"chf2"`
	Cny1         string `json:"cny1"`
	Cny2         string `json:"cny2"`
	Created      string `json:"created"`
	Day          int    `json:"day"`
	Dkk1         string `json:"dkk1"`
	Dkk2         string `json:"dkk2"`
	Emami1       string `json:"emami1"`
	Emami12      string `json:"emami12"`
	Eur1         string `json:"eur1"`
	Eur2         string `json:"eur2"`
	Gbp1         string `json:"gbp1"`
	Gbp2         string `json:"gbp2"`
	Gol18        string `json:"gol18"`
	Hkd1         string `json:"hkd1"`
	Hkd2         string `json:"hkd2"`
	Hour         string `json:"hour"`
	Inr1         string `json:"inr1"`
	Inr2         string `json:"inr2"`
	Iqd1         string `json:"iqd1"`
	Iqd2         string `json:"iqd2"`
	Jpy1         string `json:"jpy1"`
	Jpy2         string `json:"jpy2"`
	Kwd1         string `json:"kwd1"`
	Kwd2         string `json:"kwd2"`
	LastModified string `json:"last_modified"`
	Minute       string `json:"minute"`
	Mithqal      string `json:"mithqal"`
	Month        int    `json:"month"`
	Myr1         string `json:"myr1"`
	Myr2         string `json:"myr2"`
	Nok1         string `json:"nok1"`
	Nok2         string `json:"nok2"`
	Omr1         string `json:"omr1"`
	Omr2         string `json:"omr2"`
	Ounce        string `json:"ounce"`
	Qar1         string `json:"qar1"`
	Qar2         string `json:"qar2"`
	Rub1         string `json:"rub1"`
	Rub2         string `json:"rub2"`
	Sar1         string `json:"sar1"`
	Sar2         string `json:"sar2"`
	Second       string `json:"second"`
	Sek1         string `json:"sek1"`
	Sek2         string `json:"sek2"`
	Sgd1         string `json:"sgd1"`
	Sgd2         string `json:"sgd2"`
	Thb1         string `json:"thb1"`
	Thb2         string `json:"thb2"`
	Try1         string `json:"try1"`
	Try2         string `json:"try2"`
	Usd1         string `json:"usd1"`
	Usd2         string `json:"usd2"`
	Weekday      string `json:"weekday"`
	Year         int    `json:"year"`
}

func NewClient(proxyUrl string) (*Client, error) {
	client := &Client{
		client: &http.Client{},
	}

	if proxyUrl != "" {
		proxy, err := url.Parse(proxyUrl)
		if err != nil {
			return nil, err
		}

		client.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	}

	return client, nil
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("accept-language", "en-US,en-GB;q=0.9,en;q=0.8,fa;q=0.7")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"110\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("cookie", "st_bb=0")
	req.Header.Set("Referer", "https://bonbast.com/")
	req.Header.Set("Referrer-Policy", "strict-origin-when-cross-origin")
}

func (c *Client) getToken() (string, error) {
	if c.token != "" && time.Since(c.tokenCreatedAt).Hours() < 48 {
		return c.token, nil
	}

	req, err := http.NewRequest("GET", "https://bonbast.com/", nil)
	if err != nil {
		return "", err
	}
	c.setHeaders(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	c.token = strings.SplitN(strings.SplitN(string(bb), "$.post('/json', {param: \"", 2)[1], "\"},", 2)[0]
	c.tokenCreatedAt = time.Now()

	return c.token, nil
}

func (c *Client) GetData() (*Response, error) {
	c.Lock()
	defer c.Unlock()

	if c.cachedData != nil && time.Since(c.cachedAt).Seconds() < 30 {
		return c.cachedData, nil
	}

	token, err := c.getToken()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://bonbast.com/json", strings.NewReader((url.Values{"param": {token}}).Encode()))
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data Response
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	c.cachedData = &data
	c.cachedAt = time.Now()

	return c.cachedData, nil
}
