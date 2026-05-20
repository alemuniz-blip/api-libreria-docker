package security

import "time"

type Builder interface {
	CreateToken(role string, name string, image string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}