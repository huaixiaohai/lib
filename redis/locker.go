package redis
//
//import (
//	"context"
//	"errors"
//	"time"
//
//	"github.com/bsm/redislock"
//)
//
//var ErrNotObtained = errors.New("redislock: not obtained")
//
//func NewLocker() *Locker {
//	return &Locker{
//		ctx:    context.Background(),
//		client: redislock.New(client),
//	}
//}
//
//// Locker redis.Init(conf)后才可使用
//type Locker struct {
//	ctx    context.Context
//	client *redislock.Client
//	lock   *redislock.Lock
//}
//
//func (a *Locker) Lock(key string, ttl time.Duration) error {
//	var err error
//	a.lock, err = a.client.Obtain(a.ctx, key, ttl, nil)
//	if err == redislock.ErrNotObtained {
//		return ErrNotObtained
//	} else if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (a *Locker) Unlock() error {
//	return a.lock.Release(a.ctx)
//}
