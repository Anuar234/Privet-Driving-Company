# Telegram Mini App – Go + HTML

This is a simple Telegram Mini App project.

## Structure

- `frontend/`: HTML file loaded inside Telegram
- `bot/`: Go bot that sends the Mini App
- `server/`: (Optional) Backend for handling submitted data

## How to Run

1. Run `go mod tidy` to install dependencies.
2. Serve the frontend using:

   ```bash
   cd telegram-mini-app
   go run bot/main.go
