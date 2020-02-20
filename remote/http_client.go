package remote

import (
	"io/ioutil"
	"net/http"
)

var client = http.DefaultClient

func Get(url string,params map[string]string) ([]byte,error){

	req,err := http.NewRequest(http.MethodGet,url,nil)
	if err != nil {
		return nil,err
	}

	if params != nil && len(params) > 0 {
		qValues := req.URL.Query()
		for k,v:= range params {
			qValues.Add(k,v)
		}
		req.URL.RawQuery = qValues.Encode()
	}

	resp,err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer closeResp(resp)

	var b []byte
	if resp != nil {
		b,err = ioutil.ReadAll(resp.Body)
	}

	return b,err

}

func closeResp(resp *http.Response) {
	if resp != nil {
		resp.Body.Close()
	}
}