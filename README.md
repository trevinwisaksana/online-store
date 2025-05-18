# Online Store

A microservices-based online store application built using Go, Docker, and Protobuf.

## Description

This project is an implementation of a microservices architecture for an online store system. Each service is developed independently and communicates with others via an API Gateway and Protocol Buffers (Protobuf).

## Project Structure

The project consists of the following main components:

* **apigateway**: Handles routing of requests to the appropriate microservices.
* **googleapis**: Contains Protobuf-based API definitions.
* **proto**: Protocol Buffer definitions for service communication.
* **services/user-service**: The user service for authentication and user management.

## Technologies Used

* Go (Golang)
* Docker
* Docker Compose
* Protocol Buffers (Protobuf)
* API Gateway

## Prerequisites

Before you begin, make sure you have the following installed:

* [Docker](https://www.docker.com/products/docker-desktop)
* [Docker Compose](https://docs.docker.com/compose/)
* [Go](https://golang.org/dl/)

## Installation & Usage

1. **Clone the repository:**

   ```bash
   git clone https://github.com/trevinwisaksana/online-store.git
   cd online-store
   ```

2. **Build and start the Docker containers:**

   ```bash
   docker-compose up --build
   ```

3. **Access the application:**

   Once the containers are running, open your browser and go to [http://localhost:8080](http://localhost:8080).

## Development

To contribute to this project:

1. **Fork the repository.**

2. **Create a new branch:**

   ```bash
   git checkout -b your-feature-name
   ```

3. **Make your changes and commit them:**

   ```bash
   git commit -am 'Add new feature'
   ```

4. **Push to your fork:**

   ```bash
   git push origin your-feature-name
   ```

5. **Create a pull request.**

## License

This project is licensed under the [MIT License](LICENSE).

---

Let me know if you'd like me to include example API requests or expand any specific section (e.g. environment variables, service-specific details).
