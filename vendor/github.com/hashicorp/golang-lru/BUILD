load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "2q.go",
        "arc.go",
        "lru.go",
    ],
    importpath = "github.com/hashicorp/golang-lru",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/hashicorp/golang-lru/simplelru:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "2q_test.go",
        "arc_test.go",
        "lru_test.go",
    ],
    importpath = "github.com/hashicorp/golang-lru",
    library = ":go_default_library",
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/github.com/hashicorp/golang-lru/simplelru:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
