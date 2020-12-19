# Guideline for Panic/Error

-> When declaring our own error type, the variable needs to start with Err. It should also follow the camel case naming convention.
-> var ErrExampleNotAllowd= errors.New("error example text")

-> The error string should start with lowercase and not end with punctuation. One of the reasons for this guideline is that the error can be returned and concatenated with other information relevant to the error.

-> If a function or method returns an error, it should be evaluated. Errors not evaluated can cause the program to not function as expected.

-> When using panic(), pass an error type as the argument, instead of an empty value.

-> Do not evaluate the string value of an error.

-> Use the panic() function sparingly.