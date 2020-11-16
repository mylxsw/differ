# differ

Differ 是一个命令行工具，用于记录文件内容变更， Differ 在每次执行之后，会在数据目录创建对比文件的当前状态副本，下次对比时，自动与最后一个副本进行比对。

```bash
Usage:
  -context-line int
    	输出上下文行数
  -data-dir string
    	状态数据存储目录 (default "/tmp")
  -file string
    	对比的文件名，留空则从标准输入读取
  -name string
    	当前对比项目名称 (default "default")
```
