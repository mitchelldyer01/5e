package models

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Token struct {
	PlayerID int
	jwt.StandardClaims
}

type Player struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
}

func (p *Player) Insert(db *gorm.DB) error {
	err := p.hashPassword()
	if err != nil {
		return err
	}

	result := db.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Player) Update(db *gorm.DB) error {
	err := p.hashPassword()
	if err != nil {
		return err
	}

	result := db.Save(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Player) UpdateJWT(db *gorm.DB) error {
	p.generateToken()

	result := db.Model(&p).Update("token", &p.Token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Player) SelectLogin(db *gorm.DB) error {
	err := p.hashPassword()
	if err != nil {
		return err
	}

	result := db.Model(&p).First(&p).Scan(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Player) hashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Password = string(hashedPassword)
	return nil
}

func (p *Player) generateToken() error {
	p.setJWTExpiration()

	t := &Token{
		PlayerID: p.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: p.Expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)

	signed, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return err
	}

	p.Token = signed
	return nil
}

func (p *Player) setJWTExpiration() {
	p.Expiration = time.Now().Add(24 * time.Hour)
}
