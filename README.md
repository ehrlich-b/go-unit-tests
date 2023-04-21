# Go Unit Testing Pattern

I'm not saying this is "the way" to write unit tests in go, but it is a way that I have successfully used.

## Dependency injection

Dependency Injection is a design pattern that is essential when writing unit tests, and unit tests in Go are no different. The basic idea is to separate the dependencies of your code from the code itself, and to pass those dependencies into your functions or methods as parameters. This makes your code more modular and easier to test, as you can easily replace dependencies with mocks or fakes during testing.

In Go, this generally means injecting an interface for anything that is not going to exist on your test infrastructure (i.e. filesystems, external APIs, etc). Even if (e.g.) a filesystem does exist on your unit testing machine, it's generally much easier to isolate the function/struct under test via mocking, rather than having to create actual files on disk which would be subject to (e.g.) hardware failure.

See [internal/interfaces/fs.go](https://github.com/ehrlich-b/go-unit-tests/blob/main/internal/interfaces/fs.go) for an example injectable interface for a filesystem.

### How?

If you're familiar with unit testing in other languages, you may be familiar with the idea of a "DI container". This is hotly debated (of course) but the consensus (and my recommendation) is to [not use a DI container in go](https://stackoverflow.com/questions/41900053/is-there-a-better-dependency-injection-pattern-in-golang). Just inject interfaces "manually" through a New function.

Go doesn't have constructors, but by convention you should create a function called New[StructName]. E.g. (from [downloader.go](https://github.com/ehrlich-b/go-unit-tests/blob/main/internal/service/downloader.go#L17)):

```
func NewDownloader(fs interfaces.FS, httpClient *http.Client) *Downloader {
	return &Downloader{
		fs:         fs,
		httpClient: httpClient,
	}
}
```

Notice how we are "constructor injecting" all of our external dependencies.

## Surprising things that you don't need to inject in Go

Go comes with a mockable http test server. So whereas in most languages you would need to create an interface for all external web requests, in golang you can simply initialize a test server that will serve "real" responses.

See [internal/service/downloader_test.go](https://github.com/ehrlich-b/go-unit-tests/blob/main/internal/service/downloader_test.go#L14) for an example test http server.

I will still at times create a "external API" interface - if for some reason crafting the exact http responses from my server is difficult. This would look just like fs.go.

## Testify

From the [Testify](https://github.com/stretchr/testify) readme:
    
    Go code (golang) set of packages that provide many tools for testifying that your code will behave as you intend.

    Features include:
        * Easy assertions
        * Mocking
        * Testing suite interfaces and functions

The main things I want to highlight in this example are the "assertions" and "mocks". 

### Assertions

This one is pretty simple, golang doesn't have built in assertions. So rather than writing tons of:

```
if [some condition] {
    t.Fail("Some condition was not met in the test!")
}
```

You can use builtin assert functions that you're probably familiar with if you've done any kind of unit testing:

```
assert.Equal(t, a, b, "The two words should be the same.")
```

There are a wide variety of [testify assertions](https://pkg.go.dev/github.com/stretchr/testify/assert#hdr-Assertions).

### Mocks

Testify mocks allow for easy setup of mock classes. For example:

```
// Create a new instance of the mock
mockDependency := new(MockDependency)

// Set an expectation for the mock method
mockDependency.On("SomeMethod").Return(42)

// Inject the mock into the function being tested
result := myFunctionWithDependency(mockDependency)

// Assert that the function returned the expected result
assert.Equal(t, "expected result", result)

// Assert that the mock method was called once
mockDependency.AssertCalled(t, "SomeMethod")
```

Without a testify mock, you'd be "hand crafting" this mock class, so that it would return (e.g.) 42 under certain circumstances. This allows your tests to be much more declarative and readable. All function setups are nice chainable one-liners (e.g.) `mockDependency.On("SomeMethod").Return(42)`

A testify mock can be [written directly](https://pkg.go.dev/github.com/stretchr/testify/mock). But I do not recommend that at all. See the next section on Mockery.

## Mockery

[Mockery](https://github.com/vektra/mockery) is a command line tool that generates testify compatible mock classes. Mockery scans for any interface you've defined in code, and writes everything for you that you would have to otherwise manually write for a testify mock.

See [internal/interfaces/mocks/WriteCloser.go](https://github.com/ehrlich-b/go-unit-tests/blob/main/internal/interfaces/mocks/WriteCloser.go) for an example of how much boilerplate this saves you from having to write. 