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
$ cd sub
$ go run main.go
```

To access to main system,

```bash
$ for i in $(seq 5); do curl -x '' localhost:8080 &; done
```

Circuit Breaker will open after certain number of requests (`RequestVolumeThreshold`) and error rate is higher ahtn threshold (`ErrorPercentThreshold`). While opening, all requests are failed. After some certain time (`SleepWindow`), breaker try to send single test request to sub system, and request is succeeded, breaker will close. You can also set request Timeout and the number fo conccurent requests.

To try above function, let's stop sub system process :) and see what main system log says,

```bash
2015/05/08 14:15:10 / GET
2015/05/08 14:15:10 failed to get response from sub-system: hystrix: circuit open

2015/05/08 14:15:15 / GET
2015/05/08 14:15:15 hystrix-go: allowing single test to possibly close circuit my_command
2015/05/08 14:15:15 failed to get response from sub-system: Get http://localhost:9090: dial tcp 127.0.0.1:9090: connection refused

2015/05/08 14:15:33 / GET
2015/05/08 14:15:33 hystrix-go: allowing single test to possibly close circuit my_command
2015/05/08 14:15:34 hystrix-go: closing circuit my_command
2015/05/08 14:15:34 success to get response from sub-system: Hello
```

## References

- [Trying out hystrix-go](https://gist.github.com/marcusolsson/2ef585c2218d3c7f425b)



