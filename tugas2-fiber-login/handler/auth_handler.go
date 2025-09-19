package handler

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/misbahkun/test-fullstack-2025/tugas2-fiber-login/model"
	"github.com/redis/go-redis/v9"
)

var requestContext = context.Background()

type AuthHandler struct {
	RedisClient *redis.Client
}

func NewAuthHandler(client *redis.Client) *AuthHandler {
	return &AuthHandler{RedisClient: client}
}

func (handler *AuthHandler) HandleLogin(fiberContext *fiber.Ctx) error {
	
	var loginPayload model.LoginRequestBody

	if err := fiberContext.BodyParser(&loginPayload); err != nil {
		log.Printf("Gagal mem-parsing body request: %v", err)
		return fiberContext.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message":  "Format request tidak valid.",
		})
	}

	userRedisKey := "login_" + loginPayload.Username

	redisDataString, err := handler.RedisClient.Get(requestContext, userRedisKey).Result()
	
	if err == redis.Nil {
		log.Printf("Percobaan login gagal: User '%s' tidak ditemukan.", loginPayload.Username)
		return fiberContext.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"message":  "Username atau password salah.",
		})
	} else if err != nil {
		log.Printf("Error saat mengambil data dari Redis: %v", err)
		return fiberContext.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message":  "Terjadi kesalahan internal pada server.",
		})
	}

	var userFromRedis model.UserRedisData
	if err := json.Unmarshal([]byte(redisDataString), &userFromRedis); err != nil {
		log.Printf("Gagal unmarshal data Redis untuk user '%s': %v", loginPayload.Username, err)
		return fiberContext.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message":  "Gagal memproses data pengguna.",
		})
	}

	passwordHasher := sha1.New()
	passwordHasher.Write([]byte(loginPayload.Password))
	hashedInputPassword := hex.EncodeToString(passwordHasher.Sum(nil))

	if hashedInputPassword != userFromRedis.Password {
		log.Printf("Percobaan login gagal: Password salah untuk user '%s'.", loginPayload.Username)
		return fiberContext.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"message":  "Username atau password salah.",
		})
	}

	log.Printf("User '%s' berhasil login.", loginPayload.Username)
	return fiberContext.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login berhasil!",
		"user": fiber.Map{
			"realname": userFromRedis.RealName,
			"email":    userFromRedis.Email,
		},
	})
}