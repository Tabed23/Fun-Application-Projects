package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Tabed23/Fun-Application-Projects/tree/main/JobsAPI-GraphQL/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connStr string = "mongodb://mongo:27017"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	clientOptions := options.Client()
	clientOptions.ApplyURI(connStr)
	cli, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err = cli.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	cli.Database("jobs-graphQL").CreateCollection(ctx, "jobs")

	return &DB{
		client: cli,
	}

}

func (d *DB) GetJobByID(c context.Context, id string) (*model.JobList, error) {
	col := d.client.Database("jobs-graphQL").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var job *model.JobList

	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": _id}

	if err := col.FindOne(ctx, filter).Decode(job); err != nil {
		log.Fatal(err)
	}

	return job, nil

}

func (d *DB) GetJobs(c context.Context) ([]*model.JobList, error) {
	col := d.client.Database("jobs-graphQL").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var jobs []*model.JobList

	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &jobs); err != nil {
		log.Fatal(err)
	}
	log.Println(jobs)

	return jobs, nil

}

func (d *DB) CreateJob(c context.Context, data model.CreateJobInput) (*model.CreateJobResponse, error) {
	col := d.client.Database("jobs-graphQL").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	instReg, err := col.InsertOne(ctx, bson.M{"title": data.Title, "description": data.Description, "company": data.Company, "url": data.URL})
	if err != nil {
		log.Fatal(err)
	}

	ID := instReg.InsertedID.(primitive.ObjectID).Hex()

	return &model.CreateJobResponse{
		Status: 200,
		ID:     ID,
		Msg:    "Job Created",
	}, nil
}

func (d *DB) UpdateJob(c context.Context, id string, update model.UpdateJobInput) (*model.UpdateJobResponse, error) {
	col := d.client.Database("jobs-graphQL").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}

	if update.Title != nil {
		updateJobInfo["title"] = update.Title
	}
	if update.Description != nil {
		updateJobInfo["description"] = update.Description
	}
	if update.URL != nil {
		updateJobInfo["url"] = update.URL
	}

	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": _id}
	updatedData := bson.M{"$set": updateJobInfo}

	res := col.FindOneAndUpdate(ctx, filter, updatedData, options.FindOneAndUpdate().SetReturnDocument(1))
	var updateJob *model.JobList

	if err := res.Decode(&updateJob); err != nil {
		log.Fatal(err)
	}

	return &model.UpdateJobResponse{
		Status:     200,
		UpdatedJob: updateJob,
		Msg:        "Job Updated",
	}, nil
}

func (d *DB) DeleteJob(c context.Context, id string) (*model.DeleteJobResponse, error) {
	col := d.client.Database("jobs-graphQL").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return &model.DeleteJobResponse{
		Status:      200,
		DeleteJobID: id,
		Msg:         "Job Deleted",
	}, nil

}
