package main

type config struct {
	Domain        string `json:"domain"`
	Port          int    `json:"port"`
	RedisURL      string `json:"redis_url"`
	RedisPassword string `json:"redis_password"`
	RedisDB       int    `json:"redis_db"`
}

type unshortenedResp struct {
	Success bool   `json:"success"`
	Short   string `json:"short"`
	Message string `json:"msg"`
	Full    string `json:"full"`
}

type shortenResp struct {
	Success bool   `json:"success"`
	Message string `json:"msg"`
	Short   string `json:"short"`
}
