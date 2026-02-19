group "all" {
  targets = ["lqmonitor"]
}

target "lqmonitor" {
  dockerfile = "cmd/lqmonitor/Dockerfile"
  tags = ["ghcr.io/ashcycling/lqmonitor:v0.3.0"]
}