FROM qnib/go-webtest:v0.1.2 AS src

FROM qnib/uplain-cuda8:8.0.61-1.3
ENV WEBTEST_HTTP_PORT="9999"
COPY --from=src /webtest /webtest
CMD ["/webtest"]
