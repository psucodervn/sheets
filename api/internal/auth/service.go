package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"

	"api/internal/api"
	"api/model"
)

type Service struct {
	db         *sql.DB
	jwtSecret  []byte
	expireTime time.Duration
	googleConf oauth2.Config
}

func NewService(db *sql.DB, jwtSecret []byte, googleConf oauth2.Config) *Service {
	return &Service{
		db:         db,
		jwtSecret:  jwtSecret,
		expireTime: 7 * 24 * time.Hour,
		googleConf: googleConf,
	}
}

const endpointProfile string = "https://www.googleapis.com/oauth2/v2/userinfo"

func (s *Service) FetchGoogleUserWithCode(ctx context.Context, code string) (*GoogleUser, error) {
	tok, err := s.googleConf.Exchange(ctx, code, oauth2.AccessTypeOffline)
	if err != nil {
		return nil, err
	}

	resp, err := oauth2.NewClient(ctx, oauth2.StaticTokenSource(tok)).Get(endpointProfile)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("google responded with a %d trying to fetch user information", resp.StatusCode)
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var gu GoogleUser
	if err := json.Unmarshal(respBytes, &gu); err != nil {
		return nil, err
	}
	return &gu, nil
}

func (s *Service) SignWithUser(u *model.User) (string, error) {
	claims := &api.UserClaims{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email.String,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.expireTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}
