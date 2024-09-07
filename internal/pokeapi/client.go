package pokeapi

import (
	"net/http"
	"time"

	"github.com/Cap3r-0/pokedex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
