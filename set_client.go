package rejson

import (
	"github.com/go-redis/redis/v8"
	"github.com/kaiaulu/go-rejson/clients"
	"github.com/kaiaulu/go-rejson/rjs"
)

// RedisClient provides interface for Client handling in the ReJSON Handler
type RedisClient interface {
	SetClientInactive()
	SetRedisClient(conn *redis.Client)
}

// SetClientInactive resets the handler and unset any client, set to the handler
func (r *Handler) SetClientInactive() {
	_t := &Handler{clientName: rjs.ClientInactive}
	r.clientName = _t.clientName
	r.implementation = _t.implementation
}

// SetRedisClient sets Go-Redis (https://github.com/go-redis/redis) client to
// the handler
func (r *Handler) SetRedisClient(conn *redis.Client) {
	r.clientName = "redis"
	r.implementation = &clients.GoRedis{Conn: conn}
}
