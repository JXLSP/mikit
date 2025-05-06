package report

type Reporter interface {
	Connect(target, username, password string) (bool, error)
	Unauthorized(target string) (bool, error)
}
