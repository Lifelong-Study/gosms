# gosms

## Preface

gosms 是個台灣簡訊平台快速發送的工具


## Supported Platforms

| 支援平台 |
| --- | --- | --- |
| 1. | 台灣簡訊 twsms | http://www.google.com/ |
| 2. | 簡訊購 smsgo | http://www.google.com/ |

## 使用範例
```go
package main

import (
    "fmt"
    "github.com/Lifelong-Study/gosms"
)

func main() {
    config := gosms.Config{
        TWSMS: gosms.Profile{"account 1", "password 1"}, 
        SMSGO: gosms.Profile{"account 2", "password 2"},  
    }
    
    // 方式一: 指定預設平台，未來發送都使用該平台發送 
    // sms := gosms.Init(config)
    // sms.SetDefaultPlatform(gosms.PLATFORM_TWSMS)
    // result, err := sms.SendText("0963265781", "test")

    // 方式二: 每次發送前都指定平台
    sms := gosms.Init(config)
    result, err := sms.UseTaiwanSMS().
        SendText("mobile", "message")

    if err != nil {
        fmt.Print(err)
        return
    }

    fmt.Print(result)
}
```
