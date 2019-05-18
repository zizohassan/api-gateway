# How it work

![Image of Yaktocat](https://i.ibb.co/JjMZBfY/Untitled-Diagram-1.png)

# How To use

open `microservices-map.go` define new base url like this 

`urls["posts"] = "http://127.0.0.1:6060"`
 
 here i add posts micro service  with url `http://127.0.0.1:6060`
 
 in posts micro service i have these routes
 
 ```go
r.POST("posts/create", createPost)
r.POST("posts/update/:id", updatePost)
r.GET("posts/delete/:id", deletePost)
r.GET("posts/show/:id", showPost)
r.GET("posts", getPost)
```

now you can access all these routes with api-gateway

just serve api gate way and you can hit any api from `post` micro service form gateway like
`
    http://127.0.0.1:5050/posts
`

note the post routes segment 1 is `posts`  the same key we put in 
`microservices-map.go` file

the routes 1,2,3 is need token to access them 4,5 not need token 

we just need to go `auth-action-slice.go` and add the action in this case action are segment 

number 2 in url `/posts` -> segment 1 `create` -> segment 2




