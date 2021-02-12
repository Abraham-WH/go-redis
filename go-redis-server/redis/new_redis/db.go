package new_redis

type Database struct {
	data       map[Object]*Object
	expireTime map[Object]int64
}
