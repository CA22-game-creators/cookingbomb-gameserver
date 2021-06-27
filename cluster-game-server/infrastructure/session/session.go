package session

import (
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/model/session"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

var cacheInstance = New()

func CheckSessionActive(token string) bool {
	v, found := cacheInstance.GetValue(token)
	if found {
		return v.IsActive()
	}
	return false
}

func GetSessionStatus(token string) pb.ConnectionStatusEnum {
	v, found := cacheInstance.GetValue(token)
	if found {
		return v.GetStatus()
	}
	return pb.ConnectionStatusEnum_CONNECTION_UNSPECIFIED
}

func ActivateSession(token string) error {
	if CheckSessionActive(token) {
		return errors.InvalidOperation()
	}
	s := session.Session{Status: pb.ConnectionStatusEnum_CONNECTED}
	cacheInstance.SetValue(token, s)
	return nil
}

func EndSessionByServer(token string) error {
	if !CheckSessionActive(token) {
		return errors.InvalidOperation()
	}
	s := session.Session{Status: pb.ConnectionStatusEnum_DISCONNECTED_BY_SERVER}
	cacheInstance.SetValueWithExpiration(token, s)
	return nil
}

func EndSessionByClient(token string) error {
	if !CheckSessionActive(token) {
		return errors.InvalidOperation()
	}
	s := session.Session{Status: pb.ConnectionStatusEnum_DISCONNECTED_BY_CLIENT}
	cacheInstance.SetValueWithExpiration(token, s)
	return nil
}

// 基本使用しない。エラー処理等で使う
func ForceEndSession(token string) {
	s := session.Session{Status: pb.ConnectionStatusEnum_DISCONNECTED}
	cacheInstance.SetValueWithExpiration(token, s)
}

func ClearSession() {
	cacheInstance.Flush()
}
