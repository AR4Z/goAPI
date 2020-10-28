package manager

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

func DgraphClient() *dgo.Dgraph {
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func Setup(client *dgo.Dgraph) {
	err := client.Alter(context.Background(), &api.Operation{
		Schema: `
			id: string @index(exact) .
			name: string @index(exact) .
			age: int .
			price: int .
			buyer_id: uid @reverse .
			ip: string @index(exact) .
			device: string @index(exact) .
			products: [uid] @count @reverse .
			
			type Buyer {
				id
				name
				age
			}
			type Product {
				id
				name
				price
			}
			type Transaction {
				id
				buyer_id
				ip
				device
				products
			}
		`,
	})

	if err != nil {
		log.Fatal(err)
	}
}