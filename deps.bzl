load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "in_gopkg_hraban_opus_v2",
        importpath = "gopkg.in/hraban/opus.v2",
        sum = "h1:sxrRNhZ+cNxxLwPw/vV8gNsz+bbqRQiZHBYBJfpyNoQ=",
        version = "v2.0.0-20201025103112-d779bb1cc5a2",
    )
