package main

import (
	"log"
	"temporal-learning"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// 1. Koneksi ke Temporal Server (local default di port 7233)
	c, err := client.Dial(client.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		log.Fatalln("Unable to connect to Temporal Server", err)
	}
	defer c.Close()

	// 2. Buat Worker baru yang mendengar antrian "TRANSFER_MONEY_TASK_QUEUE"
	w := worker.New(c, app.TaskQueueName, worker.Options{})

	// 3. Daftarkan Workflow dan Activities
	w.RegisterWorkflow(app.TransferWorkflow)
	w.RegisterActivity(app.WithdrawActivity)
	w.RegisterActivity(app.DepositActivity)

	// 4. Jalankan Worker
	log.Println("Worker is starting...")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Worker failed to start", err)
	}
}
