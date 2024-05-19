# Docker Configuration Documentation

This documentation provides a detailed explanation of the `Dockerfile` and `docker-compose.yml` used to build and run a Go application with a PostgreSQL database.

## Dockerfile

The `Dockerfile` defines the steps to build a Docker image for the Go application.

### Builder Stage

```dockerfile
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main .
```

1. **Base Image**:
   - `FROM golang:1.22 as builder`: Uses the official Go 1.22 image as the base image for the build stage.
   
2. **Working Directory**:
   - `WORKDIR /app`: Sets the working directory inside the container to `/app`.
   
3. **Dependency Management**:
   - `COPY go.mod go.sum ./`: Copies the Go module files to the working directory.
   - `RUN go mod download`: Downloads the Go module dependencies based on the `go.mod` and `go.sum` files.
   
4. **Copy Source Code and Build**:
   - `COPY . .`: Copies the entire project directory to the working directory in the container.
   - `RUN CGO_ENABLED=0 GOOS=linux go build -v -o main`: Compiles the Go application with CGO disabled, targeting Linux, and produces an executable named `main`.

### Final Stage

```dockerfile
FROM alpine:latest
RUN apk --no-cache add ca-certificates
# Install Node.js and npm
RUN apk add --update nodejs npm
# Install javascript-obfuscator globally
RUN npm install -g javascript-obfuscator

WORKDIR /root/
COPY --from=builder /app/main .
# Ensure the entire public directory is copied into the Docker image
COPY --from=builder /app/public ./public
COPY --from=builder /app/scripts ./scripts
COPY --from=builder /app/.env .env
COPY --from=builder /app/cert ./cert

EXPOSE 80 443 587

CMD ["./main"]
```

1. **Base Image**:
   - `FROM alpine:latest`: Uses the lightweight Alpine Linux image as the base image for the final stage.
   
2. **Install Dependencies**:
   - `RUN apk --no-cache add ca-certificates`: Installs CA certificates to handle HTTPS.
   - `RUN apk add --update nodejs npm`: Installs Node.js and npm.
   - `RUN npm install -g javascript-obfuscator`: Installs the `javascript-obfuscator` globally.

3. **Working Directory**:
   - `WORKDIR /root/`: Sets the working directory inside the container to `/root/`.

4. **Copy Built Application and Other Files**:
   - `COPY --from=builder /app/main .`: Copies the built Go application from the builder stage.
   - `COPY --from=builder /app/public ./public`: Copies the `public` directory.
   - `COPY --from=builder /app/scripts ./scripts`: Copies the `scripts` directory.
   - `COPY --from=builder /app/.env .env`: Copies the `.env` file.
   - `COPY --from=builder /app/cert ./cert`: Copies the `cert` directory.

5. **Expose Ports**:
   - `EXPOSE 80 443 587`: Exposes ports 80 (HTTP), 443 (HTTPS), and 587 (SMTP) to the host.

6. **Command to Run the Application**:
   - `CMD ["./main"]`: Specifies the command to run the built Go application.

## docker-compose.yml

The `docker-compose.yml` file defines the services required to run the application, including the Go application and a PostgreSQL database.

```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
       - "80:80"   # Map port 80 on the host to port 80 in the container for HTTP
       - "443:443" # Map port 443 on the host to port 443 in the container for HTTPS
       - "587:587" # Map port 587 on the host to port 587 in the container for SMTP
    depends_on:
      - db
    environment:
      DB_HOST: db
      DB_PORT: 5432
      DB_USER: codequarry
      DB_PASSWORD: CQ1234
      DB_NAME: codequarrydb
      DB_TYPE: postgres
      URL: codequarry.dev
    volumes:
      - img_data:/root/public/img
  db:
    image: postgres:13
    ports:
      - "5432:5432"
    environment:
      POSTGRES_DB: codequarrydb
      POSTGRES_USER: codequarry
      POSTGRES_PASSWORD: CQ1234
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
  img_data:
```

### Services

1. **app**:
   - **build**: Builds the Docker image using the `Dockerfile` in the current directory.
   - **ports**: 
     - Maps port 80 on the host to port 80 in the container for HTTP.
     - Maps port 443 on the host to port 443 in the container for HTTPS.
     - Maps port 587 on the host to port 587 in the container for SMTP.
   - **depends_on**: Specifies that the `app` service depends on the `db` service, ensuring the database is started before the application.
   - **environment**: Sets environment variables for the application, including database connection details and the application URL.
   - **volumes**: 
     - Maps the `img_data` volume to the `/root/public/img` directory in the container, ensuring the `img` directory is persistent.

2. **db**:
   - **image**: Uses the official PostgreSQL 13 image.
   - **ports**: Maps port 5432 on the host to port 5432 in the container for PostgreSQL.
   - **environment**: Sets environment variables for the PostgreSQL database, including the database name, user, and password.
   - **volumes**: 
     - Maps the `postgres_data` volume to the `/var/lib/postgresql/data` directory in the container, ensuring the database data is persistent.

### Volumes

- **postgres_data**: Defines a named volume for persisting PostgreSQL data.
- **img_data**: Defines a named volume for persisting the `img` directory data.

### Summary

- **Dockerfile**: Defines a multi-stage build process for the Go application, ensuring a lightweight final image.
- **docker-compose.yml**: Defines the services and their configurations, including port mappings, dependencies, environment variables, and volumes for persistence. 

By following this setup, you ensure that your application and database are properly containerized, with persistent storage for critical data directories.