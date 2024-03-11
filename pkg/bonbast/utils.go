package bonbast

import (
	"net/http"
	"strconv"
)

// Number holds a float64. Used to Unmarshal a string to float.
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

func (c Number) Float64() float64 { return float64(c) }

func setHeaders(req *http.Request) {
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("accept-language", "en-US,en-GB;q=0.9,en;q=0.8,fa;q=0.7")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"122\", \"Not(A:Brand\";v=\"24\", \"Google Chrome\";v=\"122\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("cookie", "cookieconsent_status=true; st_bb=0")
	req.Header.Set("Referer", HOST)
	req.Header.Set("Referrer-Policy", "strict-origin-when-cross-origin")
}
