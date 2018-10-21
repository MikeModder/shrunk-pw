package main

type unshortenedResp struct {
	Success bool   `json:"success"`
	Short   string `json:"short"`
	Full    string `json:"full"`
}
