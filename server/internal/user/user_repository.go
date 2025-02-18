package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

//repository struct定義為DB

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

// #POST3
func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertID int

	query := `
        INSERT INTO users (username, password, email)
        OUTPUT INSERTED.id
        VALUES (@username, @password, @email)
    `
	// 使用命名參數傳遞值
	err := r.db.QueryRowContext(
		ctx,
		query,
		sql.Named("username", user.Username),
		sql.Named("password", user.Password),
		sql.Named("email", user.Email),
	).Scan(&lastInsertID)

	if err != nil {
		return &User{}, err
	}
	user.ID = int64(lastInsertID)

	return user, nil
}

// 先寫跟database的互動=>在來是定義service與user=>handler最後寫
func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {

	u := User{}
	query := `SELECT id, email,username,password 
			  FROM users WHERE email = @email`
	err := r.db.QueryRowContext(ctx, query, sql.Named("email", email)).Scan(&u.ID, &u.Email, &u.Username, &u.Password)

	if err != nil {
		return &User{}, err
	}
	return &u, nil
}
