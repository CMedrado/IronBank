package login

type Repository interface {
	AuthenticatedLogin(string, string) (error, string)
}
