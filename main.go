package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017" + dbName

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    int64
}

func Connect() error {

}

func main() {

}
