package infrastructure

func CheckAuthToken(token string) (bool, error) {
	cache, err := CheckAuthTokenCached(token)
	if err != nil {
		return false, err
	}
	if cache {
		return true, nil
	}

	//TODO
	return true, nil
}

func CheckAuthTokenCached(token string) (bool, error) {
	//TODO
	return false, nil
}
