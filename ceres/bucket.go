package ceres


import (
	"github.com/mediocregopher/radix.v2/redis"
)

type Bucket struct {
	client 	*redis.Client
	name	string
	length	int
}

func (b *Bucket) bucketName() string {
	return "bucket_" + b.name
}

func (b *Bucket) Add(time int, data string) error {
	return b.client.Cmd("ZADD", b.bucketName(), time, data).Err
}

func (b *Bucket) Get() []Entry {
	var r []Entry

	resp := b.client.Cmd("ZRANGE", b.bucketName(), "0", "-1", "WITHSCORES")

	if resp.Err != nil {
		return nil
	}

	if resp.IsType(redis.Nil) {
		return nil
	}

	result, _ := resp.Array()

	for i := 0; i < len(result); i += 2 {
		v1, _ := result[i].Str()
		v2, _ := result[i+1].Str()

		r = append(r, Entry { 
			atoi(v2), v1,
		})
	}

	return r
}

func CreateBucket(client *redis.Client, name string, length int) error {
	return client.Cmd("SET", name, length, "NX").Err
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