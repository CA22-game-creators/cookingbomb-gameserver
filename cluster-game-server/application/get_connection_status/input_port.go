//go:generate mockgen -source=$GOFILE -destination=../../mock/application/get_connection_status/$GOFILE
package application

type InputPort interface {
	Handle(InputData) OutputData
}
