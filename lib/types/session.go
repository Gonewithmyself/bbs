package types

type Session struct {
	Authed bool
	User   *User
	Times  int32
}

func (sess *Session) Set(user *User) {
	sess.User = user
}

func (sess *Session) Get() *User {
	return sess.User
}

func (sess *Session) DoesLogin() bool {
	return sess.Authed
}

func (sess *Session) SetAuth(status bool) {
	sess.Authed = status
}
