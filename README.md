# gomool

Toy to-golang compiler project.

## Gomool

### What does "gomool" mean?

Gomool means "junk" in Korean.

### What is "gomool"?

A ML-style toy programming language with sugar
which can be easily translated into golang.

### Why are you making this?

Go runtime is useful sometime because of GC and goroutine,
but I like Haskell or ML-style language rather than golang.

## Design Goal

What I want to make is moonscript for golang!

- ALMOST every golang code can be written in gomool with the same semantics.
- Single gomool file must be compiled into single golang file.
- Utilize indentations for block instead of `{}`.
- Remove English-word-keywords as much as possible.
- ML style language.
- Some sugars (such as defining operator, algebraic type, pattern matching, ...)
- (Wish) Better type system / checker / inferer.

## Objective

- [ ] Design language
- [ ] Add inferior editor support
- [ ] Write gomool compiler in golang
- [ ] Write gomool compiler in gomool
