package main

type LocationAreaResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []LocationArea
}

type LocationArea struct {
	Name string
	URL  string
}

type Config struct {
	NextLocationAreaURL     *string
	PreviousLocationAreaURL *string
}
