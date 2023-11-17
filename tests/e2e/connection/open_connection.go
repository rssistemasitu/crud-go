package connection

import (
	"context"
	"log"
	"time"

	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Teste e2e utilizando Docker
func OpenConnection() (database *mongo.Database, close func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
		return
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
	})
	if err != nil {
		log.Fatalf("Could not create mongo container: %s", err)
		return
	}

	mongoURI := "mongodb://localhost:" + resource.GetPort("27017/tcp")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println("Error trying to open connection")
	}

	database = client.Database("crud-go")

	close = func() {
		err := resource.Close()
		if err != nil {
			log.Println("Error trying to open connection")
			return
		}

		//Adicionar um pequeno atraso antes de encerrar o contêiner
		time.Sleep(1 * time.Second)

		//Encerrar o contêiner do MongoDB
		if err := pool.Purge(resource); err != nil {
			log.Println("Error purging resource:", err)
			return
		}
	}

	return
}
