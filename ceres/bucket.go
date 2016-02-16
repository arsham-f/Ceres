package ceres


import (
	"github.com/mediocregopher/radix.v2/redis"
	"errors"
	"time"
)

type Bucket struct {
	client 	*redis.Client
	Name	string
	length	int
}

func (b *Bucket) bucketName() string {
	return "bucket_" + b.Name
}

func (b *Bucket) Add(time int, data string) error {
	return b.client.Cmd("ZADD", b.bucketName(), time, data).Err
}

func (b *Bucket) Get() ([]Entry, error) {
	var r []Entry

	resp := b.client.Cmd("ZRANGE", b.bucketName(), "0", "-1", "WITHSCORES")

	if resp.Err != nil {
		return nil, resp.Err
	}

	if resp.IsType(redis.Nil) {
		return []Entry{}, nil
	}

	result, _ := resp.Array()

	for i := 0; i < len(result); i += 2 {
		v1, _ := result[i].Str()
		v2, _ := result[i+1].Str()

		r = append(r, Entry { 
			atoi(v2), v1,
		})
	}

	return r, nil
}

func (b *Bucket) Purge() error {
	if b.length == 0 {
		return nil
	}
	
	threshold := time.Now().UTC().Unix() - int64(b.length)
	return b.client.Cmd("ZREMRANGEBYSCORE", b.bucketName(), "-inf", threshold).Err
}

func (b *Bucket) Delete() error {
	return b.client.Cmd("DEL", b.Name, b.bucketName()).Err
}

func CreateBucket(client *redis.Client, name string, length int) error {
	exists, err := client.Cmd("EXISTS", name).Int()


	if err != nil {
		return err
	}

	if exists == 1 {
		return errors.New("Bucket already exists")
	} else {
		return client.Cmd("SET", name, length, "NX").Err
	}
}

func GetBucket(client *redis.Client, name string) (*Bucket, error) {
	length, err := client.Cmd("GET", name).Str()
	if err != nil {
		return nil, err
	}

	b := new(Bucket)
	*b = Bucket {
		client,
		name,
		atoi(length),
	}

	return b, nil
}

func AllBuckets(client *redis.Client) ([]*Bucket, error) {
	var ret []*Bucket

	keys, err := client.Cmd("KEYS", "*").Array()

	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		keyval, err := key.Str()

		if err != nil {
			continue
		}

		b, err := GetBucket(client, keyval)

		if err != nil {
			continue
		}

		ret = append(ret, b)
	}

	return ret, nil
}