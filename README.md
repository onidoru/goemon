## Goemon

Goemon is a C-like language interpreter, written in a Golang using top-down Pratt parser. It is dynamic weak-typed language.

### Start
    $ go get github.com/Onidoru/goemon
    $ cd $GOPATH/src/github.com/Onidoru/goemon/cmd

### Supported construction
Goemon supports variables, arithmetic operations,
if-else conditions, anonymous functions, arrays. Supported
data types: int, bool, fn, array.

    $ go run main.go
    >> let expr = fn(a,b) { return a * b + (a + 3 * (b / 2)) - 2; };
    >> let a = 3;
    >> let b = 6;
    >> a = 4;
    >> if (!false) {expr(a,b);}
       55
    >> let arr = [2, 2*3, expr(a, b)];
    >> expr(arr[0], arr[2]);
       299

It is a dynamic and weak-typed language, so functions will accept every parsed value, but errors will rise on evaluation.

    >> if;
        expected next token to be (, found ; instead
    >> o;
    ERROR: identifier not found: o
    >> 2 + true;
    ERROR: type mismatch: INTEGER + BOOLEAN
    >> expr(true, arr[17]);
    ERROR: type mismatch: BOOLEAN * NULL

 
