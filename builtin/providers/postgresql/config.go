package postgresql

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //PostgreSQL db
)

// Config - provider config
type Config struct {
	Host              string
	Port              int
	Database          string
	Username          string
	Password          string
	SSLMode           string
	ApplicationName   string
	Timeout           int
	ConnectTimeoutSec int
}

// Client struct holding connection string
type Client struct {
	username string
	connStr  string
}

// NewClient returns new client config
func (c *Config) NewClient() (*Client, error) {
	// NOTE: dbname must come before user otherwise dbname will be set to
	// user.
	const dsnFmt = "host=%s port=%d dbname=%s user=%s password=%s sslmode=%s fallback_application_name=%s connect_timeout=%d"

	// Quote empty strings or strings that contain whitespace
	q := func(s string) string {
		b := bytes.NewBufferString(`'`)
		b.Grow(len(s) + 2)
		for _, r := range s {
			switch r {
			case '\'':
				b.WriteString(`\'`)
			case '\\':
				b.WriteString(`\\`)
			default:
				b.WriteRune(r)
			}
		}

		b.WriteString(`'`)
		return b.String()
	}

	logDSN := fmt.Sprintf(dsnFmt, q(c.Host), c.Port, q(c.Database), q(c.Username), q("<redacted>"), q(c.SSLMode), q(c.ApplicationName), c.ConnectTimeoutSec)
	log.Printf("[INFO] PostgreSQL DSN: `%s`", logDSN)

	connStr := fmt.Sprintf(dsnFmt, q(c.Host), c.Port, q(c.Database), q(c.Username), q(c.Password), q(c.SSLMode), q(c.ApplicationName), c.ConnectTimeoutSec)
	client := Client{
		connStr:  connStr,
		username: c.Username,
	}

	return &client, nil
}

// Connect will manually connect/disconnect to prevent a large
// number or db connections being made
func (c *Client) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", c.connStr)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to PostgreSQL server: %s", err)
	}

	return db, nil
}
