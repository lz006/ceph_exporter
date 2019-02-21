package rediscom

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {

	// Access values ret from config file
	redisEndpoint := os.Getenv("REDIS_ENDPOINT")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	client = redis.NewClient(&redis.Options{
		Addr: redisEndpoint,
		DB:   redisDB, // use default DB
	})
}

func GetImagesFromRedis(poolName string) ([]*Image, error) {

	val, err := client.Get(poolName).Result()
	if err != nil {
		log.Println("Error when trying to red from redis: " + err.Error())
		return nil, err
	}

	var images []*Image
	err = json.Unmarshal([]byte(val), &images)
	if err != nil {
		log.Println("Error while trying to unmarshall result from redis: " + err.Error())
		return nil, err
	}

	return images, nil

}

type Image struct {
	Pool string  `json:"pool,omitempty"`
	Name string  `json:"name,omitempty"`
	Size float64 `json:"size,omitempty"`
	Used float64 `json:"used,omitempty"`
}
