# Go restful test

1 CRUD endpoint involving GORM, mysql, fiber, migration & seeder, swaggo, cobra, docker, unit & integration testing

## Prerequisites

Minimum system requirements

```
Go 1.19
Mysql 8~
```

## Installing & Running

Copy .env-example to .env and fill the required configuration

```
go mod tidy
go run main.go
```

Surf http://localhost:3000/swagger
To see listed endpoints available & try it on the go

U can run the migrate & seeds via command line, for example use
`./go-restful-test migrate drop_product`
`./go-restful-test seed product 10`

## Running the tests

In order to run entire test, user should have & granted docker usage

```
go test ./...
```

## Technical questions

**Q:** From 1 (Fundamental Awareness) to 5 (Expert), how well do you know about Go Language? and tell us about functions, concurrency and pointers.

**A:** 4.  It's like my daily food, from goroutines, channels to newest go feature named generic i've used it. In wide aspect, functions is made to prevent bloated syntax & store same task with parameterized variable, so we can reuse & call it many times. About concurrency, did i recall goroutine & channel before? yes that's fall under Go system, with those help, we can run tasks separately & communicate with each other using channel, and for the last one, pointer is about storing a variable, function, or other pointer by assigning the memory address of referenced value instead of possessing the value it self.


**Q:** From 1 (Fundamental Awareness) to 5 (Expert), how well do you know about Git version control? and tell us what command you most frequently used

**A:** 4, commit, pull, push, checkout. Actually, in local environment we don't need to manually enter those commands, we can use GUI instead which more easier, faster & interactive, eg: a source control feat in VSCode,

**Q:** From 1 (Fundamental Awareness) to 5 (Expert), how well do you know microservices architecture? and tell us the advantages and disadvantages of using microservices and monoliths.

**A:** 4, Splitting a service into few lil service, micro managing process by it's functionality like (separating api services, authentication, database, adaptors, notificator, cron, etc) it's often to see in containerized application. Microservice brings more modular, flexible, scalable & maintainable when it come to big architecture, with those comes great complexity too, sometime it can mess the infra topology, integration need more work, network congestion, more automated deployments, etc. All of these are relative because that's depends on how well infra is structured, how well the maintainer can provide good readibility & mitigation system, etc.


**Q:** From 1 (Fundamental Awareness) to 5 (Expert), how well do you know event-driven architecture? and tell us which event streaming platform you’ve used.

**A:** 4, More like microservice architecture, this arch involve in how data flow can be separated, by emitting a communcation to a service, then process it, forward it to the next service and so on is how event-driven work. I've used message broker service like Rabbitmq, both amqp(queue) & mqtt(subscription) as protocol.


**Q:** From 1 (Fundamental Awareness) to 5 (Expert), how well do you understand MySQL query? and tell us what we need to know to optimize query speed

**A:** 4, At least i know how to create indexing, custom view, trigger event function / procedure, make relationship of tables & join them. As for optimization in query, we need to select specific column that's only necessary, applying index to prevent entire scanning table, use proper join to prevent duplicate data set, normalize table to prevent nested data & analyze it through `explain` to profile your query. Beside from query, we can manually implement caching system too where we don't need to execute sql or others, rather we use the stored in memory data that have been queried before. When we have more resource & big data coming is on demand, invest it to infra, create cluster database to distribute the pain of single instance process.


## FAQ:

**Q:** Docker test?

**A:** Yes, bcs it's pain in the ass trying to mock up all interface, when testing database / environment related system it's better to go all out using integrated test instead


**Q:** Why there's no api versioning?

**A:** Api versioning could be ran as separate instance with help of api gateway management like traefik or kong, as development goes on, it will be hindrance to handle backward compability in same upstream.

