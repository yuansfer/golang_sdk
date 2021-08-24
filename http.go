package yuansfer

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var network = Network{
	RequestTimeout: 5,
	ConnectTimeout: 5,
	SocketTimeout:  5,
}

var client = &http.Client{
	Transport: &http.Transport{
		DialContext: func(ctx context.Context, n, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(n, addr,
				time.Second*network.RequestTimeout,
			)
			if err != nil {
				return conn, err
			}
			_ = conn.SetDeadline(time.Now().
				Add(time.Second * network.ConnectTimeout))
			return conn, err
		},
		ResponseHeaderTimeout: time.Second * network.SocketTimeout,
		MaxConnsPerHost:       0,
		IdleConnTimeout:       0,
		MaxIdleConns:          0,
	},
}

type Client struct {
	URI        string
	HTTPStatus int
	// Body indicates both Request Body & Response Body
	Body []byte
}

func NewHttpClient(uri string, b []byte) *Client {
	return &Client{
		URI:  uri,
		Body: b,
	}
}

func (c *Client) Request() (err error) {
	request, err := http.NewRequest(http.MethodPost,
		APIGateway+c.URI,
		bytes.NewReader(c.Body),
	)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("User-Agent", fmt.Sprintf("go-sdk: %s", Version()))

	resp, err := client.Do(request)
	if nil == resp {
		err = fmt.Errorf("none response received")
		return
	}

	if http.StatusOK != resp.StatusCode {
		err = fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	c.HTTPStatus = resp.StatusCode
	c.Body, err = ioutil.ReadAll(resp.Body)

	return
}
