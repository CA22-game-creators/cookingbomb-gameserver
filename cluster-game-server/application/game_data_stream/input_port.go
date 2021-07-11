//go:generate mockgen -source=$GOFILE -destination=../../mock/application/game_data_stream/$GOFILE
package application

type InputPort interface {
	Handle(InputData)
}
