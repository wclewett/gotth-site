package store

type User struct {
  Id  uint  `json:"id"`
  Email  string  `json:"email"`
  Password  string `json:"-"`
}

type Session struct {
  Id        uint  `json:"id"` 
	SessionID string `json:"session_id"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID" json:"user"`
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}

type SessionStore interface {
	CreateSession(session *Session) (*Session, error)
	GetUserFromSession(sessionID string, userID string) (*User, error)
}
