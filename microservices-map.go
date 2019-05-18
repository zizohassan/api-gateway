package main

// you must define the base url for each micro service here
// middleware will check the first segment of url with the index
// if url like this /posts so your index must be posts
func microServiceMap() map[string]string{
	var urls = make(map[string]string)
	urls["posts"] = "http://127.0.0.1:6060"
	return urls
}
