# Padrão de criação de funções

- Note: When the first letter of the function name is in lowercase, then the function is not exportable outside of a package. This means they are private and cannot be called from outside the package. They can only be called within the package.

- Keep this in mind when you use camelCase naming convention. If you want your function to be exportable, the first letter of the function name must be capitalized.

- Typically, when a function returns two types, the second type is of type error. We have not gone over errors yet so in these examples, we are not demonstrating them. It is good to know that it is idiomatic in Go for the second return type to be of type error.

 In Go, we can use what is called a blank identifier; it provides a way to ignore values in an assignment:

_, err := file.Read(bytes)

- For example, when reading a file, we might not be concerned about the number of bytes read. So, in that case, we can ignore the value being returned by using the blank identifier, "_". When there is extra data being returned from a function that does not provide any information that is needed by our program, such as the reading of a file, it is a good candidate for ignoring the return.

