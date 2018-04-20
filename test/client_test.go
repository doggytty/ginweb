package test

import (
	"testing"
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)
// 用于读取resp的body
func helpRead(resp *http.Response)  {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}
func TestGet(t *testing.T) {
	// 调用最基本的GET,并获得返回值
	resp,_ := http.Get("http://0.0.0.0:8888/test1")
	helpRead(resp)
}

func TestPost(t *testing.T)  {
	// 调用最基本的POST,并获得返回值
	resp,_ := http.Post("http://0.0.0.0:8888/test2", "",  strings.NewReader(""))
	helpRead(resp)
}

func TestParamTest(t *testing.T) {
	// 调用最基本的GET,并获得返回值
	resp,_ := http.Get("http://0.0.0.0:8888/sunlichuan/lucien")
	helpRead(resp)
}

