package main

func main() {
	ctx, redisClient := GetClient()
	defer redisClient.Close()
	// redisString(ctx, redisClient)
	redisHash(ctx, redisClient)
	redisList(ctx, redisClient)
	subscribe(ctx, redisClient)
}
