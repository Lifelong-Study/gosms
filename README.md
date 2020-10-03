<p align="center">
    <a href="https://github.com/Lifelong-Study/gosms">
        <label style="font-size: 30px;">gosms</label>
    </a>
</p>

<p align="center">
  <a href="https://travis-ci.com/GoAdminGroup/go-admin"><img alt="Go Report Card" src="https://api.travis-ci.com/GoAdminGroup/go-admin.svg?branch=master"></a>
  <a href="https://goreportcard.com/report/github.com/GoAdminGroup/go-admin"><img alt="Go Report Card" src="https://camo.githubusercontent.com/59eed852617e19c272a4a4764fd09c669957fe75/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f6368656e6867352f676f2d61646d696e"></a>
  <a href="https://goreportcard.com/report/github.com/GoAdminGroup/go-admin"><img alt="golang" src="https://img.shields.io/badge/awesome-golang-blue.svg"></a>
  <a href="https://t.me/joinchat/NlyH6Bch2QARZkArithKvg" rel="nofollow"><img alt="telegram" src="https://img.shields.io/badge/chat%20on-telegram-blue" style="max-width:100%;"></a>
  <a href="https://goadmin.slack.com"><img alt="slack" src="https://img.shields.io/badge/chat on-Slack-yellow.svg"></a>
  <a href="https://godoc.org/github.com/GoAdminGroup/go-admin" rel="nofollow"><img src="https://camo.githubusercontent.com/a9a286d43bdfff9fb41b88b25b35ea8edd2634fc/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f646572656b7061726b65722f64656c76653f7374617475732e737667" alt="GoDoc" data-canonical-src="https://godoc.org/github.com/derekparker/delve?status.svg" style="max-width:100%;"></a>
  <a href="https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/LICENSE" rel="nofollow"><img src="https://img.shields.io/badge/license-Apache2.0-blue.svg" alt="license" data-canonical-src="https://img.shields.io/badge/license-Apache2.0-blue.svg" style="max-width:100%;"></a>
</p> 

## Preface

gosms is a tool for fast delivery of Taiwanese SMS platforms


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
