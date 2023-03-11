package redis
//
//import (
//	"context"
//	"fmt"
//	log2 "github.com/huaixiaohai/lib/log"
//"time"
//
//	"github.com/go-redis/redis/v9"
//	)
//
//const KeepTTL = -1
//
//type Config struct {
//	Address      string
//	Password     string
//	DB           int
//	MaxIdleConns int
//	PoolSize     int
//}
//
//var client *redis.Client
//
//func Init(conf *Config) error {
//	client = redis.NewClient(&redis.Options{
//		Addr:         conf.Address,
//		Password:     conf.Password,
//		DB:           conf.DB,
//		MaxIdleConns: conf.MaxIdleConns,
//		PoolSize:     conf.PoolSize, // 连接池
//	})
//	pong, err := client.Ping(context.Background()).Result()
//	if err != nil {
//		//panic(fmt.Sprintf("redis 连接失败, %s, %s", pong, err.Error()))
//		log2.Error(fmt.Sprintf("redis 连接失败, %s, %s", pong, err.Error()))
//		return err
//	}
//	log2.Info("redis 连接成功：", pong)
//	return nil
//}
//
//func Set(ctx context.Context, key, value string, expiration time.Duration) error {
//	return client.Set(ctx, key, value, expiration).Err()
//}
//
//func Get(ctx context.Context, key string) (value string, has bool, err error) {
//	value, err = client.Get(ctx, key).Result()
//	if err == redis.Nil {
//		return value, false, nil
//	}
//	if err != nil {
//		return value, false, err
//	}
//	return value, true, nil
//}
//
//func Del(ctx context.Context, key string) error {
//	return client.Del(ctx, key).Err()
//}
