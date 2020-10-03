# gosms

台灣簡訊工具

支援平台: 
    台灣簡訊 twsms
    簡訊購 smsgo

```go
package main

import (
    "fmt"
    "github.com/Lifelong-Study/gosms"
)

func main() {
    config := gosms.Config{
        TWSMS: gosms.Profile{"simonly0001",  "0be8939c"}, 
        SMSGO: gosms.Profile{"3","4"},  
    }
    
    // 方式一: 指定預設平台，未來發送都使用該平台發送 
    //sms := gosms.Init(config)
    //sms.SetDefaultPlatform(gosms.PLATFORM_TWSMS)
    //result, err := sms.SendText("0963265781", "test")

    // 方式二: 每次發送前都指定平台
    sms := gosms.Init(config)
    result, err := sms.UseTaiwanSMS().
        SendText("0963265781", "998")

    if err != nil {
        fmt.Print(err)
        return
    }

    fmt.Print(result)
}
```

| First Header  | Second Header |
| ------------- | ------------- |
| Content Cell  | Content Cell  |
| Content Cell  | Content Cell  |