# Rest Client Wrappper in Go
A wrapper for making REST API calls to the appropiate host

## Context
I was once making a client for a REST API, and I always
had to write the base url of the api (host).

Most people would declare it as a constant and write their methods
or functions as:

```go
resp, err := httpClient.Get(baseURL + "somethin/something")
if err != nil {
    // ...
}
```

But this could get cumbersome. What if we culd wrap this to remember
the host and other configurations?

```go
resp, err := rest.Get(something)
if err != nil {
    // ...
}
```

This package is more or less inspired in 
[Retrofit](https://square.github.io/retrofit/), however, it
**does not** intent to be a full replacement.

It is also similar to [The Axios Package](https://github.com/axios/axios), most notably because both [Axios](https://github.com/axios/axios) and this package allow you to set up a 
[default host endpoint](https://github.com/axios/axios#creating-an-instance)
so the requests are less repetitive.

## Usage

Import and use like this:
```go
import "eacp.dev/restclient"

//...
githubAPI := restclient.New("api.github.com")

// Request a specific license
resp, err := githubAPI.Get("licenses/mit") 
```

In general, it is supposed to work like the [`net/http` package](https://pkg.go.dev/net/http/).
It contains wrappers for all the public functions in this package, such as [`http.Get()`](https://pkg.go.dev/net/http/#Get)
or [`http.Post()`](https://pkg.go.dev/net/http/#Post).
You can make a `RestClient` and use it to make requests as if it was the [`net/http` package](https://pkg.go.dev/net/http/). 

```go

//...
githubAPI := restclient.New("api.github.com")

// Request a specific license
resp, err := githubAPI.Get("licenses/mit") 
```

## Comming "soon"(ish)

### Concurrent multiple requests
I would like to replicate the [axios.all](https://github.com/axios/axios#concurrency-deprecated) function. 
It was deprecated in [axios](https://github.com/axios/axios), but I believe it could have a nice place in Go :)

```go

//...
githubAPI := restclient.New("api.github.com")

// Request multiple licenses
githubAPI.GetAll("license/mit", "licenses/gpl-3.0", "licenses/bsd-3-clause")
```

This functions still needs some though, because I'm yet to decide its return types

### Authentication and tokens
The client could be set up to use tokens and other auth methods
