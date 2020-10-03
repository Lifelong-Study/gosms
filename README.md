## Preface

gosms is a tool for fast delivery of Taiwanese SMS platforms.

## Supported Platforms

| Index | Platform | Official Website |
| --- | --- | --- |
| 1. | 台灣簡訊 | https://www.twsms.com/ |
| 2. | 簡訊購 | https://www.smsgo.com.tw/ |

## How To Use

```go
package main

import (
    "fmt"
    "github.com/Lifelong-Study/gosms"
)

func main() {

    // All platforms are optional, no need to add if not used 
    config := gosms.Config{
        TWSMS: gosms.Profile{"account 1", "password 1"}, 
        SMSGO: gosms.Profile{"account 2", "password 2"},
    }
    
    // Method 1: 
    //     Specify a preset platform, 
    //     and all future sending will use this platform to send
    // sms := gosms.Init(config)
    // sms.SetDefaultPlatform(gosms.PLATFORM_TWSMS)
    // result, err := sms.SendText("0963265781", "test")

    // Method 2: Specify the platform before each sending
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
