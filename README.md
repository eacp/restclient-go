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

## Usage
TODO