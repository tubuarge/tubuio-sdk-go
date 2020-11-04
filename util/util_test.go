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
	}{
		//everything is present
		{base_url + "/int/shortId1/method1/tag1?account=account1&args=1&args=2&args=%40&args=true",
			"shortId1", "method1", "tag1", "account1"},
		//tag is missing
		{base_url + "/int/shortId2/method2?account=account2&args=1&args=2&args=%40&args=true",
			"shortId2", "method2", "", "account2"},
		//account is missing
		{base_url + "/int/shortId3/method3/tag3?args=1&args=2&args=%40&args=true",
			"shortId3", "method3", "tag3", ""},
		//args is missing
		{base_url + "/int/shortId4/method4/tag4?account=account4&args=1&args=2&args=%40&args=true",
			"shortId4", "method4", "tag4", "account4"},
		//all optional params are missing
		{base_url + "/int/shortId5/method5?args=1&args=2&args=%40&args=true",
			"shortId5", "method5", "", ""},
		//tag and account are missing
		{base_url + "/int/shortId6/method6?account=account6&args=1&args=2&args=%40&args=true",
			"shortId6", "method6", "", "account6"},
	}

	for _, table := range tables {
		gotUrl := GetHttpGetUrl(base_url, table.shortId, table.method, table.tag, table.account, []interface{}{"1", 2, "@", true})
		if strings.Compare(gotUrl, table.url) != 0 {
			t.Errorf("Got: '%s', Want: '%s'", gotUrl, table.url)
		}
	}
}

func TestGetHttpPostUrl(t *testing.T) {
	tables := []struct {
		url     string
		shortId string
		method  string
		tag     string
	}{
		{base_url + "/int/shortId1/method1/tag1",
			"shortId1", "method1", "tag1"},
		{base_url + "/int/shortId2/method2",
			"shortId2", "method2", ""},
		{base_url + "/int/shortId3/method3/tag3",
			"shortId3", "method3", "tag3"},
		{base_url + "/int/shortId4/method4/tag4",
			"shortId4", "method4", "tag4"},
		{base_url + "/int/shortId5/method5",
			"shortId5", "method5", ""},
		{base_url + "/int/shortId6/method6",
			"shortId6", "method6", ""},
	}

	for _, table := range tables {
		gotUrl := GetHttpPostUrl(base_url, table.shortId, table.method, table.tag)
		if strings.Compare(gotUrl, table.url) != 0 {
			t.Errorf("Got: '%s', Want: '%s'", gotUrl, table.url)
		}
	}
}

func TestGetBodyRequest(t *testing.T) {
	tables := []struct {
		wantJson    string
		account     string
		args		[]interface{}
	}{
		{"{\"args\":[1,true,\"args1\",\"@\"],\"account\":\"account1\"}", "account1",[]interface{}{1, true, "args1", "@"}},
	}

	for _, table := range tables {
		gotJson, err := GetBodyRequest(table.account, table.args)
		if err != nil {
			t.Error(err)
		}

		if strings.Compare(table.wantJson, string(gotJson)) != 0 {
			t.Errorf("Got: '%s', Want: '%s'", string(gotJson), table.wantJson)
		}
	}
}
