package api

//TODO
type AuthRequestParam struct {
	sessionToken string
}

type AuthResponceSet struct {
	result bool
}

func AuthRequest(AuthRequestParam) (AuthResponceSet, error) {
	//TODO
	result := AuthResponceSet{result: true}
	return result, nil
}
