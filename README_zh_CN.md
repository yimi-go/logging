Logging
===
Try providing an easy-use logging facade api
一套易用的日志门面接口
实现一个易用的日志接口，并提供一个基于uber zap的默认实现。
# Features

# Install
TBD
# Usage
TBD
# Design Considerations
## Std lib logger?
* it's a struct.
* it too simple, and too crud. As my seen, the only useful methods are prints.
* it has dangerous panic and fatal methods, which used rarely.
* with a full functional logger interface, you wouldn't use the struct api, or an equivalent interface.
* do not support levels and structure output.
* 当将其抽象出一个 stdlogger 接口时，对外需要暴露两个接口：stdLogger、logger。
  * 如果两个接口不相关，适配层没有动力实现 stdLogger，因为他是为适配 std lib logger 设计的。
  * 所以只能把 stdLogger 作为 logger的内嵌接口。
  * 那么应该暴露哪个接口给使用方呢？getLogger 的函数/方法返回 stdLogger，那么使用方需要 assert 以获取 logger
  * 如果 getLogger 提供 logger ，那么 stdLogger 估计永远用不到。我已经能访问全功能的
## Styles
### why not glog v
* control logging level via logger names.
* long run application is a loop.
* dynamic reload levels
### why drop support of panic and fatal
* use error level logger and manual panic/exit.
* panic and fatal(exit) is dangerous, do these manually with caution.
### why not use style printing with level logger
* the code are too long. boring.
