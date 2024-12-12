# Usa la imagen oficial de Go 1.23.3
FROM golang:1.23.3

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia los archivos necesarios para la construcción
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto de los archivos de la aplicación
COPY . .

# Copia el archivo .env
COPY .env .env

# Expone el puerto 8080
EXPOSE 8080

# Compila la aplicación
RUN go build -o stori-transactions

# Ejecuta la aplicación al iniciar el contenedor
CMD ["./stori-transactions"]
