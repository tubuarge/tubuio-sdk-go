package util

import (
	"strings"
	"testing"
)

const (
	base_url = "https://prodservice-dot-dynamic-sun-260208.appspot.com"
)

func TestGetHttpGetUrl(t *testing.T) {
	tables := []struct {
		url     string
		shortId string
		method  string
		tag     string
		account string
		args    []string
	}{
		//everything is present
		{base_url+"/int/shortId1/method1/tag1?account=account1&args=arg1&args=arg2",
			"shortId1", "method1", "tag1", "account1", []string{"arg1", "arg2"}},
		//tag is missing
		{base_url+"/int/shortId2/method2?account=account2&args=arg1&args=arg2",
			"shortId2", "method2", "", "account2", []string{"arg1", "arg2"}},
		//account is missing
		{base_url+"/int/shortId3/method3/tag3?args=arg1&args=arg2",
			"shortId3", "method3", "tag3", "", []string{"arg1", "arg2"}},
		//args is missing
		{base_url+"/int/shortId4/method4/tag4?account=account4",
			"shortId4", "method4", "tag4", "account4", nil},
		//all optional params are missing
		{base_url+"/int/shortId5/method5",
			"shortId5", "method5", "", "", nil},
		//tag and account are missing
		{base_url+"/int/shortId6/method6?account=account6&args=arg1&args=arg2",
			"shortId6", "method6", "", "account6", []string{"arg1", "arg2"}},
	}

	for _, table := range tables{
		gotUrl := GetHttpGetUrl(base_url, table.shortId, table.method, table.tag, table.account, table.args)
		if strings.Compare(gotUrl, table.url) != 0 {
			t.Errorf("Got: '%s', Want: '%s'", gotUrl, table.url)
		}
	}

}
