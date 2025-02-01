# Pic-me-pls-backend

backend repository for Pic-me-pls Project

## Perequisite

- Golang ver go1.23.5 or newer https://go.dev/doc/install
- Docker https://docs.docker.com/desktop/setup/install/windows-install/
- air https://github.com/air-verse/air
- golangci-lint https://golangci-lint.run/welcome/install/

## Run local server

1. clone repository

```
git clone https://github.com/CP-RektMart/pic-me-pls-backend
cd pic-me-pls-backend
```

2. run docker compose

```
docker-compose up -d
```

3. start server
   there are 2 ways:

- normal

```
make start
```

- with hot reload

```
air
```
