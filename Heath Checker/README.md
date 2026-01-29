# Go Website Checker -- Sync vs Parallel

## Command Used

``` bash
go run main.go google.com amazon.com github.com example.com openai.com httpbin.org/status/200 httpbin.org/status/301 httpbin.org/status/400 httpbin.org/status/404 httpbin.org/status/500
```

## Synchronous Output

    time: 4.8709884s

Requests were executed one by one.

## Parallel Output

    time: 817.7524ms

Requests were executed concurrently using goroutines.

## Observations

-   Parallel execution is significantly faster.
-   Output order changes due to concurrency.
-   Some sites may return different status codes under load.
-   `403` from openai.com is expected due to access restrictions.

## Conclusion

Using goroutines and channels improves performance by running HTTP
checks in parallel.
