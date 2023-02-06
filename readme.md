<p align='center'>
  <img src='https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/9439deee6fda4cbf869c80488939931c~tplv-k3u1fbpfcp-watermark.image?' alt='CLI ChatBot Power By Gpt3' width='800'/>
</p>

<p align='center'>
  <img src='https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/ee1f8cf9cd6e4b989f51b15ee5da5afb~tplv-k3u1fbpfcp-watermark.image' alt='CLI ChatBot Power By Gpt3' width='800'/>
</p>

<p align='center'>
    CLI ChatBot Power By <b>Gpt3</b><sup><em>(speed)</em></sup><br>
</p>

## 源代码

🏀🏀
[CLI ChatBot Power By Gin](https://github.com/Leizhenpeng/gpt3-chatBot-cli)
🏀🏀

## bilibili 在线吹水

⚽️⚽️
[基于gpt3引擎实现CLI版本的chatgpt~](https://www.bilibili.com/video/BV1ae4y1N7W2)
⚽️⚽️

## Features

- 🐤 Gpt3 ChatBot
- 🍉 Store Keyring
- 🥑 Cobra CLI

## development

```bash
make build
```

```bash
go-chat key set <your-key>
go-chat run
```

## Snapshot

- 记录上下文，实现连续对话

<p align='center'>
  <img src='https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/bbda7ad297204aaebe8f47da89768876~tplv-k3u1fbpfcp-watermark.image' alt='CLI ChatBot Power By Gpt3' width='400'/>
</p>

``` go
func NewCacheHistory() CacheHistoryInterface {
   return &CacheHistory{
      db: cache.New(time.Second*60, time.Minute*10),
   }
}

func (c CacheHistory) SetQACache(q string, a string) {
   saveValue := q + "\n" + a
   c.set(userSessionContextKey, saveValue)
}

func (c CacheHistory) GetQACache() (string, bool) {
   res, ok := c.get(userSessionContextKey)
   if !ok {
      return "", false
   }
   return res.(string), true
}

func (c CacheHistory) ClearQACache() {
   c.clear()
}
```

- 封装调用gpt3接口

``` go
func GetAnswer(question string) {
  
   fmt.Print("Answer: ")
   i := 0
   ctx := context.Background()
 
   if err := client.CompletionStreamWithEngine(ctx, engine, gpt3.CompletionRequest{
      Prompt: []string{
         question,
      },
      MaxTokens:   gpt3.IntPtr(maxTokens),
      Temperature: gpt3.Float32Ptr(temperature),
   }, func(resp *gpt3.CompletionResponse) {
      if i > 1 {
         fmt.Print(resp.Choices[0].Text)
      }
      i++
   }); err != nil {
      log.Fatalln(err)
   }
   fmt.Println()
}
```

- cobra项目结构

```
├── LICENSE
├── Makefile
├── cmd
│   ├── key.go
│   ├── root.go
│   └── run.go
├── go-chat
├── go.mod
├── go.sum
├── main.go
├── readme.md
└── services
    ├── browse.go
    ├── cli.go
    ├── gpt.go
    └── key.go
    
```

## More Info

- [Go Keyring library](https://github.com/zalando/go-keyring)

- [Cobra CLI library](https://github.com/spf13/cobra-cli/blob/main/README.md)

- [golang常见库cobra](https://segmentfault.com/a/1190000023382214)

- [chatGpt3](https://github.com/PullRequestInc/go-gpt3)
