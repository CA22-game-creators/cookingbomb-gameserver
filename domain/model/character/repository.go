//go:generate mockgen -source=$GOFILE -destination=../../../mock/domain/model/character/$GOFILE
package domain

type Repository interface {
	Update(Character)
	GetAll() []Character
	Delete(Character)
}
