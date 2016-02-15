/*
	create	(name, length)
	add 	(bucket, time, data)
	get 	(bu)
*/

package main

import (
	"github.com/mediocregopher/radix.v2/redis"
	"fmt"
)

var (
	client *redis.Client
)


func main() {
	client, e := redis.Dial("tcp", "localhost:6379")
	checkerr(e, "Connecting to redis", true)

	e = CreateBucket(client, "somebucket", 10000000)
	checkerr(e, "Creating bucket", true)

	bucket, err := GetBucket(client, "somebucket")

	if err != nil {
		checkerr(err, "Getting bucket", false)
		return
	}


	bucket.Add(239, "Ssdfjdfsdf")
	fmt.Printf("%#v\n", bucket);
	fmt.Printf("%#v\n", bucket.Get());
}
