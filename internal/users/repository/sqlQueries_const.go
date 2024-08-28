package repository

import "errors"

// Errors
var (
	// ErrSignUpRepo is
	ErrSignUpRepo = errors.New("missing params in repo")
)

// SQL Queries are
const (
	signUPQuery = `INSERT INTO users (
						username, email, password_hash)
						 	VALUES ($1, $2, $3)
								RETURNING id;`
)
