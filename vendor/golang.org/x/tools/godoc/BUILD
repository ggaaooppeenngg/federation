load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cmdline.go",
        "corpus.go",
        "dirtrees.go",
        "format.go",
        "godoc.go",
        "index.go",
        "linkify.go",
        "meta.go",
        "page.go",
        "parser.go",
        "pres.go",
        "search.go",
        "server.go",
        "snippet.go",
        "spec.go",
        "spot.go",
        "tab.go",
        "template.go",
    ],
    importpath = "golang.org/x/tools/godoc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/tools/godoc/analysis:go_default_library",
        "//vendor/golang.org/x/tools/godoc/util:go_default_library",
        "//vendor/golang.org/x/tools/godoc/vfs:go_default_library",
        "//vendor/golang.org/x/tools/godoc/vfs/httpfs:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "cmdline_test.go",
        "godoc17_test.go",
        "godoc_test.go",
        "index_test.go",
    ],
    importpath = "golang.org/x/tools/godoc",
    library = ":go_default_library",
    deps = [
        "//vendor/golang.org/x/tools/godoc/vfs:go_default_library",
        "//vendor/golang.org/x/tools/godoc/vfs/mapfs:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/golang.org/x/tools/godoc/analysis:all-srcs",
        "//vendor/golang.org/x/tools/godoc/redirect:all-srcs",
        "//vendor/golang.org/x/tools/godoc/static:all-srcs",
        "//vendor/golang.org/x/tools/godoc/util:all-srcs",
        "//vendor/golang.org/x/tools/godoc/vfs:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
