# Make build system in go

example:

```
default: program

program.o: program.c
    gcc -c program.c -o program.o

program: program.o
    gcc program.o -o program
clean:
    rm -f program.o
    rm -f program
```

* targets `the_target:` - precedes `:`
    * target can have multiple steps
* dependencies - what a target depends on
* circular dependencies detection


todo:
- [x] graph processing
- [x] cycle detection
- [x] topology sort
- [ ] parsing config