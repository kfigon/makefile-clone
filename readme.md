# Make build system in go

example:

```
HEADERS = program.h headers.h

default: program

program.o: program.c $(HEADERS)
    gcc -c program.c -o program.o

program: program.o
    gcc program.o -o program

clean:
    -rm -f program.o
    -rm -f program
```

* variables `HEADERS=...`
* targets `the_target:` - precedes `:`
    * target can have multiple steps
* dependencies - what a target depends on
* circular dependencies detection