//go:generate mockgen -source=$GOFILE -destination=../../mock/application/connect/$GOFILE
package application

type InputPort interface {
	Handle(InputData) OutputData
}
