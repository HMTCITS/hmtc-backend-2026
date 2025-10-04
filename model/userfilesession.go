package model

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type UserFileSession struct {
	SessionId uuid.UUID `json:"session_id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	NRP       string    `json:"nrp" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	AlasanReq string    `json:"alasan_req" binding:"required"`
	Status    Status    `json:"status" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewUserFileSession(name string, nrp string, email string, alasan string) (UserFileSession, error) {
	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`

	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return UserFileSession{}, errors.New("invalid email")
	}

	now := time.Now()

	return UserFileSession{
		SessionId: uuid.New(),
		Name:      name,
		NRP:       nrp,
		Email:     email,
		AlasanReq: alasan,
		Status:    StatusPending,
		CreatedAt: now,
		ExpiredAt: now.Add(24 * time.Hour),
	}, nil
}

func (u *UserFileSession) IsExpired() bool {
	return time.Now().After(u.ExpiredAt)
}
