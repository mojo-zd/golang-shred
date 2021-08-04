### Option模式思路

1. 构建一个统一的`Options`收纳`OptionCreator`对象并统一遍历处理

```
type Opitons struct {
  builders []app.OptionCreator
}

NewOptions(options ...app.OptionCreator) Options{
    return &Options{
        builders: options[:]
    }
}

func (o *Options) AddFlags(flag *pflag.FlagSet) {
    for _, builder := range o.builders {
        builder.AddFlag(flag)
    }
}

func (o *Options) Initializer(cfg interface) {
    var errors []error
	for _, builder := range o.builders {
		if errs := builder.Validate(); len(errs) > 0 {
			errors = append(errors, errs...)
		}
	}
	if len(errors) > 0 {
		return MultiError(errors).AsError()
	}
	
    for _, builder := range o.builders {
        builder.ApplyToConfig(cfg)
    }
}
```

2. 抽象各个`Option`的统一方法入口

```
type OptionCreator interface {
    AddFlags(flag *pflag.FlagSet)
    Valdate() []error
    ApplyToConfig(cfg interface)
}
```

3. 处理具体`Option`对象

```
type CommonConfigInitializer interface {
    InitCommonConfig(name, component string)
}

type commonOption struct {
    name string
    component string
}

NewCommonOption() app.OptionCreator {
    return &commonOption{}
}

func (c *commonOption) AddFlags(flag *pflag.FlagSet) {
 // TODO bind flag to commonOption's atrr
}

func (c *commonOption) Validate() []error {
  // valdate params 
}

func (c *commonOption) ApplyToConfig(cfg interface) {
    if initializer, ok := cfg.(CommonConfigInitializer); ok {
        initializer.InitCommonConfig(c.name, c.component)
    }
}
```

4. 最终需要的对象

```
type CommonConfig struct {
   Name string
   Component string
}

func (c *CommonConfig) InitCommonConfig(name, component string) {
    c.Name = name
    c.Component = component
}
```