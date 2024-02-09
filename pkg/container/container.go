package container

import (
	db "bastard-proxy/db"
	"bastard-proxy/pkg/logger"
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Domain struct {
	Id          string
	TargetProxy string
	Disable     bool
}

type AppContainer struct {
	Context      context.Context
	PrismaClient *db.PrismaClient

	DomainMap *map[string]Domain
	BlockMap  *map[string]map[string]bool
	AllowMap  *map[string]map[string]bool

	RefetchDomainMap func()
	RefetchBlockMap  func()

	Logger logger.Logger
}

func InitContainer() AppContainer {
	client := db.NewClient() // Initialize the client using the imported package
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	context := context.Background()
	dm := &map[string]Domain{}
	bm := &map[string]map[string]bool{}
	am := &map[string]map[string]bool{}

	refetchDomainMap := func() {
		clear(*dm)

		log.Println("Refetching domain map")
		res, _ := client.Proxy.FindMany().Exec(context)
		for _, proxy := range res {
			(*dm)[proxy.Source] = Domain{
				Id:          proxy.ID,
				TargetProxy: proxy.Target,
				Disable:     proxy.Disable,
			}

			log.Println("Proxying " + proxy.Source + " to " + proxy.Target)
		}

		adminDomain := os.Getenv("ADMIN_DOMAIN")
		if adminDomain != "" {
			adminTarget := "localhost:5000"
			(*dm)[adminDomain] = Domain{
				Id:          "",
				TargetProxy: adminTarget,
				Disable:     false,
			}
			log.Println("Proxying " + adminDomain + " to " + adminTarget)
		}

	}

	refetchBlockMap := func() {
		clear(*bm)

		log.Println("Refetching block map")
		res, _ := client.Block.FindMany().Exec(context)
		for _, block := range res {
			proxyId, _ := block.ProxyID()

			if _, ok := (*bm)[proxyId]; !ok {
				(*bm)[proxyId] = make(map[string]bool)
			}

			(*bm)[proxyId][block.IP] = true
			log.Println("Blocking " + proxyId + ", " + block.IP)
		}
	}

	refetchAllowMap := func() {
		clear(*am)
		log.Println("Refetching allow map")
		res, _ := client.Allow.FindMany().Exec(context)
		for _, allow := range res {
			proxyId, _ := allow.ProxyID()

			if _, ok := (*am)[proxyId]; !ok {
				(*am)[proxyId] = make(map[string]bool)
			}

			(*am)[proxyId][allow.IP] = true
			log.Println("Allowing " + proxyId + ", " + allow.IP)
		}
	}

	fireInitialRefetch := func() {
		ticker := time.NewTicker(60 * time.Second)
		quit := make(chan struct{})
		go func() {
			for {
				select {
				case <-ticker.C:
					refetchDomainMap()
					refetchBlockMap()
					refetchAllowMap()
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()
	}

	refetchDomainMap()
	refetchBlockMap()
	refetchAllowMap()

	fireInitialRefetch()

	mongoUri := os.Getenv("DATABASE_URL")
	mongoClient, err := mongo.Connect(context, options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	startWatchBlock := func() {
		changeStream, err := mongoClient.Database("bastard-proxy").Collection("Block").Watch(context, mongo.Pipeline{}, options.ChangeStream())
		if err != nil {
			panic(err)
		}

		for changeStream.Next(context) {
			// fmt.Println(changeStream.Current)
			refetchBlockMap()
		}
	}

	startWatchProxy := func() {
		changeStream, err := mongoClient.Database("bastard-proxy").Collection("Proxy").Watch(context, mongo.Pipeline{}, options.ChangeStream())
		if err != nil {
			panic(err)
		}

		for changeStream.Next(context) {
			// fmt.Println(changeStream.Current)
			refetchDomainMap()
		}
	}

	startWatchAllow := func() {
		changeStream, err := mongoClient.Database("bastard-proxy").Collection("Allow").Watch(context, mongo.Pipeline{}, options.ChangeStream())
		if err != nil {
			panic(err)
		}

		for changeStream.Next(context) {
			// fmt.Println(changeStream.Current)
			refetchAllowMap()
		}
	}

	go startWatchBlock()
	go startWatchProxy()
	go startWatchAllow()

	logger := logger.New()

	return AppContainer{
		Context:          context,
		PrismaClient:     client,
		DomainMap:        dm,
		BlockMap:         bm,
		AllowMap:         am,
		RefetchDomainMap: refetchDomainMap,
		RefetchBlockMap:  refetchBlockMap,
		Logger:           logger,
	}
}
