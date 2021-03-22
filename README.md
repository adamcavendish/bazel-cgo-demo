# Building Golang CGO with Bazel

Bazel is a fast building system open sourced by Google. It has go  support but it does not have a detailed instruction on how to integrate  CGO with the bazel build system. Building everything from scratch with `cc_library` and manually written `go_library` is do-able. However, we do prefer the amazing [gazelle](https://github.com/bazelbuild/bazel-gazelle) and sometimes it is hard to use prebuilt libraries with CGO build without digging up in the source code in the `gazelle`.



The post is available at https://blog.modest-destiny.com/posts/building-golang-cog-with-bazel/



The major idea is to use `patch` to patch the `go_repository()` to use libraries written in `cc_import` and `cc_library`.

# Source code

For each steps in the post, I've created a corresponding branch for testing.

| Step                                                         | Git command         | URL Link                                                     |
| ------------------------------------------------------------ | ------------------- | ------------------------------------------------------------ |
| Step 1: Initialize the repository                            | git checkout step-1 | [Step 1](https://github.com/adamcavendish/bazel-cgo-demo/tree/step-1) |
| Step 2: Declare the System libopus Pre-built Dependency      | git checkout step-2 | [Step 2](https://github.com/adamcavendish/bazel-cgo-demo/tree/step-2) |
| Step 3: Patch the `go_repository()` Rule for `in_gopkg_hraban_opus_v2` to Depend on libopus | git checkout step-3 | [Step 3](https://github.com/adamcavendish/bazel-cgo-demo/tree/step-3) |
| Step 4: Add ogg Libraries for Linking                        | git checkout step-4 | [Step 4](https://github.com/adamcavendish/bazel-cgo-demo/tree/step-4) |

