package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hiroshiyoka/task-5-vix-btpns-Raka/config"
	"github.com/hiroshiyoka/task-5-vix-btpns-Raka/helpers"
	"github.com/hiroshiyoka/task-5-vix-btpns-Raka/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Take input from json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// Take data based on username
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username atau Password salah"}
			helpers.ResponseJson(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helpers.ResponseJson(w, http.StatusInternalServerError, response)
			return
		}
	}

	// Password Check Validation
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username atau Password salah"}
		helpers.ResponseJson(w, http.StatusUnauthorized, response)
		return
	}

	// Making a jwt token process
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "mux-restapi",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// Declare algorithm for signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// Set token to cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login berhasil"}
	helpers.ResponseJson(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// Take input from json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// Hash password using bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// Insert to database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Success"}
	helpers.ResponseJson(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Delete token in cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Logout berhasil"}
	helpers.ResponseJson(w, http.StatusOK, response)
}
