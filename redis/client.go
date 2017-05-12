package Redis

import (
	//"fmt"
	"time"

	"github.com/fzzy/radix/redis"
)

type Client struct {
	client *redis.Client
}

func NewClient(network, addr, pass string, timeout time.Duration) (*Client, error) {
	c, e := redis.DialTimeout(network, addr, timeout)
	if e != nil {
		return nil, e
	}

	this := &Client{
		client: c,
	}
	r := this.client.Cmd("AUTH", pass)
	if r.Err != nil {
		this.client.Close()

		return nil, r.Err
	}

	return this, nil
}

func (this *Client) Close() {
	this.client.Close()
}

func (this *Client) Hset(key, field, value string) error {
	r := this.client.Cmd("HSET", key, field, value)
	if r.Err != nil {
		return r.Err
	}

	return nil
}

func (this *Client) Hmset(key string, fieldValues ...string) error {
	args := []interface{}{key}
	for i := 0; i < len(fieldValues); i++ {
		args = append(args, fieldValues[i])
	}

	r := this.client.Cmd("HMSET", args...)
	if r.Err != nil {
		return r.Err
	}

	return nil
}

func (this *Client) Hget(key, field string) (string, error) {
	r := this.client.Cmd("HGET", key, field)
	if r.Err != nil {
		return "", r.Err
	}

	str, err := r.Str()
	if err != nil {
		return "", err
	}
	return str, nil
}

func (this *Client) Hgetall(key string) (map[string]string, error) {
	r := this.client.Cmd("HGETALL", key)
	if r.Err != nil {
		return nil, r.Err
	}

	hash, err := r.Hash()
	if err != nil {
		return nil, err
	}
	return hash, nil
}