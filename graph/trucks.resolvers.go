package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/smtp"
	"os"

	"github.com/Henoch-kargo/kargo-trucks/graph/generated"
	"github.com/Henoch-kargo/kargo-trucks/graph/model"
)

func (r *mutationResolver) SaveTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	truck := &model.Truck{
		ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
		PlateNo: plateNo,
	}
	r.Trucks[truck.ID] = *truck
	return truck, nil
}

func (r *mutationResolver) DeleteTruck(ctx context.Context, id *string) (*model.Truck, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendTruckDatatoEmail(ctx context.Context, email string) (*model.Email, error) {
	fmt.Println(email)
	emails := &model.Email{
		Email: email,
	}

	from := "henoch0chendra@gmail.com"
	password := os.Getenv("EMAILPASSWORD")

	// Receiver email address.
	to := []string{
		email,
		"henoch0chendra@gmail.com",
	}

	// // smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// // Message.
	message := []byte("This is a test email message.")

	// // Authentication.
	auth := smtp.PlainAuth("", from, password,
		smtpHost)

	// // Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Email Sent Successfully!")
	return emails, nil
}

func (r *queryResolver) PaginatedTrucks(ctx context.Context) ([]*model.Truck, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
