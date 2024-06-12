package models

import "time"

type Client struct {
	Name     string
	Table    int
	JoinTime time.Time
}
