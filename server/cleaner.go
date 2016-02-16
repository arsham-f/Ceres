package main

import (
	"github.com/mediocregopher/radix.v2/redis"
	ceres "github.com/arsham-f/Ceres/ceres"
	"time"
)

func cleaner(client *redis.Client) {
	for {
		buckets, _ := ceres.AllBuckets(client)

		for _, b := range buckets {
			b.Purge()
		}

		time.Sleep(time.Second)
	}
}