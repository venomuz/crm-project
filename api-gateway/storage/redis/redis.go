package redis

import (
	redis "github.com/gomodule/redigo/redis"
	"github.com/venomuz/crm-project/api-gateway/storage/repo"
)

type redisRepo struct {
	rConn *redis.Pool
}

func NewRedisRepo(rds *redis.Pool) repo.RepositoryStorage {
	return &redisRepo{
		rConn: rds,
	}

}

func (r *redisRepo) Set(key, value string) error {
	conn := r.rConn.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	return err
}

// SetWithTTL

func (r *redisRepo) SetWithTTL(key, value string, second int64) error {
	conn := r.rConn.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, second, value)
	return err
}

//Get
func (r *redisRepo) Get(key string) (interface{}, error) {
	conn := r.rConn.Get()
	defer conn.Close()

	return conn.Do("GET", key)
}
func (r *redisRepo) Delete(key string) (interface{}, error) {
	conn := r.rConn.Get()
	defer conn.Close()

	return conn.Do("DEL", key)
}
func (r *redisRepo) Search(key string) (interface{}, error) {
	conn := r.rConn.Get()
	defer conn.Close()

	return conn.Do("INCR", key)
}
