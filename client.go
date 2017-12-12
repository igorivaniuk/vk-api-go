package vk_api_go

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"github.com/drossha/vk-api-go/support"
	"bytes"
)

const BaseApiUrl = "https://api.vk.com/method/"
const ApiVersion = "5.69"

const (
	ApiLanguageRu = "ru"
	ApiLanguageUa = "ua"
	ApiLanguageBe = "be"
	ApiLanguageEn = "en"
	ApiAlnguageEs = "es"
	ApiLanguageFi = "fi"
	ApiLanguageDe = "de"
	ApiLanguageIt = "it"
)

type ApiRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type ClientInterface interface {
	// Do some api request
	Do() ([]byte, error)
}

type Client struct {
	ApiUrl       string
	ApiVersion   string
	AccessToken  string
	Language     string
	ClientID     string
	ClientSecret string
	TestMode     bool
	UseHttps     bool
	UserXml      bool
	HttpClient   http.Client
}

func NewClient() Client {
	return Client{
		ApiUrl:     BaseApiUrl,
		Language:   ApiLanguageRu,
		ApiVersion: ApiVersion,
		TestMode:   false,
		UseHttps:   true,
		HttpClient: http.Client{},
	}
}

func (c *Client) Do(m string, r interface{}) ([]byte, error) {

	urlMethod := c.ApiUrl + m

	if c.UserXml {
		urlMethod += ".xml"
	}

	params := support.Map(r)

	form := url.Values{}
	for k, v := range params {
		if v != "" && v != "0" {
			form.Add(k, v)
		}
	}
	if c.TestMode {
		form.Add("test_mode", "1")
	}
	if c.UseHttps {
		form.Add("https", "1")
	}

	form.Add("lang", c.Language)
	form.Add("v", c.ApiVersion)

	body := bytes.NewBufferString(form.Encode())
	resp, err := http.Post(urlMethod, "application/x-www-form-urlencoded", body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
