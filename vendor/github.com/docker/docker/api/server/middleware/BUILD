load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cors.go",
        "debug.go",
        "experimental.go",
        "middleware.go",
        "version.go",
    ],
    importpath = "github.com/docker/docker/api/server/middleware",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/errors:go_default_library",
        "//vendor/github.com/docker/docker/api/server/httputils:go_default_library",
        "//vendor/github.com/docker/docker/api/types/versions:go_default_library",
        "//vendor/github.com/docker/docker/pkg/ioutils:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "debug_test.go",
        "version_test.go",
    ],
    importpath = "github.com/docker/docker/api/server/middleware",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/api/server/httputils:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
