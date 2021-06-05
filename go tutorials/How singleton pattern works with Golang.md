## Problem to be solved

We have to open a connection to the database and keep it in a single  instance so we do not overload the database and have problems or errors  with the limit of concurrent connections exceeded, and of course, of  optimizing our code. We know that a connection is very costly  computationally.

Pooling natively, then we need to control all the connections that our  application will open to communicate with it, so in every action we  execute, be it a select, delete, insert or update, we need a connection  we will use the singleton pattern in the memory of atomic form. Examples of ***Singleton async\*** and ***sync connections\***.

Before we implement our solution using ***Singleton Pattern\*** to solve the problem of connection pooling, let’s look at some other examples of the need to use the singleton pattern.

Here are some examples of ways to try to protect our ***global variable\***.

## Protecting Global Variable Example 1

No Thread Safe

```Go
// type global
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {

	if instance == nil {

		instance = make(singleton) // <-- not thread safe
	}

	return instance
}

```



## Protecting Global Variable Example 2

Thread Safe

```Go
var lock = &sync.Mutex{}

// type global
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {

		instance = make(singleton) // <-- thread safe
	}

	return instance
}
```



## Protecting Global Variable Example 3

Thread Safe, the best way to our singleton

```Go
var once sync.Once

// type global
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {

	once.Do(func() { // <-- atomic, does not allow repeating

		instance = make(singleton) // <-- thread safe

	})

	return instance
}
```



Above we show three ways to protect our global variable “instance”, for  synchronous solutions example 1 will already work, for asynchronous  solutions example 2 and 3 are the most indicated being that example 3 is the best solution to protect from “[***race condition\***](https://blog.golang.org/race-detector)“ by ensuring that the instance is declared a single time by becoming atomic.

The purpose of the article is to show exactly how to try to protect the ***global variable,\*** in Golang we have several ways of doing this.
Below are a few more examples to complement the reasoning, showing how we instantiate our **singleton**, be it *synchronously or asynchronously* using *goroutine*.

## Examples of how to instantiate our Singleton

The code below is a good example of openings for multiple connections in a  synchronous way, depending on how we implement our connection instance  the bank would not support.


Multiple connections open in the database in a Synchronous way

```Go
func main() {

	// conexao
	Db := Conn.Connet()

	// busca sempre
	// da memoria
	Db = Conn.Connet()
	Db = Conn.Connet()
	Db = Conn.Connet()
	Db = Conn.Connet()
	Db = Conn.Connet()
	Db = Conn.Connet()
	Db = Conn.Connet()
  
	// da memoria
	fmt.Println("connect", Db.Ping())
	fmt.Println("connect", Conn.Connet().Ping())
	fmt.Println("connect", Db.Ping())
	fmt.Println("connect", Conn.Connet().Ping())
	fmt.Println("connect", Db.Ping())
	fmt.Println("connect", Conn.Connet().Ping())

	time.Sleep(time.Second * 1)
}
```



Now we will see an example of opening ***connections asynchronously*** using ***goroutine***. We made hundreds of goroutines and thousands of concurrent connections  in our example below, works 100% without crashing and or bugging the  amount of database connections.


Asynchronous connections

```Go
// na base
func main() {

	// conexao
	Db := Conn.Connet()

	// Exemplo 2 
	// criando 200 goroutines
	for x := 0; x < 200; x++ {
		go func(x int) {
			for j := 0; j < 10000000; j++ {
				fmt.Printf("Goroutine2 Connect: %d -> %d", x, j)
				fmt.Printf(" login: %s", Conn.Connet().GetUserEmail(x))
				time.Sleep(time.Millisecond * 150)
			}
		}(x)
	}
	fmt.Scanln()
}
```

# Goroutines and the Singleton Pattern

When we think of possible solutions for the implementation of the ***singleton pattern\*** using Golang, we bump into the Goroutines it will be responsible for providing our code to execute ***asynchronously and concurrently\*** and when we use Goroutine in our application our whole way of thinking and implementing changes, that is is no longer a ***synchronous\*** application and because of this we have to think of the Golang way of being.

Goroutines is a powerful resource and when used correctly it becomes a strong ally to fight the day-to-day battles. Every time we implement codes that use competition we have some known scenarios that we have to deal with,  program scope, global variables, locations, parameter passing, pointers  all this has to be handled so we can work with competition in a correct  and optimized way .

A good example of possible problems using competition is the use of ***global variables\***. Due to the goroutines our implementation of the ***Singleton Pattern\*** and our possible solutions will be written to accept the use of  competition. Let’s write our code using good programming practices to  try to mitigate the possible stealth bugs that may occur in runtime of  your program written in ***Golang\***.

## 200 goroutines being initialized

In our example below we are ***creating 200 goroutines\*** and putting all of them in competition and ten thousand interactions  are being made in our database ie “select email from login where id =?”.

Goroutine making hundreds of calls in our driver

```Go
	for x := 0; x < 200; x++ {
    
		go func(x int) {

			for j := 0; j < 10000; j++ {

				fmt.Println(" login: ", Conn.Connet().GetUserEmail(x))
				time.Sleep(time.Millisecond * 150)
			}
		}(x)
	}
```

Good understanding of the goroutines will help us write better and more  powerful codes in Golang, this I have no doubt. In a next article we  will describe some cases using goroutines, I am eager to write about  Goroutines and show how we solved some of our problems in the company.

The subject about ***Singleton Pattern\*** is so interesting that I had to do a much deeper research on the  subject before writing this article and as I use a lot in my day to day I decided to collaborate a little showing some important points of this  pattern and why a lot of the times it is considered an ***Anti Pattern\***. Describing the technical details in the article we clearly notice  several good practices that we can use on our day when we are coding in  Golang, good practices, clean implementation forms, less complex and  less costly implementations and with higher performance.

# What is Singleton Pattern ?

The transcript of what is a ***Singleton Pattern\*** would be: “**Singleton** is a software design standard. This standard guarantees the existence  of only one instance of a class, while maintaining a global point of  access to its object.”
**Singleton** is a design pattern that restricts instantiation to an object, we have to ensure that this occurs only once.

Basically ***singleton\*** is a way to use ***global variables\***. We know how dangerous the use of *global variables* is, our code is **vulnerable** to the access of the global variable or in any part of the system we  can change its value. So when we try to debug our program it will not be an easy task to figure out which code path leads to the current state,  which is why I do not consider ***Singleton Pattern\*** an ***Anti Pattern\***, but a way to ***protect global variables.\***

However, the problem with ***Singleton\*** using competition that will be our goal, in a ***multi-threaded\*** environment, initialization must be protected to avoid reboot, atomically.

The ***Singleton Pattern\*** is a feature derived from the ***object oriented\*** paradigm, so how will we implement Golang if it does not have OO support?

To answer this question we have to understand that **Object Oriented** Programming is a concept and can be implemented in any programming  language even though it is of other paradigms. It is clear that the  level of abstraction and difficulty becomes an arduous task greatly  increasing the level of complexity of the code depending on the  language, our intention is purely didactic in order to better understand the proposed scenario when we speak in **singleton pattern**.

# go run -race singleton.go

> It executes all the codes using -race as parameter: go run -race ..

The “***-race detector\***” is a feature we have available in *Golang* to detect improper accesses in memory when we are using competition in  our application. It is possible to generate report that contains stack  traces for conflicting accesses, as well as piles in which the  goroutines involved were created, I will soon create an article  addressing exactly this subject and we will talk about “***Profiling\***” in Go. 
Below are the ways we can make the “***-race\***” call.

```
$ go test -race seupkg    // to test the package
$ go run -race seusrc.go  // to run the source file
$ go build -race seucmd   // to build the command
$ go install -race seupkg // to install the package
```

# Singleton Pattern in Golang

The solution presented below would be ideal if our application were  synchronous, and the problem would be solved with code below, but as our goal is to implement using competition the solution below is far from  ideal.
Let’s check the code below and start our implementation possibilities:

Singleton Not Thread Safe

```Go
type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect() *DriverPg {

	if instance == nil {
		// <--- NOT THREAD SAFE / Quando usarmos Goroutine
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	}

	return instance
}

func main() {

	// chamada
	go func() {

		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 600)
			fmt.Println(*Connect(), " - ", i)
		}
	}()

	go func() {

		fmt.Println(*Connect())
	}()

	fmt.Scanln()
}
```



In the above implementation, we have our first approach implementing ***Singleton\*** in ***Golang\*** and with goroutines to simultaneously run our ***Singleton\***.

The problem visible in this code is that ***several goroutine\*** routines could evaluate the first check and all would create a  singleton instance and replace one another. There is no guarantee of  which instance will be returned in the code above and other additional  operations on the instance may be inconsistent with developer  expectations and stealth problems may occur at run time.

***Too bad this approach\***, several very subtle errors can occur, if references to the singleton  instance are being maintained through code, there is a great chance that there are potentially multiple instances of the type with different  states, generating potential code behaviors. It also becomes a real  nightmare during debugging and it becomes really difficult to detect the bug, since at debug time nothing seems to be wrong due to runtime  breaks, minimizing the potential of a “***Not Thread Safe\***” execution, obfuscating totally the problem for who is coding.

# Locks with Mutex

In the code below is a poor solution for attempting to solve the “***Thread Safe\***” problem. In fact, this solves the “***Thread Safe\***” problem, but creates other serious potential problems. It introduces a  containment in the goroutines executing an aggressive blocking of the  whole function, let’s check the code below:

Singleton Using Go and Lock with Mutex

```Go
// nosso lock mutex
var lock = &sync.Mutex{}

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect() *DriverPg {
	// <--- Desnecessario a lock
	// se a instancia já tiver
	// sido criada muito agressivo
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		// ainda não é a melhor implementação devido
		// os bloqueios
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	}
	
	return instance
}

func main() {

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 600)
			fmt.Println(*Connect(), " - ", i)
		}
	}()

	go func() {
		fmt.Println(*Connect())
	}()

	fmt.Scanln()
}
```

The “Thread Safe” issue has been resolved with the above implementation using ***sync.Mutex\*** where the Lockdown occurs before creating the ***Singleton\*** instance. The big problem with this approach is excessive blocking even when it would not be necessary to do so in case the instance has  already been created and should simply have returned the singleton  instance. If our program is written to become highly concurrent, this  can generate a bottleneck, since only one goroutine routine can get the *singleton* instance at a time, making it our slowest solution.

Let’s check another solution, because it is not our best approach above.

# Check-Lock-Check in Go

One way to improve and ensure a minimum lock and still be safe for the goroutine is to use the pattern called “**Check-Lock-Check**” when acquiring locks. But you have to use ***Mutex with atomic\*** so it **is not “\*the tread safe\*”** otherwise it will become a “***the thread not safe\***”. We use this same Patter in C and C ++. The pattern works with the idea of checking first, *to minimize any aggressive blocking*, since an IF statement is less expensive than locking. Next time, we  would have to wait and get the exclusive lock so that only one execution is inside that block at a single time. With the first scan and  exclusive lock, there may be another goroutine that has the lock, so we  would need to double-check inside the lock to avoid replacing the  instance with another. ***Check the code below\***:


Singleton Using Go and Blocking with Mutex less aggressive and not yet fully atomic

```Go
var lock = &sync.Mutex{}

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect() *DriverPg {
	
	// ainda não está perfeita, não é totalmente atomico
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	}
	return instance
}

func main() {

	go func() {

		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 600)
			fmt.Println(*Connect(), " - ", i)
		}
	}()

	go func() {

		fmt.Println(*Connect())
	}()

	fmt.Scanln()
}
```

The above approach is the best so far, ***but still not perfect\***. There ***is no atomic check\*** on the storage state of the instance. Taking all technical  considerations into account, this is still not perfect. But it is much  better than the initial approaches.

Using the ***sync/atomic\*** package, we can load and set atomically a flag that indicates whether or not we initialize our instance. 
**Check the code below:**


Singleton “**Thread Safe**” & ensuring the uniqueness of our instance with **sync package**

```Go
// manter o estado
var atomicinz uint64

// lock mutex
var lock = &sync.Mutex{}

// driver
type DriverPg struct {
	conn string
}

// instancia Global
var instance *DriverPg

// funcao retornando
// o ponteiro de nossa
// struct
func Connect() *DriverPg {

	// garantindo que já entrou
	if atomic.LoadUint64(&atomicinz) == 1 {

		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	// entra somente uma
	// únic vez
	if atomicinz == 0 {

		instance = &DriverPg{conn: "DriverConnectPostgres"}
		atomic.StoreUint64(&atomicinz, 1)
	}

	return instance
}

func main() {

	// chamada
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println(*Connect())
	}()

	// 50 goroutine
	for i := 0; i < 50; i++ {
		go func(i int) {
			for {
				time.Sleep(time.Millisecond * 60)
				fmt.Println(Connect().conn, " - ", i)
			}
		}(i)
	}

	fmt.Scanln()
}
```

The ***sync/atomic\*** library allows us to flag and store content securely and ensuring its uniqueness, much like ***sync.Map\***, where it is storing in a storage and accepting competition in its  implementation. The problem we had to use more functionality, more  feature and a little slower implementation sequentially.

# sync.Once clean and powerful

We have the type “***Once\***” in the ***sync\*** library, remembering that this native library in Golang is very  powerful and we have to exploit it as much as we can, the sync.Once  object will perform an action exactly once and not more, it was what was missing for our code to get even more powerful and clean.

```Go
// call somente
// uma unica vez
var once sync.Once

type DriverPg struct {
	conn string
}

// variavel Global
var instance *DriverPg

func Connect() *DriverPg {

	once.Do(func() {

		instance = &DriverPg{conn: "DriverConnectPostgres"}
	})

	return instance
}

func main() {

	// chamada
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println(*Connect())
	}()

	// 100 goroutine
	for i := 0; i < 100; i++ {

		go func(ix int) {
			time.Sleep(time.Millisecond * 60)
			fmt.Println(ix, " = ", Connect().conn)
		}(i)
	}

	fmt.Scanln()
}
```

Singleton Go using ***sync.Once\*** the clean code / “The Thread Safe”

With this approach and proposed scenario our code in addition to getting  cleaner was much better, the sync.Once function guarantees the  uniqueness of our instance, our code can now have ***100 goroutines\*** or more according to its need, puts them in competition not we will  have Thread Safe problems or aggressive checking. A simple and secure  way to write the Golang code for implementing Singleton Pattern.

# init () Another approach

Another valid approach is to use ***init\***( ), it runs only once and is called before the maim function. Check the code below:

```Go

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect() *DriverPg {

	instance = &DriverPg{conn: "DriverConnectPostgres"}
	return instance
}

func init() {

	Connect()
}

func main() {

	// chamada
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println(instance.conn)
	}()

	go func() {

		fmt.Println(*Connect())
	}()

	// 100 goroutine
	for i := 0; i < 100; i++ {

		go func(ix int) {
			time.Sleep(time.Millisecond * 60)
			fmt.Println(ix, " = ", instance.conn)
		}(i)
	}

	fmt.Scanln()
}
```

Singleton using Init()

But this approach is a disadvantage when we use ***init\***( ). Note clearly that it is not secure, nothing prevents you from making a direct call in the “***Connect\***” function as occurred in ***line #47\***, in addition there is a limitation in the use of init in relation to its load time and the most important in *Golang* we can have multiple ***init\***( ) running not only on one file or package but on multiple, and there is a running order between them.

The ***init( )\*** function does not accept arguments nor returns any value. In contrast to our approach using ***sync.Once\***, the ***init ( )\*** identifier is not declared, so it can not be referenced.

The best scenario is to write codes that do not depend on the boot order, in previous versions of *golang* there were some complaints and several problems reported, do not write codes in an ***init ( )\*** that you need guarantees of execution at any given time. The solution  when it needs explicit assurance is to write explicit calls.

For more details have a page only of this in *Golang* https://golang.org/doc/effective_go.html#initialization, it is worth the detailed reading about ***init( )\***, it is a powerful and robust implementation in the current versions of Golang but always it’s good to stay tuned.

# Variable receiving function

Another valid approach is to use a global variable to get the function in a ***global scope\***, we know that in Golang the variables are assigned and declared before the ***init ( )\*** and main function call, so in this approach the method returns exactly one instance. Check the code below:

Singleton with global variable receiving function

```Go

type DriverPg struct {
	conn string
}

var instance *DriverPg

var instanceNew = *Connect()

func Connect() *DriverPg {

	if instance == nil {

		// <--- NOT THREAD SAFE
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	}

	return instance
}

func main() {

	// chamada
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println("goroutine 1: ", instanceNew.conn)
	}()

	go func() {

		fmt.Println("goroutine 2: ", *Connect())
	}()

	fmt.Scanln()
}
```

This approach is flawed because nothing guarantees that the function will be called again in some part of the code, as occurred in line #25 the  function was called again, in line #30 is made the instance of our  singleton, but nothing guarantees this uniqueness .

# Conclusion

The ideal is undoubtedly the use of ***sync.Once\*** that guarantees us the uniqueness and that is “**Thread Safe**” guaranteeing that a “[***Race Condition\***](https://blog.golang.org/race-detector)***”\*** does not occur, it only allows the function to be executed only once Golang  flexibilized and automated all the complexity we would have in other  languages if we were to work with competition and simultaneity. Golang  really became powerful in these approaches, making it simple to  implement and understand.

When we talk about competition our whole way of thinking and coding solutions using ***Golang\*** changes drastically. ***There are several scenarios\*** that we need to ***apply standards\*** and practices in our projects to take full advantage of the power that Golang offers.