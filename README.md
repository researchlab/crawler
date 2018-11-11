# crawler

- support crawling user data from the given website.

```shell
1.seed --(requests)--> engine ---> task queue(requests) ---> fetcher(fetch the page)--->parser(parse the page content)

2.parser--(new requests)--> task queue(requests)
```
