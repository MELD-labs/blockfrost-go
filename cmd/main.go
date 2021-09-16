package main

import (
	"context"
	"fmt"
	"log"
	"os"

	blockfrostgo "github.com/blockfrost/blockfrost-go/pkg/api"
)

const (
	stakeAddrTestnet = "stake_test1uzaghtuxs0z569hnc68enjpnzy5tarqeg54k9p6rj5jaakq4dwqjg"
	stakeAddrMainnet = "stake1u9ylzsgxaa6xctf4juup682ar3juj85n8tx3hthnljg47zctvm3rc"
)

func main() {

	client := blockfrostgo.NewBlockfrostAPI(
		os.Getenv("API_KEY"),
		blockfrostgo.CardanoMainnet,
		false,
		nil,
		os.Stdout,
	)

	appinfo, err := client.Info(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appinfo)

	health, err := client.Health(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(health.IsHealthy)

	healthClock, err := client.HealthClock(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(healthClock.ServerTime)

	metrics, err := client.Metrics(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(metrics)
	metricsEndpoints, err := client.MetricsEndpoints(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(metricsEndpoints)

	accountAddr, err := client.Account(
		context.Background(),
		stakeAddrMainnet,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accountAddr)

	q := blockfrostgo.QueryParmasAPI{}
	accountRewardHist, err := client.AccountRewards(
		context.Background(),
		stakeAddrMainnet,
		q,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accountRewardHist)
}
