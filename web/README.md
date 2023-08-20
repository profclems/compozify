# Compozify

Compozify is a simple (yet complicated) tool to generate a `docker-compose.yml` file from a `docker run` command.

## Getting Started

To start the development server, run the following command:

Clone the repository:

```bash
git clone https://github.com/profclems/compozify.git
```

Navigate to the directory:

```bash
cd compozify/web
```

### DEVELOPMENT

#### Install dependencies

```bash
pnpm install
```

#### Start the backend server

##### Prerequisites

- [Go 1.18+](https://golang.org/dl/)

##### Run the server

Open a new terminal and route to `cmd/compozify-web` and run:

```bash
go run main.go
```

#### Start the development server

In `compozify/web/next.config.js` file

1. Enable comment the `async rewrites()` function.

2. Disable comment the `output: 'export'` function.

Then run the following command:

```bash
pnpm dev
```

**NB**:

- Make sure you have the backend server running in another terminal in order to make requests to the API endpoints.

- Return the `async rewrites()` function and the `output: 'export'` function to its initial state before building the app for production or making commits.
