package redisstorage

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Dnlbb/telegram-bot/internal/config"
	"github.com/Dnlbb/telegram-bot/internal/storage"
	"github.com/go-redis/redis/v8"
)

var (
	ErrNoSavedPages = errors.New("no saved pages")
)

type Storage struct {
	client *redis.Client
}

func New(addr string) *Storage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	return &Storage{client: rdb}
}

func (s *Storage) Save(page *storage.Page) error {
	ctx := context.Background()

	hash, err := page.Hash()
	if err != nil {
		return fmt.Errorf("failed to create hash: %w", err)
	}

	key := fmt.Sprintf("user:%s:page:%s", page.UserName, hash)
	err = s.client.Set(ctx, key, page.URL, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to save page to Redis: %w", err)
	}

	return nil
}

func (s *Storage) PickRandom(userName string) (*storage.Page, error) {
	ctx := context.Background()

	pattern := fmt.Sprintf("user:%s:page:*", userName)
	keys, err := s.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get keys: %w", err)
	}

	if len(keys) == 0 {
		return nil, ErrNoSavedPages
	}

	rand.Seed(time.Now().UnixNano())
	randomKey := keys[rand.Intn(len(keys))]

	url, err := s.client.Get(ctx, randomKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get page: %w", err)
	}

	return &storage.Page{
		URL:      url,
		UserName: userName,
	}, nil
}

func (s *Storage) Remove(page *storage.Page) error {
	ctx := context.Background()

	hash, err := page.Hash()
	if err != nil {
		return fmt.Errorf("failed to create hash: %w", err)
	}

	key := fmt.Sprintf("user:%s:page:%s", page.UserName, hash)
	err = s.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to remove page: %w", err)
	}

	return nil
}

func (s *Storage) IsExists(page *storage.Page) (bool, error) {
	ctx := context.Background()

	hash, err := page.Hash()
	if err != nil {
		return false, fmt.Errorf("failed to create hash: %w", err)
	}

	key := fmt.Sprintf("user:%s:page:%s", page.UserName, hash)
	exists, err := s.client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}

	return exists > 0, nil
}
