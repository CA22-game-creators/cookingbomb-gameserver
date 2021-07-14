//go:generate mockgen -source=$GOFILE -destination=../../../mock/domain/model/character/$GOFILE
package domain

type Repository interface {
	Add(Character) int
	Update(Character, int)
	GetAll() *[]Character
}
