//go:build !debug

package store

import (
	"context"
	"fmt"
	"github.com/amirkhaki/cnbzdtr/config"
	"github.com/amirkhaki/cnbzdtr/entity"
	"github.com/amirkhaki/cnbzdtr/protocol"
	"github.com/go-redis/redis/v8"
)

type redisStore struct {
	rdb *redis.Client
}

func (rS *redisStore) AddUser(ctx context.Context, u *entity.User) error {
	if u.ID == "" {
		return fmt.Errorf("User ID is empty!")
	}
	if _, err := rS.rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		data := make(map[string]any)
		data["score"] = u.Score
		data["most_score"] = u.MostScore
		data["prev_most_score"] = u.PrevMostScore
		data["referral_count"] = u.ReferralCount
		data["message_count"] = u.MessageCount
		for k, v := range data {
			if err := rdb.HSet(ctx, u.ID, k, v).Err(); err != nil {
				return fmt.Errorf("Could not set key %s: %w", k, err)
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("could not set user data pipelined: %w", err)
	}
	return nil
}
func (rS *redisStore) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	u, err := entity.NewUser(id)
	if err != nil {
		return nil, fmt.Errorf("Could not create user with id: %w", err)
	}
	if err = rS.rdb.HGetAll(ctx, u.ID).Scan(u); err != nil {
		return nil, fmt.Errorf("Could not get user data from redis: %w", err)
	}
	return u, nil
}
func (rS *redisStore) GetUserOrCreate(ctx context.Context, id string) (*entity.User, error) {
	if u, err := rS.GetUserByID(ctx, id); err == nil {
		return u, nil
	} else {
		u, err = entity.NewUser(id)
		if err != nil {
			return nil, fmt.Errorf("Could not create user with id: %w", err)
		}
		return u, rS.AddUser(ctx, u)
	}
}

func (rS *redisStore) UpdateUser(ctx context.Context, u *entity.User) error {
	return rS.AddUser(ctx, u)
}
func (rS *redisStore) DeleteUser(ctx context.Context, u *entity.User) error {
	err := rS.rdb.Del(ctx, u.ID).Err()
	if err != nil {
		return fmt.Errorf("Could not delete user: %w", err)
	}
	return nil
}
func New(cfg config.Config) (protocol.Store, error) {
	opt, err := redis.ParseURL(cfg.Redis_DSN)
	if err != nil {
		return nil, fmt.Errorf("Could not parse redis dsn: %w", err)
	}
	rdb := redis.NewClient(opt)
	rS := &redisStore{rdb: rdb}
	return rS, nil
}
