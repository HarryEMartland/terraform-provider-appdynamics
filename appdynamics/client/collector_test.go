package client

import (
	"fmt"
	godotenv "github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
)

func TestCollector(t *testing.T) {
	fmt.Println("Testing colllector")
	godotenv.Load("../.env")

	baseURL := os.Getenv("APPD_CONTROLLER_BASE_URL")
	secret := os.Getenv("APPD_SECRET")

	client := AppDClient{
		BaseUrl: baseURL,
		Secret:  secret,
	}

	id, err := client.CreateCollector(&Collector{
		Name:      "dummytest-2",
		Type:      "MYSQL",
		Hostname:  "host",
		Port:      33,
		Username:  "aaa",
		Password:  "bb",
		AgentName: "dbagent",
		Enabled:   true,
	})

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("Successfully created collector id = %v \n", id)

	collectorId, _ := strconv.Atoi(id)

	_, getCollectorErr := client.GetCollector(collectorId)

	if getCollectorErr != nil {
		fmt.Println(getCollectorErr)
		t.Fail()
	}

	fmt.Printf("Successfully read collector id = %v \n", collectorId)

	deleteCollectorErr := client.DeleteCollector(collectorId)

	if deleteCollectorErr != nil {
		fmt.Println(deleteCollectorErr)
		t.Fail()
	}

	fmt.Printf("Successfully deleted collector id = %v \n", collectorId)
}
