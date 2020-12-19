# In Go, package names should be all lowercase with no underscores. Don't use camel case or snake case styling. There are multiple packages with pluralized names.

# Avoid package names such as misc, util, common, or data. These package names make it harder for the user of your package to understand its purpose.

# If a function, type, variable, and so on starts with an uppercase letter, it is exportable; if it starts with a lowercase letter, it is unexportable. There are no access modifiers to be concerned with in Go. If the function name is capitalized, then it is exported, and if it is lowercase, then it is unexported.

# Note: It is good practice to only expose code that we want other packages to see. We should hide everything else that is not needed by external packages.

# Alias package

# The package aliasing syntax is very simple. You place the alias name before the import package path: import  f "fmt"

# There can be more than one init() function in a package. This enables you to modularize your initialization for better code maintenance. 

# For example, suppose you need to set up various files and database connections and repair the state of the environment that your program will be executed in. Doing all that in one init() function would make it complicated for maintaining and debugging. 