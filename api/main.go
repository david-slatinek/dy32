package main

import (
	"log"
	"main/producer"
)

func main() {
	//err := env.Load("env/.env")
	//if err != nil {
	//	log.Fatalf("error with env.Load: %v", err)
	//}
	//
	//uri := os.Getenv("MONGODB_URI")
	//if uri == "" {
	//	log.Fatal("MONGODB_URI is not set")
	//}
	//
	//serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	//clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//
	//client, err := mongo.Connect(ctx, clientOptions)
	//if err != nil {
	//	log.Fatalf("error with Connect: %v", err)
	//}
	//
	//if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
	//	log.Fatalf("error with Ping: %v", err)
	//}

	pro := producer.KafkaProducer{
		Address: "localhost:9092",
		Topic:   "kafka-invoices",
	}

	err := pro.Init()
	if err != nil {
		log.Printf("failed to dial leader: %v", err)
	}
	defer func(p producer.KafkaProducer) {
		if err := p.Close(); err != nil {
			log.Printf("failed to close writer: %v", err)
		}
	}(pro)

}
