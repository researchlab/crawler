# crawler

- support crawling user data from the given website.

```shell
seed --(requests)--------------> engine
                                 |
												         V
fetcher(fetch the page) <--- task queue(requests) for len(task queue) <=0 {quit}
  |                              ^
	V                              |
parser--(new requests)-----------^
```
