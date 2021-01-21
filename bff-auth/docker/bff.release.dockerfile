FROM  rust:1.48.0 AS builder

RUN useradd -m app

USER app
ENV USER=app

WORKDIR /home/bff
RUN mkdir app

COPY ./ /home/bff/

WORKDIR /home/bff

RUN cargo build --release && \
    strip /home/bff/app/src/target/x86_64-unknown-linux-musl/release/app

FROM scratch

COPY --from=builder \
    /home/bff/app/src/target/x86_64-unknown-linux-musl/release/app /

EXPOSE 8080

ENTRYPOINT ["/app]