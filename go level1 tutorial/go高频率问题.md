## init()跟 sync.Once 的区别

Package init() functions are guaranteed by the spec to be called only once and all called from a single thread (not to say they couldn't start goroutines, but they're thread safe unless you make them multi-threaded).

The reason you'd use sync.Once is if you want to control if and when some code is executed. A package init() function will be called at application start, period. 

sync.Once allows you to do things like lazy initialization, for example creating a resource the first time it is requested (but only once, in case multiple "first" requests come in at the same time) rather than at application start; or to only initialize a resource if it is actually going to be needed.
