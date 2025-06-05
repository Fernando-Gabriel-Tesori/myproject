package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func ConnectMongo(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	mongoClient = client
	log.Println("✅ Conectado ao MongoDB com sucesso!")
}

func GetCollection(dbName string, collection string) *mongo.Collection {
	if mongoClient == nil {
		log.Fatal("❌ MongoDB ainda não foi inicializado. Chame ConnectMongo() primeiro.")
	}
	return mongoClient.Database(dbName).Collection(collection)
}
