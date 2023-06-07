FROM ubuntu:20.04

WORKDIR /root


RUN buildOeps='apt-utils ca-certificates curl unzip wget' \
	&& apt-get update \
	&& apt-get upgrade -y \
	&& apt-get install -y --no-install-recommends $buildDeps


ADD ./tmp/go1.19.5.linux-amd64.tar.gz /usr/local/
ADD ./tmp/jdk-11.0.1_linux-x64_bin.tar.gz /usr/local/java/

ENV JAVA_HOME /usr/local/java/jdk-11.0.1
ENV GO_ROOT /usr/local/go

ENV PATH $PATH:$JAVA_HOME/bin:$GO_ROOT/bin

RUN java -version \
&&	go version


