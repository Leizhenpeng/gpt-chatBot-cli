package services

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type CacheHistory struct {
	db *cache.Cache
}

func (c CacheHistory) get(key string) (interface{}, bool) {
	res, ok := c.db.Get(key)
	if !ok {
		return "", false
	}
	return res, true
}

func (c CacheHistory) set(key string, value interface{}) {
	c.db.Set(key, value, time.Minute*10)
}

func (c CacheHistory) clear() {
	c.db.Flush()
}

type CacheHistoryInterface interface {
	get(key string) (interface{}, bool)
	set(key string, value interface{})
	clear()

	SetQACache(q string, a string)
	GetQACache() (string, bool)
	ClearQACache()
}

const (
	userSessionContextKey = "userSessionContext"
)

func NewCacheHistory() CacheHistoryInterface {
	return &CacheHistory{
		db: cache.New(time.Second*60, time.Minute*10),
	}
}

func (c CacheHistory) SetQACache(q string, a string) {
	saveValue := q + "\n" + a
	c.set(userSessionContextKey, saveValue)
}

func (c CacheHistory) GetQACache() (string, bool) {
	res, ok := c.get(userSessionContextKey)
	if !ok {
		return "", false
	}
	return res.(string), true
}

func (c CacheHistory) ClearQACache() {
	c.clear()
}
