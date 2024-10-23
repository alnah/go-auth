// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
)

type UserCore struct {
	ID    int32  `json:"id"`
	Email string `json:"email"`
	// bcrypt
	Hash      string       `json:"hash"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type UserCoreRole struct {
	CoreID int32 `json:"core_id"`
	RoleID int32 `json:"role_id"`
}

type UserRole struct {
	ID        int32        `json:"id"`
	RoleName  string       `json:"role_name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type UserToken struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
	// bcrypt
	Hash      string       `json:"hash"`
	UserAgent string       `json:"user_agent"`
	IpAddress string       `json:"ip_address"`
	CreatedAt sql.NullTime `json:"created_at"`
	ExpiresAt sql.NullTime `json:"expires_at"`
	RevokedAt sql.NullTime `json:"revoked_at"`
}