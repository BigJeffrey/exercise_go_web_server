package services

import (
	"fmt"
	"go.uber.org/zap"
	"webserver/consts"
	"webserver/interfaces"
	"webserver/models"
)

type DeleteProductService struct {
	Log *zap.SugaredLogger
	PG  interfaces.PostgresqlInterface
}

func (d *DeleteProductService) DeleteProductService(idRequest models.IdRequest) error {
	productRetrieved, err := d.PG.GetProductById(idRequest.ID)
	if err != nil {
		d.Log.Error(err.Error())
		return err
	}

	if productRetrieved == nil {
		d.Log.Infof("nothing to delete: %v", idRequest.ID)
		return fmt.Errorf(consts.ProductNotFoundError)
	}

	_, err = d.PG.DeleteProduct(idRequest.ID)
	if err != nil {
		d.Log.Error(err.Error())
		return err
	}

	return nil
}
