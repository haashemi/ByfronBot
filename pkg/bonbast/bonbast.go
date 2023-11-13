package bonbast

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
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
	Aed1         Number `json:"aed1"`
	Aed2         Number `json:"aed2"`
	Afn1         Number `json:"afn1"`
	Afn2         Number `json:"afn2"`
	Amd1         Number `json:"amd1"`
	Amd2         Number `json:"amd2"`
	Aud1         Number `json:"aud1"`
	Aud2         Number `json:"aud2"`
	Azadi1       Number `json:"azadi1"`
	Azadi12      Number `json:"azadi12"`
	Azadi120     Number `json:"azadi1_2"`
	Azadi122     Number `json:"azadi1_22"`
	Azadi14      Number `json:"azadi1_4"`
	Azadi142     Number `json:"azadi1_42"`
	Azadi1G      Number `json:"azadi1g"`
	Azadi1G2     Number `json:"azadi1g2"`
	Azn1         Number `json:"azn1"`
	Azn2         Number `json:"azn2"`
	Bhd1         Number `json:"bhd1"`
	Bhd2         Number `json:"bhd2"`
	Bitcoin      Number `json:"bitcoin"`
	Bourse       Number `json:"bourse"`
	Cad1         Number `json:"cad1"`
	Cad2         Number `json:"cad2"`
	Chf1         Number `json:"chf1"`
	Chf2         Number `json:"chf2"`
	Cny1         Number `json:"cny1"`
	Cny2         Number `json:"cny2"`
	Created      string `json:"created"`
	Day          int    `json:"day"`
	Dkk1         Number `json:"dkk1"`
	Dkk2         Number `json:"dkk2"`
	Emami1       Number `json:"emami1"`
	Emami12      Number `json:"emami12"`
	Eur1         Number `json:"eur1"`
	Eur2         Number `json:"eur2"`
	Gbp1         Number `json:"gbp1"`
	Gbp2         Number `json:"gbp2"`
	Gol18        Number `json:"gol18"`
	Hkd1         Number `json:"hkd1"`
	Hkd2         Number `json:"hkd2"`
	Hour         string `json:"hour"`
	Inr1         Number `json:"inr1"`
	Inr2         Number `json:"inr2"`
	Iqd1         Number `json:"iqd1"`
	Iqd2         Number `json:"iqd2"`
	Jpy1         Number `json:"jpy1"`
	Jpy2         Number `json:"jpy2"`
	Kwd1         Number `json:"kwd1"`
	Kwd2         Number `json:"kwd2"`
	LastModified string `json:"last_modified"`
	Minute       string `json:"minute"`
	Mithqal      Number `json:"mithqal"`
	Month        int    `json:"month"`
	Myr1         Number `json:"myr1"`
	Myr2         Number `json:"myr2"`
	Nok1         Number `json:"nok1"`
	Nok2         Number `json:"nok2"`
	Omr1         Number `json:"omr1"`
	Omr2         Number `json:"omr2"`
	Ounce        Number `json:"ounce"`
	Qar1         Number `json:"qar1"`
	Qar2         Number `json:"qar2"`
	Rub1         Number `json:"rub1"`
	Rub2         Number `json:"rub2"`
	Sar1         Number `json:"sar1"`
	Sar2         Number `json:"sar2"`
	Second       string `json:"second"`
	Sek1         Number `json:"sek1"`
	Sek2         Number `json:"sek2"`
	Sgd1         Number `json:"sgd1"`
	Sgd2         Number `json:"sgd2"`
	Thb1         Number `json:"thb1"`
	Thb2         Number `json:"thb2"`
	Try1         Number `json:"try1"`
	Try2         Number `json:"try2"`
	Usd1         Number `json:"usd1"`
	Usd2         Number `json:"usd2"`
	Weekday      string `json:"weekday"`
	Year         int    `json:"year"`
}

type Number float64

func (c Number) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(strconv.FormatFloat(float64(c), 'f', 0, 64))), nil
}

func (c *Number) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*c = Number(i)
	return nil
}

func (c Number) Float64() float64 {
	return float64(c)
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
