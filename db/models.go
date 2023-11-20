package db

import (
	"crypto/sha256"
	"encoding/base64"
	uuid "github.com/satori/go.uuid"
)

type BaseModel struct {
	Uid       uuid.UUID `gorm:"type=uuid;primary_key;"`
	CreatedAt int       `gorm:"autoCreateTime:milli"`
	UpdatedAt int       `gorm:"autoUpdateTime:milli"`
}
type User struct {
	BaseModel
	Username string `gorm:"unique"`
	Phone    string `gorm:"unique"`
	Password string
}

func (user *User) SetPassword(rawPassword string) {
	hash := sha256.Sum256([]byte(rawPassword))
	user.Password = base64.StdEncoding.EncodeToString(hash[:])
}

func (user *User) CheckPassword(rawPassword string) bool {
	hash := sha256.Sum256([]byte(rawPassword))
	return user.Password == base64.StdEncoding.EncodeToString(hash[:])
}

type Chat struct {
	BaseModel
	ChatterOneID int
	ChatterOne   User `gorm:"foreignKey:ChatterOneID"`
	ChatterTwoID int
	ChatterTwo   User `gorm:"foreignKey:ChatterTwoID"`
}

func (chat *Chat) clear() {
	// TODO: impl
}

func (chat *Chat) delete() {

}

type Message struct {
	BaseModel
	Body     string
	ChatID   int
	Chat     Chat
	SenderID int
	Sender   User `gorm:"foreignKey:SenderID"`
}

//TODO: contacts
//TODO: add contacts
