# crawler

- support crawling user data from the given website.

single-crawler

```shell
seed --(requests)--------------> engine
                                 |
                                 V
fetcher(fetch the page) <--- task queue(requests) for len(task queue) <=0 {quit}
  |                              ^
  V                              |
parser--(new requests)-----------^
```

concurrent-crawler

```shell

engine <----(requests) <-------seeds
 |
 V
scheduler ---> requestqueue(request chan) --->activerequester
 |                ^                               |
 |                |--------worker(queue)<-----    | 
 |                                            ^   |
 V                                            |   V
 workerqueue(worker chan) --------------- activeworker

 ---(requests)-------------engine--------->data(items)
 |                           ^
 V                           |(requests, items)
 scheduler---(requests)-> worker(queue)
```

截图

![](resouces/pics/searchresult.png)
