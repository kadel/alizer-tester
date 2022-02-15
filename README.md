# alizer-tester

Simple tool for quick testing of  [Alizer](https://github.com/redhat-developer/alizer) Go implementation.



example:
```
$ go build -o alizer main.go

$ ./alizer -path ./
# Runing recognizer.Analyzer("./")
{Name:Shell Aliases:[sh shell-script bash zsh] UsageInPercentage:25 Frameworks:[] Tools:[] CanBeComponent:false}
{Name:Modula-2 Aliases:[] UsageInPercentage:25 Frameworks:[] Tools:[] CanBeComponent:false}
{Name:AMPL Aliases:[] UsageInPercentage:25 Frameworks:[] Tools:[] CanBeComponent:false}
{Name:Go Aliases:[golang] UsageInPercentage:25 Frameworks:[] Tools:[1.17] CanBeComponent:true}
{Name:GCC Machine Description Aliases:[] UsageInPercentage:25 Frameworks:[] Tools:[] CanBeComponent:false}

# Runing recognizer.SelectDevFileFromTypes("./")
{Name:go Language:go ProjectType:go Tags:[Go]}

# Runing recognizer.DetectComponents("./")


```