# Imagem base para o contêiner
FROM golang:latest

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o código-fonte para o contêiner
COPY . .

# Compila e executa a lambda operations
WORKDIR /app/cmd/operations
RUN go build -o main

# Compila e executa a lambda redirectionUrl
WORKDIR /app/cmd/redirectionUrl
RUN go build -o main

# Compila e executa a lambda urls
WORKDIR /app/cmd/urls
RUN go build -o main

# Define o comando padrão para o contêiner
ENTRYPOINT ["/bin/ls"]
CMD ["-l"]