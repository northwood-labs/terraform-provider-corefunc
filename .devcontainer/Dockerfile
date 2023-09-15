# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM mcr.microsoft.com/devcontainers/base:ubuntu AS base-builder

# Switch to proper Bash
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

# Informational
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM."

# Environment variables
ENV BASH_ENV "~/.bashrc"
ENV CGO_ENABLED 0
ENV CHECKPOINT_DISABLE true
ENV PATH /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin:/go/bin
ENV PIP_ROOT_USER_ACTION ignore
ENV TERRAGRUNT_SOURCE_UPDATE true
ENV TF_PLUGIN_CACHE_DIR "/root/.terraform.d/plugin-cache"
# ENV TZ "/usr/share/zoneinfo/Etc/UTC"

# Install the xz-utils package
RUN apt-get -y update
# RUN apt-get -y install --no-install-recommends

# bash \
# curl \
# exa \
# git \
# git-lfs \
# go \
# gofumpt \
# golangci-lint \
# gpg \
# grype \
# hadolint \
# infracost \
# jq \
# make \
# markdownlint-cli \
# nodejs \
# npm \
# openssh-client \
# python3 \
# python3-dev \
# shellcheck \
# shfmt \
# snyk-cli \
# terraform-docs \
# tfswitch \
# tgswitch \
# trivy \
# tz \
# tzdata \
# xz \
# yamllint \
# zip \

# coreutils \
# gcompat \
# git \
# ifupdown \
# libc6-compat \
# perl \
# poetry \
# procps \
# util-linux-misc \
# aws-cli
# syft

RUN python3 -m ensurepip
RUN python3 -m pip install --upgrade --no-cache-dir --no-cache --force pip
RUN python3 -m pip install --upgrade --no-cache-dir --no-cache --force wheel
RUN sleep 1
RUN python3 -m pip install --no-cache-dir --upgrade \
    pre-commit \
    ;

RUN mkdir -p /root/.ssh
RUN ssh-keyscan -H "github.com" >> /root/.ssh/known_hosts

# Cleanup
RUN rm -Rf /tmp/*
RUN find /tmp -type f -print0 | xargs -0 rm -Rf

# Code mounted here
WORKDIR /workspace

# Labels
ARG BUILD_DATE
ARG VCS_REF
ARG VCS_TAG
ARG VCS_URL

# https://github.com/opencontainers/image-spec/blob/master/annotations.md
LABEL org.opencontainers.image.title=""
LABEL org.opencontainers.image.version="$VCS_TAG"
LABEL org.opencontainers.image.description=""
LABEL org.opencontainers.image.url=""
LABEL org.opencontainers.image.documentation=""
LABEL org.opencontainers.image.vendor=""
LABEL org.opencontainers.image.authors=""
LABEL org.opencontainers.image.created="$BUILD_DATE"
LABEL org.opencontainers.image.source="$VCS_URL"
LABEL org.opencontainers.image.revision="$VCS_REF"
