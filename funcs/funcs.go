package funcs

import (
	"cache/structures"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v2"
)

const link = "https://jsonplaceholder.typicode.com/users"

var redisClient *redis.Client

func RetrieveAll() ([]structures.User, error) {
	var Users []structures.User
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	data, _ := redisClient.Get(ctx, link).Result() //retrieving data from cache
	if data != "" {                                //take data from cache
		json.Unmarshal([]byte(data), &Users)
		return Users, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return Users, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return Users, err
	}

	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("Status code %d\n", resp.StatusCode)
		fmt.Println(err)
		return Users, errors.New(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return Users, err
	}
	if err = json.Unmarshal(body, &Users); err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return Users, err
	}

	marshaled, err := json.Marshal(Users) //sending data to cache
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return Users, err
	}
	redisClient.Set(context.Background(), link, marshaled, time.Minute*5)

	return Users, nil
}

func RetrieveWithId(id int) (structures.User, error) {
	var User structures.User
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	linkWithId := fmt.Sprintf("%s/%d", link, id)

	data, _ := redisClient.Get(ctx, linkWithId).Result() //retrieving data from cache
	if data != "" {                                      //take data from cache
		json.Unmarshal([]byte(data), &User)
		return User, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, linkWithId, nil)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return User, err
	}

	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("Status code %d\n", resp.StatusCode)
		fmt.Println(err)
		return User, errors.New(err)
	}

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return User, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return User, err
	}
	if err = json.Unmarshal(body, &User); err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return User, err
	}

	marshaled, err := json.Marshal(User) //sending data to cache
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		return User, err
	}
	redisClient.Set(context.Background(), linkWithId, marshaled, time.Minute*5)

	return User, nil
}

func RedisClientInit() {
	var redisClt redis.Options
	yamlFile, err := os.ReadFile("redisClient.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Parse YAML data into the Config struct

	err = yaml.Unmarshal(yamlFile, &redisClt)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	// Initialize Redis client in the init function
	redisClient = redis.NewClient(&redisClt)
}
