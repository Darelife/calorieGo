# calorieGo
Learning how to use golang as a backend server + production level backend basics (caching, ratelimiting, batch processing, reverse proxy, databases, queues, and containerization)

## todo
I'm using ChatGPT btw, trying to understand how everything works. I myself don't know (or have forgotten) about a lot of the imp things. So, i'll go and revise them
- [ ] Revise GoLang
- [ ] Revise Redis
- [ ] Revise Docker
- [ ] Learn how to use Nginx (Reverse Proxies)
- [ ] Learn how to use RabbitMQ
- [x] Graceful shutdown & signal handling
- [ ] Rate limiting (token bucket, leaky bucket)
- [ ] Caching strategies
- [ ] External API failures


## running
```bash
go run ./cmd/api
curl -X POST http://localhost:8080/v1/food/barcode \
  -H "Content-Type: application/json" \
  -d '{"barcode":"737628064502"}'
```

```bash
sudo docker exec -it calorie-redis redis-cli
```

```bash
docker compose up -d
docker compose down
docker ps
docker ps -a
```
