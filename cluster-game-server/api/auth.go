package api

//TODO
type AuthRequestParam struct {
	authToken string
}

type AuthResponceSet struct {
	result bool
}

func AuthRequest (AuthRequestParam) (AuthResponceSet, error) {
	//TODO
	result AuthResponceSet := {Result := true};
	return 
}