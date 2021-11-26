package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

func HTTPGet(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func HTTPPostForm(uri string, params url.Values) ([]byte, error) {
	client := &http.Client{}
	param := strings.NewReader(params.Encode()) //把form数据编码

	request, err := http.NewRequest(http.MethodPost, uri, param)
	if err != nil {
		logrus.Info(err.Error(), `请求失败`)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, errors.New(fmt.Sprintf(`STATUS ERROR:%s`, response.Status))
}

func HTTPPostJson(uri string, data interface{}) ([]byte, error) {
	client := &http.Client{}

	dt, err := json.Marshal(data)
	if err != nil {
		logrus.Info("Parse Json Error")
		return nil, err
	}
	reqNew := bytes.NewBuffer(dt)
	request, err := http.NewRequest(`POST`, uri, reqNew)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, errors.New(fmt.Sprintf(`STATUS ERROR:%s`, response.Status))
}
