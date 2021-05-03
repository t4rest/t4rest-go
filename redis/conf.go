package redis

import (
	"time"

	"github.com/pkg/errors"
)

// Conf config struct
type Conf struct {
	Database          int
	RedisPoolSize     int
	MasterName        string
	Address           string
	Password          string
	SentinelAddresses []string
	DialTimeout       time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	PoolTimeout       time.Duration
}

// Validate validates kafka config
func (c Conf) Validate() error {

	if c.Database == 0 {
		return errors.New("no redis Database")
	}

	if c.MasterName == "" {
		return errors.New("no redis MasterName")
	}

	if len(c.SentinelAddresses) == 0 {
		return errors.New("no redis SentinelAddresses")
	}

	if c.Password == "" {
		return errors.New("no redis Password")
	}

	if c.RedisPoolSize == 0 {
		return errors.New("no redis RedisPoolSize")
	}

	if c.DialTimeout == 0 {
		return errors.New("no redis DialTimeout")
	}

	if c.ReadTimeout == 0 {
		return errors.New("no redis ReadTimeout")
	}

	if c.WriteTimeout == 0 {
		return errors.New("no redis WriteTimeout")
	}

	if c.PoolTimeout == 0 {
		return errors.New("no redis PoolTimeout")
	}

	return nil
}
