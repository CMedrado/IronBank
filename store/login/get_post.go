package login

func (a *StoredToken) PostToken(id int, token string) {
	accountToken[id] = Token{Token: token}
}

func (a *StoredToken) GetTokenID(id int) Token {
	return accountToken[id]
}
