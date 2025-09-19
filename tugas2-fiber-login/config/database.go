package config

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/misbahkun/test-fullstack-2025/tugas2-fiber-login/model"
	"github.com/redis/go-redis/v9"
)

var backgroundContext = context.Background()

func ConnectToRedis() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",              
		DB:       0,               
	})

	if err := redisClient.Ping(backgroundContext).Err(); err != nil {
		log.Printf("Gagal melakukan ping ke Redis: %v", err)
		return nil, err
	}

	log.Println("Berhasil terhubung ke database Redis.")
	return redisClient, nil
}

func SeedDatabaseWithDummyUser(redisClient *redis.Client) {
	passwordHasher := sha1.New()
	passwordHasher.Write([]byte("password123"))
	passwordHashString := hex.EncodeToString(passwordHasher.Sum(nil))

	dummyUserData := model.UserRedisData{
		RealName: "Aberto Doni Sianturi",
		Email:    "adss@gmail.com",
		Password: passwordHashString,
	}

	dummyUserJSON, err := json.Marshal(dummyUserData)
	if err != nil {
		log.Fatalf("Gagal marshal data dummy: %v", err)
	}

	err = redisClient.Set(backgroundContext, "login_doni", dummyUserJSON, 0).Err()
	if err != nil {
		log.Printf("Gagal menyimpan data dummy ke Redis: %v", err)
	} else {
		log.Println("Data dummy 'login_doni' berhasil disimpan ke Redis.")
	}
}