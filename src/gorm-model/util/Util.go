package util

import (
	"github.com/gin-gonic/gin"

	dto "gorm-model/model/dto"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// ReformatDate function
func ReformatDate(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

// CreateToken function
func CreateToken(email string) (*dto.TokenDetails, error) {
	tokenDetails := &dto.TokenDetails{}
	tokenDetails.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tokenDetails.AccessUUID = uuid.NewV4().String()
	tokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshUUID = uuid.NewV4().String()

	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = tokenDetails.AccessUUID
	atClaims["user_id"] = email
	atClaims["exp"] = tokenDetails.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenDetails.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tokenDetails.RefreshUUID
	rtClaims["user_id"] = email
	rtClaims["exp"] = tokenDetails.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	tokenDetails.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))

	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}

// StoreToken function
func StoreToken(context *gin.Context, userid uint64, td *dto.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	redisClient := context.MustGet("redis").(*redis.Client)

	errAccess := redisClient.Set(td.AccessUUID, strconv.Itoa(int(userid)), at.Sub(now)).Err()

	if errAccess != nil {
		return errAccess
	}

	errRefresh := redisClient.Set(td.RefreshUUID, strconv.Itoa(int(userid)), rt.Sub(now)).Err()

	if errRefresh != nil {
		return errRefresh
	}

	return nil
}
