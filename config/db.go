package config

import (
    "context"
    "os"

    "github.com/luqmanshaban/go-eccomerce/initializers"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectToDB() {
    initializers.LoadEnv()
    dbString := os.Getenv("DB_STRING")

    // Use Connect instead of deprecated NewClient
    ctx := context.Background()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbString))
    if err != nil {
        panic(err)
    }

    println("DB CONNECTED")
    DB = client.Database("Cluster0")

    // Defer client.Disconnect(context.Background()) is no longer necessary
    // The Connect method automatically manages connection lifetime
}
