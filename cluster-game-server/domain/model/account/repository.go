//go:generate mockgen -source=$GOFILE -destination=../../../mock/domain/model/account/$GOFILE
package domain

type Repository interface {
	Find(sessionToken string) (Account, error)
	GetSessionStatus(sessionToken string) StatusEnum
	CheckSessionActive(sessionToken string) bool
	Connect(sessionToken string)
	Disconnect(sessionToken string)
}
