package repository

import (
	"agent/ai-agent-platform/internal/pkg/database"
	"agent/ai-agent-platform/internal/server/model"
	"context"
	"database/sql"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}

type userRepository struct {
	db *database.Postgres
}

func NewUserRepository(db *database.Postgres) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	conn, err := r.db.GetReadConnection(ctx)
	if err != nil {
		return nil, err
	}

	query:= `SELECT 
		id, 
		username, 
		email, 
		phone, 
		first_name, 
		last_name, 
		password, 
		status, 
		last_login_at, 
		created_at, 
		updated_at 
	FROM users WHERE id = $1`

	var user model.User
	err = conn.QueryRowxContext(ctx, query, id).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, err
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	conn, err := r.db.GetReadConnection(ctx)
	if err != nil {
		return nil, err
	}

	query:= `SELECT 
		id,
		username,
		email,
		phone,
		first_name,
		last_name,
		password,
		status,
		last_login_at,
		created_at,
		updated_at
		FROM users WHERE email = $1`

	var user model.User
	err = conn.QueryRowxContext(ctx, query, email).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, err
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	conn, err := r.db.GetWriteConnection(ctx)
	if err != nil {
		return nil, err
	}

	query:= `
	INSERT INTO users (
		id,
		username,
		email, 
		phone,
		first_name,
		last_name,
		password,
		status,
		last_login_at,
		created_at,
		updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		RETURNING *`

	err = conn.QueryRowxContext(
		ctx,
		query,
		user.ID,
		user.Username,
		user.Email,
		user.Phone,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Status,
		user.LastLoginAt,
		user.CreatedAt,
		user.UpdatedAt).StructScan(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}