package main

import (
	"context"
	"fmt"
	proto "tomato/proto"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TomatoService struct{}

type Feedback struct {
	ID       primitive.ObjectID `bson:"_id"`
	FullName string             `bson:"full_name"`
	Email    string             `bson:"email"`
	Content  string             `bson:"content"`
	Phone    string             `bson:"phone"`
}

// AddFeedback used to get non-loggedin users' feedback
func (g *TomatoService) AddFeedback(ctx context.Context, req *proto.AddFeedbackRequest, rsp *proto.AddFeedbackResponse) error {
	feedback := Feedback{
		ID:       primitive.NewObjectID(),
		FullName: req.FullName,
		Email:    req.Email,
		Content:  req.Content,
		Phone:    req.Phone}

	fmt.Println(req)

	insertResult, err := MongoCollection.InsertOne(context.TODO(), feedback)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Feedback: %+v\n", insertResult)
	rsp.IsCreated = true
	return nil
}

func (g *TomatoService) GetFeedback(ctx context.Context, req *proto.GetFeedbackRequest, rsp *proto.GetFeedbackResponse) error {
	return nil
}
