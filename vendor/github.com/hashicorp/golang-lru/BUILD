load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "2q.go",
        "arc.go",
        "doc.go",
        "lru.go",
    ],
    importpath = "github.com/hashicorp/golang-lru",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/hashicorp/golang-lru/simplelru:go_default_library"],
)
