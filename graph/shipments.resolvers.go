package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/Henoch-kargo/kargo-trucks/graph/model"
)

func (r *mutationResolver) SaveShipment(ctx context.Context, id *string, name string, origin string, destination string, deliveryDate string, truckID string) (*model.Shipment, error) {
	if _, ok := r.Trucks[truckID]; !ok {
		return nil, errors.New("TRUCK_UNAVAILABLE")
	}

	shipment := &model.Shipment{
		ID:           fmt.Sprintf("Shipment-%d", len(r.Shipments)+1),
		Name:         name,
		Origin:       origin,
		Destination:  destination,
		DeliveryDate: deliveryDate,
		TruckID:      truckID,
	}
	r.Shipments[shipment.ID] = *shipment
	return shipment, nil
}

func (r *queryResolver) PaginatedShipments(ctx context.Context, id *string, origin *string, page int, first int) ([]*model.Shipment, error) {
	panic(fmt.Errorf("not implemented"))
}
