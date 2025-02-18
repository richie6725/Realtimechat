package user

import (
	"Go_realtime_chat/server/internal/util"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

const (
	secretKey = "secret"
)

//Service struct定義為Repository

type service struct {
	Repository
	timeout time.Duration
}

// 初始化#3
// 建立一個回傳repository的service
func NewService(repository Repository) Service { //Repository&Service=interface
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

// #POST2

// service接收到資料後，將密碼hash，且傳送相關資料給repository業務
func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserResp, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	r, err := s.Repository.CreateUser(ctx, u) //接Repository的CreateUser
	if err != nil {
		return nil, err
	}
	res := &CreateUserResp{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil

}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserResp, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserResp{}, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserResp{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserResp{}, err
	}
	return &LoginUserResp{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}
