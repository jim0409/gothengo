# intro

golang 透過`DI(Dependency Injection)`代碼對於`interface`做單元測試


如果今天代碼是OOP，那麼依賴注入(depenency injection)就很容易實現

如果依賴注入被大量使用，那替換掉依賴將會變成一件輕而易舉的事情

把兩者結合，就可以得到一種編寫可測試代碼的模式

1. 將代碼的依賴抽象出來，抽象成一個接口，並且這個接口的實例不是自己創建出來，而是由上層調用方式注入進來
2. 將第三方依賴封裝成上面接口的一種實現，調用方負責創建具體的實例，並注入進業務代碼


# flow
根據依賴關係，在測試代碼裡。可以`mock`出另一種接口的實現

```
									  -- implement -- Real Dependency
									  |
MyFunction -- Depend -> <interface> --|
									  |
									  -- implement -- Mock Dependency
```


# refer:
- https://blog.betacat.io/post/2020/03/a-pattern-for-writing-testable-go-code/
