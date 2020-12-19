# There is a Go proverb that states "Accept interfaces, return structs." It can be restated as accept interfaces and return concrete types. This proverb is talking about accepting interfaces for your APIs (functions, methods, and so on) and the return to be structs or concrete types. This proverb follows Postel's Law, which states "Be conservative with what you do, be liberal with what you accept." We are focusing on the "be liberal with what you accept."

# By accepting interfaces, you are increasing the flexibility of the API for your function or method.

# Since an empty interface specifies no methods, that means that every type in Go implements an empty interface automatically. All types satisfy the empty interface.

