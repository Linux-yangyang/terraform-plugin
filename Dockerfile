FROM debian:stable-slim

ENV TZ Asia/Shanghai

RUN apt-get update && apt-get install -y --no-install-recommends \
                ca-certificates && apt install -y wget && apt-get install -y gpg && wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg && \
    gpg --no-default-keyring --keyring /usr/share/keyrings/hashicorp-archive-keyring.gpg --fingerprint && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com bullseye main" | tee /etc/apt/sources.list.d/hashicorp.list && \
    apt-get update && apt-get install -y terraform && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y \

ENV GO_VERSION 1.19

RUN apt-get update && apt-get install -y --no-install-recommends \
    g++ \
    gcc \
    libc6-dev \
    make \
    pkg-config \
    git \
    curl \
    && rm -rf /var/lib/apt/lists/*

RUN curl -SLO "https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz" \
    && tar -xzf "go$GO_VERSION.linux-amd64.tar.gz" -C /usr/local

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

COPY terraform-plugin /src/terraform
WORKDIR /src/terraform

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy && go build cmd/main.go

RUN mkdir -p /app/logs
RUN mkdir -p /app/conf

COPY --from=builder /src/terraform/main /app/bin/
COPY --from=builder /src/terraform/config.yaml /app/conf/config.yaml

WORKDIR /app

EXPOSE 4399

CMD ["./bin/main"]