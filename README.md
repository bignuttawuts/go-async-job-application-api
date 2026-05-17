cp .env.example .env
docker compose -f /deployments/docker-compose.yaml up -d
make run