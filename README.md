# project-GROS
<div align="center">
  <picture>
    <img alt="GROS - Making Rust better with Go" src="img/gros.png" height="300">
  </picture>
  <p>Making Rust better with Go</p>
</div>

## Raison d'être

Microsoft recently [announced](https://devblogs.microsoft.com/typescript/typescript-native-port/) the porting of the existing Typescript compiler from TypeScript/JavaScript to the native language, Go. Doing so is reducing the compile times by an order of magnitude.

One key takeaway from the FAQ item titled "Why Go?" in the project's Github repo:
Go was chosen for performance reasons. The team evaluated other languages like Rust, C and C++, but found Go to be the best fit for their needs. Go is designed to be memory safe. It provides features like garbage collection, strong typing, and built-in memory safety mechanisms to prevent common memory-related issues such as buffer overflows, null pointer de-references, and data races.

This has raised the question of what other programming languages could benefit from a similar treatment. 

Rust is gaining traction in various industries, including systems programming and cloud-native development. Compile times in Rust, while improving, are still painfully slow. 

Go has been optimized for blazing fast compilation time since the very start. 

Rust implements memory safety with its borrow-checker, however so much Rust code makes use of unsafe that it begs the question : Is what Rust really needs a *Garbage Collector*?

The objective of this project is to reimplement Rust in Go, improving Rust's deficiencies with Go's strengths. As this is a merging of Go and Rust, we've decided to also merge names: the result - Project "GROS" aka "GROSlang".

## Implementation status

We have started with the lexer and parser, compiler design is ongoing and we are creating placeholders with TODOs as we progress.

## How to Contribute

Please read our latest [project contribution guide](https://github.com/UoCCS/project-GROS/blob/main/CONTRIBUTING.md).

## Getting involved

Even if you do not plan to contribute to GROSlang itself, we'd be happy to have you involved:

- Follow our activity on [GitHub issues](https://github.com/UoCCS/project-GROS/issues)
- Share the repo on social media
- Contribute code

## Acknowledgements

Many thanks to the _Union of Concerned Computer Scientists - Memory Safety working group_ for your contributions and support.
