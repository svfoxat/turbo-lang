# Turbo Programming Language

- Inspiration and guidance from https://interpreterbook.com/



## Supported Types
* `boolean`
    * `true`, `false`
* `integer`
    * `1`, `100`
* `float`
    * `1.0`, `2.5f`
* `strings`
    * `"Hello"`, `"World"`  

## Operators
* `+, -, *, /, ==, =, !=, >=, <=`

## Conditionals
```
if (<expression>) {
        // do something
    } else {
        // do something
}
```

## Variable bindings
```
let x = 1;
let y = 2;
let hawi = 420 * x + y;
>> 422
```

## Functions
- Functions are implicitly returning, you can also use the `return` keyword


```
let double = fn(x) { 
    x * 2; 
}; 
double(5);

>> 10
```

- IIFEs are also supported
```
fn (x) {
    x * x
}(2)

>> 4
```

### Higher Order Functions

* return functions from functions
```
let newAdder = fn(x) { fn(y) { x + y }}
let addToTwo = newAdder(2)
addToTwo(3)

>> 5
```

* functions as arguments
```
let add = fn(a, b) { a + b }
let sub = fn(a, b) { a - b }
let applyFunc = fn(a, b, func) { fun(a,b) }

applyFunc(2, 2, add)
>> 4

applyFunc(2, 2, sub)
>> 0
```
