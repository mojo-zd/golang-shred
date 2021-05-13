代码开闭原则,对扩展开放对修改关闭

- Option模式  options.go
- 工厂模式  factory.go
1. 简单工厂(静态工厂),扩展性不好
2. 工厂方法
3. 抽象工厂

- 适配器模式 adapter/
1. Target对象
2. Adapter 用于连接Target和Adaptee的业务
3. Adaptee 被适配对象

- 单例模式  single/
1. 懒汉模式
2. 饿汉模式
3. 枚举类型
3. 双重检查 
- 原型模式  prototype/
- 代理模式  proxy/ proxy.go
1. 客户调用抽象类
2. 抽象类调用代理类
3. 抽象类调用真实角色

- 观察者模式 observer/
- Command模式 command.go
- 策略模式 strategy.go
- 装饰器模式 decorator.go
- visitor模式 visitor/
- 模板方法模式  templatemethod.go
可以在父类实现某些具体业务, 同时不确定业务定义为接口. 提供统一模板方法执行具体流程
- 状态模式 
用于解决系统中复杂对象的状态切换以及不同状态下行为的封装问题
- 结构
1. Context环境类, 环境中维护一个State对象,它定义了当前状态
2. State抽象状态类
3. ConcreateState具体状态类实现,每一个类封装了一个状态的行为
- 备忘录模式 memento/
恢复到某一个时间点的数据。e.g: 历史记录、撤销
1. 源发器
2. 备忘录对象保存源发器某个时间点数据
3. 负责人存储备忘对象,需要恢复时从负责人获取
- 建造者模式
针对一个大型对象拆分成小对象进行构建
- 桥接模式
- 组合模式
适用于树形结构,可以统一处理部分对象和整体对象.包括叶子节点和容器节点,
容器节点多了add、remove、get等方法
- 装饰器模式
动态的为一个对象增加新功能,用于代替继承。包含组件:
1. Component 抽象构建角色
2. ConcreteComponent 具体构建角色
3. Decorator 装饰角色,持有一个抽象构建的引用