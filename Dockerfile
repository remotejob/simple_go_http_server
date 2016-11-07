FROM scratch
EXPOSE 8080

COPY server /
CMD ["/server"]
