# Ceres

Ceres, pronounced "series," is a time-series data store. It's an abstraction of a key-value store where your key is always some increasing integer (i.e. a timestamp) and your values will be returned to you in ascending order of key. 

You can create "buckets" to store these keys. Each bucket has a lifetime. Entries which are too old will be pruned if their key is too old.

This is all built on Redis ZSET.


