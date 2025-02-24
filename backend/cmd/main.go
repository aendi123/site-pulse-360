package main

import (
    "os"
    "log"
    "time"
	"context"

    "github.com/gofiber/fiber/v2"
    "github.com/go-co-op/gocron/v2"

    "gitlab.diller.gl/site-pulse-360/backend/database"
    "gitlab.diller.gl/site-pulse-360/backend/internal/scanner"
    "gitlab.diller.gl/site-pulse-360/backend/internal/leaderelection"
)

func main() {
    database.ConnectDb()

    podName := os.Getenv("POD_NAME") // gets defined in deployment.yaml
    namespace := os.Getenv("NAMESPACE") // gets defined in deployment.yaml
    leaderElection := leaderelection.NewKubernetesLeaderElection(context.Background(), podName, "leaderelection-scanner", namespace)

    // init listener for leader election
    listener := leaderelection.NewListener(func(cancelChan <-chan struct{}) {
        // init cronjob.
        s, err := gocron.NewScheduler()
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        s.NewJob(
            gocron.DurationJob(1*time.Minute),
            gocron.NewTask(scanner.Scan),
            gocron.WithStartAt(gocron.WithStartImmediately()),
        )
        s.Start()
        <-cancelChan // block till leader changes
        // run any cleanups, canceling the cronjob.
        err = s.StopJobs()
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
    })
    
    leaderElection.AddListener(listener)
    go leaderElection.RunElection() // run leader election async

    app := fiber.New()
    setupRoutes(app)    
    app.Listen(":3000")
}