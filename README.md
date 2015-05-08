# hystrix-go

This is sample project to use [hystrix-go](https://github.com/afex/hystrix-go).
This repository also includes sub-system which main process tries to send request.
So you can try its functionality by hand. 

## Install

To use hystrix-go, 

```bash
$ go get github.com/afex/hystrix-go/hystrix
```

## Usage

To run main (with hystrix circuit breaker),

```bash
$ go run main.go
```

To run sub system, it would takes time to response,

```bash
$ cd sub && go run main.go
```

To access to main system,

```bash
$ for i in $(seq 5); do curl -x '' localhost:8080 &; done
```

Circuit Breaker will open after certain number of requests (`RequestVolumeThreshold`) and error rate is higher ahtn threshold (`ErrorPercentThreshold`). While opening, all requests are failed. After some certain time (`SleepWindow`), breaker try to send single test request to sub system, and request is succeeded, breaker will close.

You can also set request Timeout and the number fo conccurent requests. 

## References

- [Trying out hystrix-go](https://gist.github.com/marcusolsson/2ef585c2218d3c7f425b)



