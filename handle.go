package yuansfer

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

func struct2Map(obj Yuansfer) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		data[strings.TrimSuffix(
			t.Field(i).Tag.Get("json"),
			",omitempty")] = v.Field(i).String()
	}
	return data
}

func md5Token(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	md5Data := hash.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}

func getSignature(req Yuansfer) string {
	//req.LoadCredentials()
	data := struct2Map(req)
	pre := map2Str(data) + md5Token(Token)
	log.Println(pre)
	return md5Token(pre)
}

func map2Str(m map[string]string) string {
	var keys []string
	for k := range m {
		if m[k] != "" {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)

	var buf bytes.Buffer
	for _, key := range keys {
		buf.WriteString(key)
		buf.WriteString("=")
		buf.WriteString(m[key])
		buf.WriteString("&")
	}

	return buf.String()
}

func postToYuansfer(uri string, req Yuansfer) ([]byte, error) {
	var (
		buf []byte
		err error
	)
	buf, err = json.Marshal(req)
	if nil != err {
		return nil, err
	}
	c := NewHttpClient(uri, buf)
	err = c.Request()
	return c.Body, err
}

func values2Map(m url.Values) map[string]string {
	var r = make(map[string]string)
	m.Del("verifySign")
	for k := range m {
		r[k] = m[k][0]
	}
	return r
}

//VerifySignNotify checks the parameters from Yuansfer with the value of verifySign.
func VerifySignNotify(values url.Values, token string) (m map[string]string, r bool) {
	verifySign := values.Get("verifySign")

	m = values2Map(values)
	pre := map2Str(m) + md5Token(token)
	vs := md5Token(pre)

	return m, vs == verifySign
}
