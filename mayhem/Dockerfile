FROM golang:1.18 as builder

COPY . /yml2dot
WORKDIR /yml2dot/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzParser.a && \
    clang -fsanitize=fuzzer fuzzParser.a -o fuzzParser.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /yml2dot/mayhem/fuzzParser.libfuzzer /