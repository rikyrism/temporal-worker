package main

import (
	"context"
	"log"
	"temporal-learning"

	"go.temporal.io/sdk/client"
)

func main() {
	// 1. Hubungkan Client ke Temporal Server
	c, err := client.Dial(client.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		log.Fatalln("Unable to connect to Temporal Server", err)
	}
	defer c.Close()

	// 2. Persiapkan data dummy
	transferInput := app.TransferInput{
		FromAccount: "Budi",
		ToAccount:   "Ani",
		Amount:      50000.0,
		ReferenceID: "TRX-001",
	}

	// 3. Konfigurasi cara menjalankan workflow (id unik dan task queue)
	workflowOptions := client.StartWorkflowOptions{
		ID:        "transfer-money-workflow", // ID unik untuk workflow ini
		TaskQueue: app.TaskQueueName,
	}

	// 4. Jalankan Workflow
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, app.TransferWorkflow, transferInput)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Printf("Started workflow with WorkflowID: %s, RunID: %s\n", we.GetID(), we.GetRunID())

	// 5. Tunggu hasilnya (Opsional: dalam praktek nyata biasanya tidak menunggu sync begini)
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Workflow failed to complete", err)
	}

	log.Printf("Workflow Result: %s\n", result)
}
