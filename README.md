# Temporal Money Transfer (Example)

This project demonstrates a simple money transfer application using [Temporal](https://temporal.io/).

## Prerequisites

1.  **Go** installed (v1.20 or later).
2.  **Temporal Server** running locally. The easiest way is using the Temporal CLI:
    ```bash
    temporal server start-dev
    ```
    *This will start the server and the Temporal Web UI at [http://localhost:8080](http://localhost:8080).*

## Project Structure

- `workflow.go`: Contains the Workflow logic (`TransferWorkflow`) and the Activity logic (`WithdrawActivity`, `DepositActivity`).
- `worker/`: Contains the code to start a Temporal Worker that listens for tasks.
- `starter/`: Contains the code to trigger/start a Workflow execution.

## How to Run

### 1. Start the Worker
The worker must be running to process the workflow tasks.
```bash
go run worker/main.go
```
You should see `Worker is starting...`.

### 2. Start the Workflow
In a new terminal window, run the starter to initiate a transfer:
```bash
go run starter/main.go
```
You will see logs indicating the workflow has started and eventually the final result.

## How It Works

1.  **Workflow Definition**: The `TransferWorkflow` orchestrates the process. It first calls `WithdrawActivity` and then `DepositActivity`.
2.  **Activities**: These are the actual units of work (e.g., updating a database). In this demo, they just simulate a database operation with `time.Sleep`.
3.  **Temporal Server**: Acts as the orchestrator. It maintains the state of the workflow and tells the Worker which task to execute next.
4.  **Worker**: Polls the `TRANSFER_MONEY_TASK_QUEUE` for tasks. When it gets one, it executes the corresponding function.
5.  **Fault Tolerance**: If the Worker crashes during the transfer, the Temporal Server will notice and re-assign the task to another worker (or wait for this one to restart), ensuring the transfer is never "lost" in an inconsistent state.

## Viewing Progress
Open [http://localhost:8080](http://localhost:8080) to see your workflow execution in the Temporal Web UI. You can see the history of events, input/output data, and any errors that occurred.
