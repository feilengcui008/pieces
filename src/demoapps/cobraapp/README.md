demo for cobra pkg

+ `cobra init appName` to init a app with root cmd
+ `cobra add cmdName` to add a cmd to root cmd
+ each cmd can have persistent and local flags 
+ flags can have long name and short name(only one letter), flags can created from existing pkg level variables or string name
+ define global flags by adding persistent flags to root cmd
+ `cmd.AddCommand` to add sub cmd to existing cmd
