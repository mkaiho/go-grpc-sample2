FROM debian:bullseye

ARG GO_VERSION=1.19.1
ARG NODE_VERSION=16.15.0
ARG PB_VERSION=21.6
ARG SRC_MODULE=github.com/mkaiho/go-grpc-sample2

RUN apt-get update -y && apt-get install -y \
  curl gcc make git zip unzip jq openssl ca-certificates gnupg lsb-release clang-format

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
  && unzip awscliv2.zip \
  && rm awscliv2.zip \
  && ./aws/install \
  && rm -r ./aws \
&& curl -L -o direnv https://github.com/direnv/direnv/releases/download/v2.29.0/direnv.linux-amd64 \
  && chmod +x direnv \
  && mv direnv /usr/local/bin/ \
  && echo 'eval "$(direnv hook bash)"' >> ~/.bashrc
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg \
&& echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null \
&& apt-get update -y && apt-get -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin

RUN mkdir -p /go/src/$SRC_MODULE && chmod -R 777 /go
RUN chsh -s /bin/bash

WORKDIR /go/src/$SRC_MODULE

RUN git clone https://github.com/anyenv/anyenv ~/.anyenv \
&& echo 'export PATH="$HOME/.anyenv/bin:$PATH"' >> ~/.bash_profile \
&& echo 'eval "$(anyenv init -)"' >> ~/.bash_profile \
&& . ~/.bash_profile \
&& echo y | anyenv install --init \
&& anyenv install nodenv \
&& anyenv install goenv \
&& . ~/.bash_profile \
&& nodenv install $NODE_VERSION \
&& nodenv global $NODE_VERSION \
&& npm install -g aws-cdk aws-cdk-local \
&& goenv install $GO_VERSION \
&& goenv global $GO_VERSION \
&& echo 'alias aws-local="aws --endpoint http://localstack:4566"' >> ~/.bashrc
RUN curl -LO "https://github.com/protocolbuffers/protobuf/releases/download/v${PB_VERSION}/protoc-${PB_VERSION}-linux-x86_64.zip" \
&& unzip protoc-${PB_VERSION}-linux-x86_64.zip -d ${HOME}/.local \
&& echo 'export PATH="$PATH:${HOME}/.local/bin"' >> ~/.bash_profile \
&& echo 'export PATH="$PATH::$(go env GOPATH)/bin"' >> ~/.bash_profile \
&& . ~/.bash_profile \
&& go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
