//go:generate mockgen -source=$GOFILE -destination=../../mock/application/disconnect/$GOFILE
package application

type InputPort interface {
	Handle(InputData) OutputData
}
