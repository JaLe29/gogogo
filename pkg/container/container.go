package container

import (
	db "bastard-proxy/db"
	"context"
	"log"
	"os"
)

type Domain struct {
	Id          string
	TargetProxy string
}

type AppContainer struct {
	Context          context.Context
	PrismaClient     *db.PrismaClient
	DomainMap        *map[string]Domain
	RefetchDomainMap func()
}

func InitContainer() AppContainer {
	client := db.NewClient() // Initialize the client using the imported package
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	context := context.Background()
	dm := &map[string]Domain{}

	refetchDomainMap := func() {
		clear(*dm)

		log.Println("Refetching domain map")
		res, _ := client.Proxy.FindMany().Exec(context)
		for _, proxy := range res {
			(*dm)[proxy.Source] = Domain{
				Id:          proxy.ID,
				TargetProxy: proxy.Target,
			}

			log.Println("Proxying " + proxy.Source + " to " + proxy.Target)
		}

		adminDomain := os.Getenv("ADMIN_DOMAIN")
		adminTarget := "localhost:5000"
		(*dm)[adminDomain] = Domain{
			Id:          "",
			TargetProxy: adminTarget,
		}
		log.Println("Proxying " + adminDomain + " to " + adminTarget)
	}

	return AppContainer{
		Context:          context,
		PrismaClient:     client,
		DomainMap:        dm,
		RefetchDomainMap: refetchDomainMap,
	}
}
