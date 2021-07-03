package domain

import (
	"regexp"

	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
)

type Name string

func NewName(v string) (Name, error) {
	isValid, err := regexp.MatchString("^([0-9０-９a-zA-Zぁ-んァ-ヶｦ-ﾟ一-龠]{1,10})$", v)
	if err != nil {
		return "", errors.InvalidOperation()
	}
	if !isValid {
		return "", errors.InvalidArgument("ユーザー名は半角英数字か日本語の1-10文字である必要があります")
	}

	return Name(v), nil
}
