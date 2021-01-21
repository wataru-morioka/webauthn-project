FROM  rust:1.48.0

RUN useradd -m bff && \
    cargo install cargo-watch && \
    cargo install sccache

# USER bff
ENV USER=app
ENV RUSTC_WRAPPER=/usr/local/cargo/bin/sccache
ENV CARGO_BUILD_TARGET_DIR=/tmp/target

COPY ./.aws /root/.aws

WORKDIR /root
RUN mkdir app
WORKDIR /root/app

EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]