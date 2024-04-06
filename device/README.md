# device

Device Info

## usage 

```go
// 查看我的设备
device.List(false, 0) 

//询问小爱音响天气
device.Action(&ActionDetail{
    Did:  "{did}",
    Siid: 5,
    Aiid: 4,
    In:   []any{"今天天气", 1}, // 0 : silent-execution
})
```