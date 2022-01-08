**27**Jul 2019

# Testing in Go: Subtests

*Before we begin: The content in this article assumes knowledge of table-driven tests in Go. If you are unfamiliar with the concept, read [this article](https://ieftimov.com/post/testing-in-go-table-driven-tests) to familiarize yourself.*

With table-driven tests as the most popular testing approach, there is one annoying problem that every programmer will face: selective running of tests. That’s because the traditional approach of testing using table-driven tests in a single test function is not decomposable in granular subfunctions.

In other words, we cannot ask our `go test` tool to run a particular test case from a slice of them. Here's an example of a small test function that uses table-driven tests:

```go
func TestOlder(t *testing.T) {
	cases := []struct {
		age1     int
		age2     int
		expected bool
	}{
                // First test case
		{
			age1:     1,
			age2:     2,
			expected: false,
		},

                // Second test case
		{

			age1:     2,
			age2:     1,
			expected: true,
		},
	}

	for _, c := range cases {
		_, p1 := NewPerson(c.age1)
		_, p2 := NewPerson(c.age2)

		got := p1.older(p2)

		if got != c.expected {
			t.Errorf("Expected %v > %v, got %v", p1.age, p2.age, got)
		}
	}

}
```

There's no need to understand what the function under test does, although by looking at the tests you might figure it out. The question is: how do I run the `TestOlder` function with the second test case, without running the first case?

With the approach used above that is not possible. `go test -run regex` can target function names based on the supplied `regex`, but it has no way of understanding the internals of the function.

That's one of the reasons Marcel van Lohuizen in 2016 proposed the addition of programmatic sub-tests and sub-benchmarks. The changes were added to the language as of version 1.7. You can read more about it in [the proposal](https://github.com/golang/proposal/blob/master/design/12166-subtests.md) and [the related discussion](https://github.com/golang/go/issues/12166).

## What are subtests and how do they work?

Subtests are a construct in Go's `testing` package that split our test functions in granular test processes. They unlock useful functionality such as better handling of errors, more control over running tests, parallelism and a simpler and easier to organise code.

The actualization of subtests in the `testing` package is the [`Run` method](https://golang.org/pkg/testing/#T.Run). It takes two arguments: the names of the subtest and the sub-test-function. The name is an identifier of the subtests, which unlocks running a specific subtest using the `go test` command. Like with ordinary test functions, subtests are reported after the parent test function is done, meaning all of the subtests have finished running.

Note

Parallel tests are a topic we will explore in one of the next articles, so feel free to ignore that part for now. If you would like to get familiar with it now, there is a good section about it on [the official blog post about subtests](https://blog.golang.org/subtests).

Please note that there are some gotchas with parallel tests, that's why we are going to look into them separately.

Without going into too much detail, under the hood `Run` runs the function in a separate goroutine and blocks until it returns, or calls `t.Parallel` to become a parallel test. What happens under the hood and how subtests are architected is an interesting topic to explore, yet it's quite extensive to be covered in this article.

## How to use `t.Run`

Let's look at the `TestOlder` function again, this time refactored to use `t.Run` for each of the test cases runs:

```go
func TestOlder(t *testing.T) {
	cases := []struct {
		name     string
		age1     int
		age2     int
		expected bool
	}{
		{
			name:     "FirstOlderThanSecond",
			age1:     1,
			age2:     2,
			expected: false,
		},
		{
			name:     "SecondOlderThanFirst",
			age1:     2,
			age2:     1,
			expected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, p1 := NewPerson(c.age1)
			_, p2 := NewPerson(c.age2)

			got := p1.older(p2)

			if got != c.expected {
				t.Errorf("Expected %v > %v, got %v", p1.age, p2.age, got)
			}
		})
	}

}
```

There are a few notable changes. We changed the `struct` of the `cases`, to include a `string` attribute `name`. Each of the test cases has a name that describes the case itself. For example, the first case has the name `FirstOlderThanSecond` because `age1` is bigger than `age2` in the same case.

Next, in the `for` loop, we wrap the whole test in a `t.Run` block, where the first argument is the name of the test case. The second argument is a function which will (not) mark the test as failed based on the inputs and expected output.

If we run the test, we'll see something like this:

```bash
$ go test -v -count=1
=== RUN   TestOlder
=== RUN   TestOlder/FirstOlderThanSecond
=== RUN   TestOlder/SecondOlderThanFirst
--- PASS: TestOlder (0.00s)
    --- PASS: TestOlder/FirstOlderThanSecond (0.00s)
    --- PASS: TestOlder/SecondOlderThanFirst (0.00s)
PASS
ok  	person	0.004s
```

From the output, it's noticeable that right after `go test` runs `TestOlder` it spawns off two more test functions: `TestOlder/FirstOlderThanSecond` and `TestOlder/SecondOlderThanFirst`. It's worth noting that, `TestOlder` will not finish until these two functions do not exit.

The next few lines of the output paint that picture better, because the output is nested and it makes it clear that `TestOlder` is a parent to the other two functions. This is the effect on the output of spawning off two subtests in a test function. Also, we should also take note of the naming of the subtests – they are prefixed with the function that spawns them.

## Selectively running subtests with `go test`

As we already saw when using the traditional approach, running a specific test case is rather impossible. One of the pros of using subtests is that running only a specific subtest is very easy and intuitive.

Reusing the examples from before, running any of the subtests is just a matter of supplying the full name of the subtest: its parent test function, followed by a slash and the subtest name.

For example, if we would like to run the subtest `FirstOlderThenSecond` from the `TestOlder` test function, we can execute:

```bash
$ go test -v -count=1 -run="TestOlder/FirstOlderThanSecond"
=== RUN   TestOlder
=== RUN   TestOlder/FirstOlderThanSecond
--- PASS: TestOlder (0.00s)
    --- PASS: TestOlder/FirstOlderThanSecond (0.00s)
PASS
```

That's really it. Just by supplying the full name of the subtest, we can run a specific subtest. Given that the `-run` flag takes any regex, if we would like to run all of the subtests that are run in the `TestOlder` test function, we can do it by providing an “umbrella” regex:

```bash
$ go test -v -count=1 -run="TestOlder"
=== RUN   TestOlder
=== RUN   TestOlder/FirstOlderThanSecond
=== RUN   TestOlder/SecondOlderThanFirst
--- PASS: TestOlder (0.00s)
    --- PASS: TestOlder/FirstOlderThanSecond (0.00s)
    --- PASS: TestOlder/SecondOlderThanFirst (0.00s)
PASS
```

By supplying `TestOlder` to the `-run` flag, we run both the `TestOlder/FirstOlderThanSecond` and the `TestOlder/SecondOlderThanFirst` subtests.

## Shared Setup and Teardown

Another somewhat hidden side of subtests is unlocking the ability to create isolated setup and teardown functions.

The setup function is run to set up a test's state before the actual testing happens. For example, if we would have to open a connection to a database and fetch some records that would be used in the test, we would put such functionality in the setup function. In line with that, the teardown function of that test would close down the connection to the database and clean up the state. That's because teardown functions are run after the test finishes.

Let's use our `TestOlder` function from earlier to explore how setup and teardown are made and how they work:

```go
func setupSubtest(t *testing.T) {
	t.Logf("[SETUP] Hello 👋!")
}

func teardownSubtest(t *testing.T) {
	t.Logf("[TEARDOWN] Bye, bye 🖖!")
}

func TestOlder(t *testing.T) {
	cases := []struct {
		name     string
		age1     int
		age2     int
		expected bool
	}{
		{
			name:     "FirstOlderThanSecond",
			age1:     1,
			age2:     2,
			expected: false,
		},
		{
			name:     "SecondOlderThanFirst",
			age1:     2,
			age2:     1,
			expected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setupSubtest(t)
			defer teardownSubtest(t)

			_, p1 := NewPerson(c.age1)
			_, p2 := NewPerson(c.age2)

			got := p1.older(p2)

			t.Logf("[TEST] Hello from subtest %s \n", c.name)
			if got != c.expected {
				t.Errorf("Expected %v > %v, got %v", p1.age, p2.age, got)
			}
		})
	}

}
```

We introduce two new functions here: `setupSubtest` and `teardownSubtest`. While they do not contain any special functionality, understanding their invocation is essential here. Looking at the two lines where they are invoked, we can see that the `setupSubtest` is called right inside when the subtest is run.

The next line is where the `teardownSubtest` function is invoked, but this time using the `defer` keyword. This is a feature of Go that we use to our advantage here: `defer` allows us to invoke a function that will be executed at the end of the calling function. In other words, when the subtest function finishes, the `teardownSubtest` function will be invoked. This is the way Go makes setup and teardown functions really easy: they are not defined particularly, nor they contain any remarkable setup. They are two simple functions that use Go's built-in functionality.

If we run the test again, we can see the following output:

```bash
$ go test -v -count=1 -run="TestOlder"
=== RUN   TestOlder
=== RUN   TestOlder/FirstOlderThanSecond
=== RUN   TestOlder/SecondOlderThanFirst
--- PASS: TestOlder (0.00s)
    --- PASS: TestOlder/FirstOlderThanSecond (0.00s)
        person_test.go:33: [SETUP] Hello 👋!
        person_test.go:71: [TEST] Hello from subtest FirstOlderThanSecond
        person_test.go:37: [TEARDOWN] Bye, bye 🖖!
    --- PASS: TestOlder/SecondOlderThanFirst (0.00s)
        person_test.go:33: [SETUP] Hello 👋!
        person_test.go:71: [TEST] Hello from subtest SecondOlderThanFirst
        person_test.go:37: [TEARDOWN] Bye, bye 🖖!
PASS
ok  	person	0.005s
```

We can see that the setup is always run before the teardown. Looking at the output, it is clear that the assertion and error marking is done in between. In cases where the assertion would fail the particular subtest would be marked as failed and the error output would be presented at the end of the test run.

## `TestMain`

Before we wrap up this post, we will look at one last feature that the `testing` package has: `TestMain`.

There are times when our test file has to do some extra setup or teardown before or after the tests in a file are run. Therefore, when a test file contains a `TestMain` function, the test will call `TestMain(m *testing.M)` instead of running the tests directly.

# Coming soon!

I am planning a course on testing in Go. Would you like to get a (free) sneak peek while I am building it? Sign up now!

Subscribe

I respect your privacy. Unsubscribe at any time. No spam, ever.

Think of it in this way: every test file contains a “hidden” `TestMain` function, and its contents look something like this:

```go
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
```

`TestMain` will run in the main goroutine and it does the setup or teardown necessary around a call to `m.Run`. `m.Run` simply runs all of the test functions in the test file. `TestMain` will take the result of the `m.Run` invocation and then call `os.Exit` with the result as an argument. One important thing to note is that when we use `TestMain`, `flag.Parse` is not run so if our tests depend on command-line flags we have to call it explicitly.

There are a few use-cases where you would use `TestMain`: global startup and shutdown callbacks, or other state setups. You can read some more in Chris Hines’ 2015 [article](http://cs-guy.com/blog/2015/01/test-main/) on the topic.