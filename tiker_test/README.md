# for select with ticker cpu负载测试

## 测试步骤
基准测试，生成pprof文件
`go test -cpuprofile cpu.prof -bench=.
`

```shell
go test -cpuprofile TickDo.prof -bench  BenchmarkLongTimeTickDo
go test -cpuprofile TickDoWithoutDefault.prof -bench  BenchmarkLongTimeTickDoWithoutDefault
go test -cpuprofile TickDoWithSleep.prof -bench  BenchmarkLongTimeTickDoWithSleep
```

可视化查看pprof
`go tool pprof -http=:9091 <xxx.prof>`


## 测试结果
### 测试机器
```yaml
goos: windows
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
```

### cpu


## 相关资料
- https://github.com/golang/go/issues/27707
- https://github.com/golang/go/issues/25471
- https://loesspie.com/2020/08/05/golang-runtime-futex-high-cpu/


