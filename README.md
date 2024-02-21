# Social TODO Service with Clean Architecture in Microservices

This project is a Social TODO Service implemented using the Go programming language, Gin framework, and follows the principles of Clean Architecture.

It utilizes various technologies and services to provide a robust and scalable solution. The key technologies and services used in this project are:

- **[AppContext](https://github.com/hoangtk0100/app-context)**: A custom library that manages common components, some used in this repository as described below.
- **gRPC**: An efficient and high-performance remote procedure call (RPC) framework designed for building scalable and reliable distributed systems.
- **Go Gin**: A lightweight web framework for building RESTful APIs in Go.
- **Pub/Sub with NATS**: A messaging system for building distributed systems and microservices architecture.
- **Redis Cache**: A fast in-memory data store used for caching frequently accessed data.
- **Asynchronous Job Group**: A mechanism for executing tasks asynchronously to improve performance and responsiveness.
- **Jaeger**: A distributed tracing system used to monitor and troubleshoot the service's performance.
- **Cloudflare R2**: A cloud storage service used for storing and retrieving files in a distributed manner.
- **PASETO**: A secure token format designed for authentication and authorization purposes. It aims to provide a more secure and versatile alternative to traditional JSON Web Tokens (JWTs).
- **MySQL**: The main database.

## Features

The Social TODO Service offers the following features:

- User management: Registration, authentication, and authorization.
- TODO creation and management: Users can create, update, and delete their TODO items.
- Social interactions: Users can like/unlike, view their TODOs, and interact with them.

## Architecture

The Social TODO Service follows a layered architecture pattern consisting of the following components:

- **Transport/Handler**: This layer is responsible for handling incoming requests such as RESTful, gRPC, or any from clients and parsing the request data into the desired struct format for the Business layer. It acts as a bridge between the external world and the business logic.

- **Business**: The Business layer handles the business logic of the application. It receives input from the Transport/Handler layer and implements the necessary computations, validations, and business-specific operations according to the defined use cases. It generates the appropriate response and communicates with the Repository layer if needed.

- **Repository**: The Repository layer interacts with local/remote storage and is responsible for communicating with specific database engines. It handles the actual storage and retrieval of data, interacting with databases such as MySQL and Redis. It provides an interface to perform CRUD (Create, Read, Update, Delete) operations on the data.

- **Entity**: The Entity layer defines the models or data structures. It encapsulates the business entities or objects in the system and provides methods and properties related to those entities. The Entity layer helps in managing the state and behavior of the application's core objects. It serves as a central representation of the application's data model and is utilized by the Business and Repository layers.

The overall flow in the architecture is as follows: The Transport/Handler layer receives requests, parses them into the appropriate struct format, and passes them to the Business layer. The Business layer processes the requests, performs necessary computations and business logic, and generates responses. If required, the Business layer can interact with the Repository layer to retrieve or store data. The Entity layer plays a central role in defining and manipulating the application's data model, providing a consistent and unified representation of entities across the system.

By incorporating the Entity layer, the Social TODO Service achieves better organization and abstraction of the application's data entities, promoting code reuse, and encapsulation of business logic. The Entity layer helps ensure a clear and unified understanding of the core entities and their relationships within the application.

This architecture promotes separation of concerns, modularization, and maintainability by decoupling the different layers and their responsibilities. Each layer has a specific role and can be developed, tested, and scaled independently.

## Contribution

Contributions to this project are welcome! If you find any issues or have ideas for enhancements, feel free to open an issue or submit a pull request. Please follow the project's guidelines for contribution and code formatting.