package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestAllPage(t *testing.T) {
	baseUrl:="http://localhost:8080"
	//声明加初始化测试数据
	var tests=[]struct{
		method string
		url string
		expected int
	}{
		{"GET","/",200},
		{"GET","/about",200},
		{"GET","/notfound",404},
		{"GET","/articles",200},
		{"GET","/articles/create",200},
		{"GET","/articles/12",200},
		{"GET","/articles/12/edit",200},
		{"POST","/articles/12",200},
		{"POST","/articles",200},
		{"POST","/articles/12/delete",200},
	}
	for _, test := range tests {
		t.Logf("当前请求URL:%v \n",test.url)
		var(
			resp *http.Response
			err error
		)
		switch{
		case test.method=="POST":
			data:=make(map[string][]string)
			resp,err=http.PostForm(baseUrl+test.url,data)
		default:
			resp,err=http.Get(baseUrl+test.url)
		}
		assert.NoError(t, err,"请求"+test.url+"时报错")
		assert.Equal(t, test.expected,resp.StatusCode,test.url+"应返回状态码"+strconv.Itoa(test.expected))
	}
}