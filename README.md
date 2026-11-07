# Acetlisto

A Go-based shopping list REST API developed as part of the DevOps & SecOps course at Hochschule Heilbronn.

This team project explored modern DevSecOps practices, including automated testing, security scanning, containerization, cloud deployment, and application observability.

This repository is a personal portfolio copy maintained to demonstrate the project, its technologies, and local development and CI practices. The original work was developed collaboratively as part of the university team.

---

## Project Overview

Acetlisto is a REST API for managing shopping list items, developed using Go.

The team project focused on combining application development with a modern software delivery workflow.

Key areas included:

- REST API development using Go
- CI/CD automation with GitHub Actions
- Containerization using Docker
- Automated testing and security scanning
- Cloud deployment using Google Cloud
- Application observability with OpenTelemetry

---

## Technologies

| Area | Technologies |
|---|---|
| Application | Go, REST API |
| CI/CD | GitHub Actions |
| Containerization | Docker |
| Security | GoSec, Govulncheck, Snyk |
| Cloud | Google Cloud Run, Google Artifact Registry |
| Authentication | Workload Identity Federation |
| Observability | OpenTelemetry, Google Cloud Trace |

*These technologies reflect the broader university team project. Not every cloud integration or service is configured in this personal repository.*

---

## DevSecOps Highlights

The team project explored:

- Automated CI/CD workflows using GitHub Actions
- Automated testing and test coverage
- Static security analysis with GoSec
- Vulnerability detection with Govulncheck
- Application security scanning with Snyk
- Docker-based application containerization
- Container registry integration
- Cloud deployment workflows
- Distributed tracing with OpenTelemetry

---

## Personal Repository CI Pipeline

This repository includes a GitHub Actions workflow for validating the application independently of the university team's deployment infrastructure.

The workflow performs:

1. Repository checkout
2. Go environment setup and dependency download
3. Go formatting checks
4. Static analysis using `go vet`
5. Automated tests with race detection and coverage
6. Security scanning using GoSec
7. Vulnerability checking using Govulncheck
8. Docker image build

The personal workflow builds and validates the application. It does not publish container images or deploy to Google Cloud.

---

## Cloud Deployment and Observability

The original university team project included Google Cloud deployment capabilities using:

- Google Cloud Run
- Google Artifact Registry
- Workload Identity Federation

The team also integrated OpenTelemetry for distributed tracing, with Google Cloud Trace as the tracing backend.

These are features of the broader team project and should not be interpreted as active deployment configuration in this personal repository.

---

## My Contributions

This was a collaborative university project. My individual contributions should be distinguished from the team's overall work.

**Individual contributions:** This section will be updated to describe the specific implementation, CI/CD, testing, security, containerization, or troubleshooting tasks I personally completed.

---

## Running Locally

### Prerequisites

- Go 1.27 or compatible toolchain
- Docker (optional, for containerized execution)

### Run the API

From the repository root:

```bash
go run ./cmd
```

The API uses port `8080` by default.

### Build the Docker image

```bash
docker build -t acetlisto:local .
```

### Run the container

```bash
docker run --rm -p 8080:8080 acetlisto:local
```

---

## API Examples

### Create an item

```bash
curl -X POST http://localhost:8080/items/ \
  -H "Content-Type: application/json" \
  -d '{"Name":"Oil","Description":"Should be nice!"}'
```

### List items

```bash
curl http://localhost:8080/items/
```

### Get a single item

Replace `ITEM_ID` with an ID returned by the API.

```bash
curl http://localhost:8080/items/ITEM_ID
```

### Update an item

```bash
curl -X PUT http://localhost:8080/items/ITEM_ID \
  -H "Content-Type: application/json" \
  -d '{"Name":"Olive Oil","Description":"Shopping list item"}'
```

### Delete an item

```bash
curl -X DELETE http://localhost:8080/items/ITEM_ID
```
