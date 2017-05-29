package authentication

import (
	"errors"
	"net/http"
	"strings"

	"gopkg.in/square/go-jose.v2"

	"gopkg.in/gin-gonic/gin.v1"

	"gopkg.in/square/go-jose.v2/jwt"
)

var (
	password  = "password"
	login     = "login"
	jwtSecret = []byte("JWTSecret")
)

func extractToken(r *http.Request) (*jwt.JSONWebToken, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if len(authorizationHeader) > 7 && strings.EqualFold(authorizationHeader[0:7], "BEARER ") {
		return jwt.ParseSigned(authorizationHeader[7:])
	}
	return nil, errors.New("Token not found")
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	// Récupération de l'identifiant et du mot de passe
	loginRequest := LoginRequest{}
	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if loginRequest.Login != login || loginRequest.Password != password {
		c.AbortWithError(http.StatusUnauthorized, errors.New("Invalid credentials"))
	}

	// Création du JWT à l'aide du mot de passe partagé
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: jwtSecret}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	cl := jwt.Claims{
		Issuer:   "glmf-go-api",
		Audience: jwt.Audience{"glmf-go-api-client"},
	}

	raw, err := jwt.Signed(signer).Claims(cl).CompactSerialize()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusOK, raw)

}

func JWTMiddleware(c *gin.Context) {

	// Extraction du JWT
	token, err := extractToken(c.Request)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// Validation du JWT
	cl := jwt.Claims{
		Issuer:   "glmf-go-api",
		Audience: jwt.Audience{"glmf-go-api-client"},
	}
	if err := token.Claims(jwtSecret, &cl); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.Next()
}
