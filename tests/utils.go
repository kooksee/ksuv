package tests

import (
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
)


type Request struct {
	Prefix string
	Url    string
	Method string
	Token  string
	Data   string
}

type Response struct {
	*http.Response
	Body string
}

func NewRequest(prefix, token string) *Request {
	return &Request{Prefix: prefix, Token: token}
}

func (req *Request) Post(url string, data string) (resp *Response) {
	req.Url = url
	req.Method = "POST"
	req.Data = data

	return req.Do()
}

func (req *Request) Get(url string, data string) (resp *Response) {
	req.Url = url
	req.Method = "GET"
	req.Data = data

	return req.Do()
}

func (req *Request) Do() (resp *Response) {
	r, err := http.NewRequest(req.Method, req.Prefix+req.Url, bytes.NewBuffer([]byte(req.Data)))
	if err != nil {
		log.Fatal("creating NewRequest failed", err)
	}

	r.Header.Set("Content-Type", "application/json")
	if req.Token != "" {
		r.Header.Set("Authorization", "Bearer "+req.Token)
	}

	re, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal("client do failed", err)
	}

	//check if status ok
	if re.StatusCode != http.StatusOK {
		log.Fatalf("net status: %d, check gateway and rpc", re.StatusCode)
	}

	defer re.Body.Close()
	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		log.Fatal("this may never happened", err)
	}

	resp = &Response{re, string(body)}
	return resp
}

func (resp *Response) GetStringValue(key string) string {
	result := make(map[string]interface{})
	json.Unmarshal([]byte(resp.Body), &result)
	data := result["data"].(map[string]interface{})
	return data[key].(string)
}

func (resp *Response) GetFloat64Value(key string) float64 {
	result := make(map[string]interface{})
	json.Unmarshal([]byte(resp.Body), &result)
	data := result["data"].(map[string]interface{})
	return data[key].(float64)
}

func (resp *Response) JsonBody() *simplejson.Json {
	j, err := simplejson.NewJson([]byte(resp.Body))
	if err != nil {
		log.Fatal(err)
	}
	return j
}

