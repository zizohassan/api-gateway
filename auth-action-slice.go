package main

// here you must define only auth action
// like if you want to only auth user can create posts
// your url must be like this /posts/create
// create index posts and slice with actions {create}
func authActionsSlice()  map[string][]string {
	var actions = make(map[string][]string)
	actions["posts"] = []string{"create" , "update" , "delete"}
	return actions
}