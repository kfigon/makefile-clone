# Make build system in go

example:

<!-- todo - or maybe do it in yaml? -->
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


todo:
- [x] graph processing
- [x] cycle detection
- [x] topology sort
- [ ] parsing config