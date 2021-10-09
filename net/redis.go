package net

import (

  // Import builtin packages.
  "sync"

  // Import 3rd party packages.
  "github.com/go-redis/redis"
  "github.com/rs/zerolog/log"
)

// RedisConfig defines the configuration of a Redis client.
type RedisConfig struct {
  ClientName string   `json:"client_name"`
  Address    string   `json:"address"`
  Password   string   `json:"password"`
  DB         int      `json:"db"`
}

var redisSessions map[string]*redis.Client = map[string]*redis.Client{}
var redisSessionsMu sync.Mutex

func RedisConnect(config RedisConfig) *redis.Client {

  // Declare the options and connect.
  options := &redis.Options{
    Addr:     config.Address,
    Password: config.Password,
    DB:       config.DB,
  }
  client := redis.NewClient(options)

  log.Info().Msgf("Connected to redis. Client name: %q", config.ClientName)

  // Put the driver in the public map.
  redisSessionsMu.Lock()
  redisSessions[config.ClientName] = client
  redisSessionsMu.Unlock()

  // Return the driver object.
  return client
}

// RedisGetClient gets the Redis client keyed by the specified name. Panics if the client hasn't
// been created as of this call.
func RedisGetClient(name string) *redis.Client {

  redisSessionsMu.Lock()
  client, ok := redisSessions[name]
  redisSessionsMu.Unlock()
  if !ok {
    panic("unitialized redis client: " + name)
  }

  return client
}