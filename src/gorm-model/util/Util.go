package util

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	dto "gorm-model/model/dto"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/golang-jwt/jwt"
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
	tokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.AccessUUID = uuid.NewV4().String()
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

// CreateAuth function
func CreateAuth(context *gin.Context, userid uint64, td *dto.TokenDetails) error {
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

// ExtractToken function
func ExtractToken(request *http.Request) string {
	bearToken := request.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

// VerifyToken function
func VerifyToken(request *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(request)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// TokenAuthMiddleware function
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := VerifyToken(context.Request)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  err.Error()})
			context.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {
			accessUUID, ok := claims["access_uuid"].(string)

			if !ok {
				context.JSON(http.StatusUnauthorized, gin.H{
					"status": http.StatusUnauthorized,
					"error":  err.Error()})
				context.Abort()
				return
			}

			fetchData, getErr := FetchAuth(context, accessUUID)

			if getErr != nil || fetchData == 0 { //if any goes wrong
				context.JSON(http.StatusUnauthorized, gin.H{
					"status": http.StatusUnauthorized,
					"error":  "Unauthorized"})
				context.Abort()
				return
			}
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  err.Error()})
			context.Abort()
			return
		}

		context.Next()
	}
}

// ExtractTokenMetadata function
func ExtractTokenMetadata(request *http.Request) (*dto.AccessDetails, error) {
	token, err := VerifyToken(request)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)

		if !ok {
			return nil, err
		}

		userID, ok := claims["user_id"].(string)
		// userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)

		if !ok {
			return nil, err
		}

		return &dto.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}

	return nil, err
}

// FetchAuth function
func FetchAuth(context *gin.Context, givenUUID string) (uint64, error) {
	redisClient := context.MustGet("redis").(*redis.Client)
	fetched, err := redisClient.Get(givenUUID).Result()

	if err != nil {
		return 0, err
	}

	fetchData, _ := strconv.ParseUint(fetched, 10, 64)
	return fetchData, nil
}

// DeleteAuth function
func DeleteAuth(context *gin.Context, givenUUID string) (int64, error) {
	redisClient := context.MustGet("redis").(*redis.Client)
	deleted, err := redisClient.Del(givenUUID).Result()

	if err != nil {
		return 0, err
	}

	return deleted, nil
}
