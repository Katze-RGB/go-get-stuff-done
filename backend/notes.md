initial takeaways
- go is really fast wow
- structs/interfaces are actually kinda different from classes.
- is go actually object oriented? I'm not 100% sure!
- was pretty painful to write/debug until I took a step back and setup air. That's more a result of me having like 12 total hours in the language though.
- JSON handling is super easy and painless wowie.
- go is also a lot more verbose than most other languages I've gotten paid to write in. Is there a better way to handle the boilerplate?
self-improvement
- ORMs are still mid. Starting with GORM rather than just building queries was a mistake maybe? Made testing a lot harder for sure
- do a better job seperating the business logic and data. This made testing pretty hard. Abstract it through an interface or something?
- look into db mocking libraries.
- all the querying should probably go somewhere other than in the handlers. Probably model funcs or something
- look more into how go talks to DBs. Seems like potential for connection overload/storms without some form of pooling?
- what's a goroutine? Is it just a coroutine but go?
notes for reviews
- the cat gif in the frontend is structurally critical, don't remove it
