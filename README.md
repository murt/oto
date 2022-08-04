# oto
(O)ther (T)ool (O)rchestration - development pipeline tool for multi-language repositories.

`oto` is designed to work seamlessly with the existing tools of each target in your project.

### You should use `oto`...

__Monorepo with multiple different languages and toolchains__

> Tying together multiple tools is what `oto` is designed for

### You should _not_ use `oto`...

__Single target projects__

> There is most likely a better (default) toolchain for your project.

__Monorepos where all targets are written in the same language__
  
> There are often better tools specifically for a monorepo in your language (eg. turbo-repo for multiple Javascripty things)


## Goals

- Latest Starlark support
- Conventional tasks (install, build, lint, test, clean)
- File monitoring and caching
- Containerise
- Arbitrary tasks
- Linux, Mac, and Windows support
- Auto inform via known tools
- Monorepo focused experience

# Why?

I was working on a project that required python (via poetry), godot (via C#), C++ (via Scons), javascript, docker, and protobuf. I wanted to use the best tools for each job but also take care of dependencies and isolation for distribution in a language agnostic way (eg. I wanted to be able to compile protobuf into python *and* C# as a shared dependency of two different parts of the project).

I initially gravetated towards [Bazel](https://bazel.build/) but found that, while it supported a wide range of tools, it required more conformity and explicitness than I would like. For example, redeclaring dependencies that are specified in `requirements.txt` felt unnecessary.

## Supported Tools

### Python
- poetry
- pip
- pipenv
### Go
- go
### Rust
- cargo
### Docker
- docker
### Javascript
- npm
- yarn
### protobuf
### dotnet
- nuget
### godot
