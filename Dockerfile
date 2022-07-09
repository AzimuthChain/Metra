FROM golang:latest

ARG METRA_DIR=/bor
ENV METRA_DIR=$METRA_DIR

RUN apt-get update -y && apt-get upgrade -y \
    && apt install build-essential git -y \
    && mkdir -p /metra

WORKDIR ${METRA_DIR}
COPY . .
RUN make metra-all

ENV SHELL /bin/bash
EXPOSE 8545 8546 8547 30303 30303/udp

ENTRYPOINT ["metra"]
