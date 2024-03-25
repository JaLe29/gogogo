package container

import (
	db "bastard-proxy/db"
	config "bastard-proxy/pkg/config"
	"bastard-proxy/pkg/logger"
	"bastard-proxy/pkg/metrics"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Domain struct {
	Id          string
	TargetProxy string
	Disable     bool
	Cache       bool
}

type AppContainer struct {
	Config       config.Config
	Context      context.Context
	PrismaClient *db.PrismaClient

	GuardMap        *map[string]bool
	GuardExcludeMap *map[string]bool
	DomainMap       *map[string]Domain
	BlockMap        *map[string]map[string]bool
	AllowMap        *map[string]map[string]bool

	RefetchDomainMap func()
	RefetchBlockMap  func()

	Logger logger.Logger

	MongoClient *mongo.Client

	Metrics *metrics.Metrics
}

func InitContainer(conf config.Config) AppContainer {
	client := db.NewClient() // Initialize the client using the imported package
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	context := context.Background()

	metricsInstance := metrics.New()

	dm := &map[string]Domain{}
	bm := &map[string]map[string]bool{}
	am := &map[string]map[string]bool{}
	gm := &map[string]bool{}
	gem := &map[string]bool{}

	// check admin domain
	adminDomain := conf.AdminDomain
	adminDomainEntity, _ := client.Proxy.FindFirst(
		db.Proxy.Source.Equals(adminDomain),
	).Exec(context)

	if adminDomainEntity == nil {
		client.Proxy.CreateOne(
			db.Proxy.Source.Set(adminDomain),
			db.Proxy.Target.Set("localhost:5000"),
			db.Proxy.Disable.Set(false),
			db.Proxy.Cache.Set(false),
		).Exec(context)
	}

	refetchDomainMap := func() {
		res, _ := client.Proxy.FindMany().Exec(context)
		clear(*dm)
		for _, proxy := range res {
			(*dm)[proxy.Source] = Domain{
				Id:          proxy.ID,
				TargetProxy: proxy.Target,
				Disable:     proxy.Disable,
				Cache:       proxy.Cache,
			}

			log.Println("Proxying " + proxy.Source + " to " + proxy.Target)
		}

	}

	refetchBlockMap := func() {
		res, _ := client.Block.FindMany().Exec(context)
		clear(*bm)
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
		res, _ := client.Allow.FindMany().Exec(context)
		clear(*am)

		for _, allow := range res {
			proxyId, _ := allow.ProxyID()

			if _, ok := (*am)[proxyId]; !ok {
				(*am)[proxyId] = make(map[string]bool)
			}

			(*am)[proxyId][allow.IP] = true
			log.Println("Allowing " + proxyId + ", " + allow.IP)
		}
	}

	refetchGuardMap := func() {
		resProxy, _ := client.Proxy.FindMany().Exec(context)
		resGuards, _ := client.Guard.FindMany().Exec(context)

		clear(*gm)

		for _, proxy := range resProxy {
			proxyId := proxy.ID
			(*gm)[proxyId] = false

			// find at leas one guard of proxy id
			for _, guard := range resGuards {
				guardId, _ := guard.ProxyID()

				if guardId == proxyId {
					(*gm)[proxy.ID] = true
					break
				}
			}

		}
	}

	refetchGuardExcludeMap := func() {
		res, _ := client.GuardExclude.FindMany().Exec(context)
		fmt.Println("Refetching guard exclude map")
		fmt.Println(res)
		clear(*gem)

		for _, ge := range res {
			gid, _ := ge.ProxyID()

			key := gid + ge.Path
			fmt.Println("Insertin key: ", key)
			(*gem)[key] = true

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
					refetchGuardMap()
					refetchGuardExcludeMap()
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
	refetchGuardMap()
	refetchGuardExcludeMap()

	fireInitialRefetch()

	mongoClient, err := mongo.Connect(context, options.Client().ApplyURI(conf.DatabaseUrl))
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

	startWatchGuard := func() {
		changeStream, err := mongoClient.Database("bastard-proxy").Collection("Guard").Watch(context, mongo.Pipeline{}, options.ChangeStream())
		if err != nil {
			panic(err)
		}

		for changeStream.Next(context) {
			// fmt.Println(changeStream.Current)
			refetchGuardMap()
		}
	}

	startWatchGuardExclude := func() {
		changeStream, err := mongoClient.Database("bastard-proxy").Collection("GuardExclude").Watch(context, mongo.Pipeline{}, options.ChangeStream())
		if err != nil {
			panic(err)
		}

		for changeStream.Next(context) {
			// fmt.Println(changeStream.Current)
			refetchGuardExcludeMap()
		}
	}

	go startWatchBlock()
	go startWatchProxy()
	go startWatchAllow()
	go startWatchGuard()
	go startWatchGuardExclude()

	logger := logger.New()

	return AppContainer{
		Config:           conf,
		Context:          context,
		PrismaClient:     client,
		DomainMap:        dm,
		BlockMap:         bm,
		AllowMap:         am,
		GuardMap:         gm,
		RefetchDomainMap: refetchDomainMap,
		RefetchBlockMap:  refetchBlockMap,
		Logger:           logger,
		MongoClient:      mongoClient,
		Metrics:          &metricsInstance,
		GuardExcludeMap:  gem,
	}
}
