Logging
===
尝试提供一套易用的日志门面API接口。
# 设计特性
* 可用易用
* 非vendor绑定。仅定义接口和数据struct，内置 Nop （什么也不输出）。
* 方便适配。提供一个可用参考实现：[zap-logging](https://github.com/yimi-go/zap-logging)
# 安装
TBD
# 使用方法
TBD
# 适配方法
TBD
# 设计考量
## kratos，or not kratos
整个 yimi-go 项目是为了实现一个方便部署到 service-mesh 环境的微服务框架。
在设计上大量的借鉴其他开源 kit 框架。

kratos 是优秀的框架项目，其代码实现是很有参考价值。其实现基本处处体现着简洁。
但在日志部分，我觉得其在设计上有些过犹不及，过于简洁，只提供了一个简单的接口供适配：
```go
// Logger is a logger interface.
type Logger interface {
	Log(level Level, keyvals ...interface{}) error
}
```
同时提供一个 Helper 来提供方便使用的 Info、Error 等。

这种设计有以下几个问题：
* 同时暴露两个接口，调用方可能使用任意一个。
  咋一看，这不会是什么问题，但到需要日志携带 caller 信息时，问题就来了。
  vendor 适配时难以确定需要 skip 多少 stack 层级来计算 caller 和 stacktrace。
  为了解决这个问题，kratos 贴心的提供了 `Caller` 的 `Valuer` 构造器函数，
  方便调用者提供 caller 信息。但问题并没完美解决，还带来了以下问题：
    * 是否需要计算 caller 无法在运行时通过配置控制，因为已经在代码里写死了。
    * 开发者必须关心 vendor，vendor 锁定。
      开发者必须在开发阶段确定使用哪个 vendor 的适配，才能准确计算 caller depth，
      且必须保证之后不能替换 vendor。
* Log 接口的 Log 方法使用了变长参数 keyvals 。这个参数提供要输出的日志结构键值对。
  如果提供的键值对不成对，可能产生错位，需要人工检查（可靠性?）或打warn日志(为时晚矣？)
  或panic(Oh no!)。
* 项目其他功能模块直接使用 Helper，而 Helper 是 struct ，无法变更实现，
  对测试 mock 也不友好。经典的设计模式、最佳工程实践都告诉我们要基于接口而非实现编程，
  Kratos 项目基于 Helper 使用日志框架也算是个反模式了。

Kratos 提供了一系列工程的最佳实践，尤其是在服务性可用性方面，log 模块成为了整个项目的一大败笔。

> 当然，以上一些问题也可能不是问题，毕竟代码是要人来写的，通过行政手段管好人不犯错就可以了😅。
## 标准库 logger?
项目一开始计划提供标准库 log 的接口抽象，毕竟标准库提供的 logger 是个 struct，
存在难以 mock 等问题。但经过考虑还是放弃了对标准库 logger 的支持：
* 标准库的功能实在过于简单，或者说简陋。
  以本人愚见，只有 prints 方法在比较简单的场景下有些可用价值。
  标准库logger 不支持结构化输出，要输出结构化，需要大量代码组装结构化数据再序列化，使用复杂。
* 其 caller 计算是在内部实现的。外部完全无法介入。
  一旦需要将其封装到工具函数/方法里，就需要类似 kratos caller valuer 的方案了。
* 标准库logger 提供了危险的 panic 和更危险的 fatal 类方法。没有特殊的理由，业务不应主动 panic。
  即使要 panic 也必须妥善的 recover 处理，否则导致进程不断退出重启，就是 bug、事故了。
  fatal 类方法更是没有 recover 机会。
  通常只有启动 init 阶段检查启动条件不满足时才会 panic，且此时需要返回特定进程退出码时才会使用 fatal/exit。
  相比将 panic/fatal 能力提供到 logger API 里，建议真正需要panic/fatal的时候使用 error 级别日志 +
  手动 panic / os.Exit。
* 如果日志框架提供了更全功能的日志接口时，谁还会用提供的标准化库 logger 接口呢？🤔
    * 当将其抽象出一个 stdlogger 接口时，对外需要暴露两个接口：stdLogger、logger。
        * 如果两个接口不相关，适配层没有动力实现 stdLogger。
          标准库接口功能比较简单，使用这个接口开发都的基本都是确定自己的日志需求已经可以被这个API满足了。
          那么 vendor 实现这个接口时如果功能比标准库有加强就会打破开发者的预期；
          如果不加强，为什么不直接使用标准库的 logger struct？
        * 所以只能把 stdLogger 作为 logger的内嵌接口。 那么应该暴露哪个接口给使用方呢？
            * 如果框架的 getLogger 的函数/方法返回 stdLogger，那么使用方需要 assert 以获取全功能的 Logger。
              一方面如前面所述，如果调用方是要用标准库相同的 API，何必使用这个框架？
              另一方面，如果调用方是为了使用全功能的 Logger ，那代码里大量的 assert 会很麻烦、很蠢。
            * 如果 getLogger 提供 logger ，那么 stdLogger 估计永远用不到。
              标准库的 Logger 也一定实现不了全功能 Logger 接口。
## 不支持 glog 系的 V 方法
glog 支持 按 module 控制日志输出和 V 方法分级输出。但是：
* 按 module 控制只能按 package 最后路径部分控制，控制精度有问题。
* 按V方法分级输出会有过度输出的问题。对于长时间运行的服务端应用来说，日志量放大可能导致服务可用性问题。

而且这两种方法基本都是以启动参数的方式来设置，不适用动态调整。
yimi-go 最终目标是要支持 service-mash 的云原生应用，应该是日志配置文件化、版本化。
通过文件系统监听可以实现配置的动态重载。
通过按 logger name 配置日志等级，可以解决 glog 按文件目录 module 控制精度问题。
## 不支持 panic 和 fatal 类日志方法
在不支持标准库 logger 部分已有说明，这里不再赘述。
## 其他风格问题考虑
* 之前有读到过一篇博文，其表述了一个观点：日志应取消 Warn 级别，因为不是Error所以没有人会看，并导致日志级别错乱。
  初读之时大为赞同，但细想，用错日志等级还是程序员的问题。取消 Warn 级别还是不现实。
  我可以不用，但你不能没有啊！而且 Warn 日志也远没有 panic / fatal 那样的"负"作用。
  另外，使用日志收集系统和 metric 系统，是可以基于日志等级做报警系统的。
* 有些日志框架API提供了获取 Leveled Logger 再 print 的方式调用。
  实现框架时也考虑过这样，因为这样可以提供比较简洁的API界面。但实现使用上代码会冗长不少，
  当前日志接口的设计是要易用易接入，与设计方向相背，故放弃。
