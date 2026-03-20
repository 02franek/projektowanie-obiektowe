FROM ubuntu:24.04

RUN apt-get update && apt-get install -y fpc

WORKDIR /app

COPY bubblesort.pas .

RUN fpc bubblesort.pas

CMD ["./bubblesort"]
