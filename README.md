[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/thegalactiks/giteway/ci.yml?branch=main&label=core%20build&style=for-the-badge)](https://github.com/thegalactiks/giteway/actions/workflows/ci.yml)
![Latest version](https://img.shields.io/github/v/release/thegalactiks/giteway?sort=semver&style=for-the-badge)
[![Github Repo Stars](https://img.shields.io/github/stars/thegalactiks/giteway?style=for-the-badge)](https://github.com/thegalactiks/giteway)
![License](https://img.shields.io/github/license/thegalactiks/giteway?style=for-the-badge)

# Giteway - HTTP Gateway for Git Services

Giteway is an HTTP Gateway for Git services, providing a standardized interface for fetching Git information and commits from various service providers, including GitHub, GitLab, Bitbucket *(soon)*, and SSH Git *(soon)*.

## Features

- **Standardized Interface**: Giteway exposes a consistent API endpoint for interacting with Git repositories, irrespective of the underlying service provider.
- **Multi-Provider Support**: Easily fetch Git data from GitHub and GitLab.
- **Authentication**: Securely handle authentication for different Git service providers, supporting various authentication mechanisms.
- **HTTP Gateway**: Access Git data through standard HTTP requests, making it easy to integrate into web applications.
- **Modular Design**: Designed with modularity in mind, making it extensible for adding support for new Git providers.

## Getting Started

To get started with Giteway, you can deploy it as a Docker container using the following steps:

1. Pull the Giteway Docker image from the Docker Hub:

    ```bash
    docker pull galactiks/giteway:latest
    ```

2. Run the Docker container, exposing the necessary ports:

    ```bash
    docker run -d -p 5000:5000 galactiks/giteway:latest
    ```

    This command will start the Giteway container and map port 5000 of the container to port 5000 on your host machine.

3. Verify that Giteway is running by accessing the API endpoint in your browser or using a tool like cURL:

    ```bash
    curl http://localhost:5000/repos/github.com/thegalactiks
    ```

    You should receive a response with the list of public repositories.

4. You can now start using Giteway to interact with Git repositories through the exposed API.

For more detailed information on how to use Giteway and its API, refer to the [API Documentation](https://docs.galactiks.com/giteway/reference/api/).

## Configuration

Giteway can be configured using environment variables to customize its behavior. The following environment variables are supported:

| Variable Name            | Description | Default Value |
|--------------------------|-------------|---------------|
| `GITEWAY_SERVE_BASE_URL` | The base URL at which Giteway is served. |  |
| `GITEWAY_SERVE_PORT`     | The port at which Giteway is served. | `5000` |
| `GITEWAY_SERVE_CORS_ENABLED` | Indicates whether CORS is enabled for Giteway. | `false` |
| `GITEWAY_SERVE_CORS_ALLOW_ORIGINS` | The list of allowed origins for CORS. | `*` |
| `GITEWAY_SERVE_CORS_ALLOWED_METHODS` | The list of allowed HTTP methods for CORS. | `GET,POST` |
| `GITEWAY_SERVE_CORS_ALLOW_HEADERS` | The list of allowed headers for CORS. | `Authorization, Content-Type, Cookie` |
| `GITEWAY_SERVE_CORS_EXPOSE_HEADERS` | The list of exposed headers for CORS. | `Content-Type, Set-Cookie` |
| `GITEWAY_SERVE_CORS_ALLOW_CREDENTIALS` | Indicates whether CORS allows credentials. | `true` |
| `GITEWAY_SERVE_TIMEOUT` | The timeout duration for Giteway requests. | `5000` |
| `GITEWAY_LOGGING_LEVEL` | The logging level for Giteway. | `-1` |
| `GITEWAY_LOGGING_ENCODING` | The encoding format for Giteway logs. | `console ` |
| `GITEWAY_LOGGING_DEVELOPMENT` | Indicates whether Giteway is in development mode. | `true` |
| `GITEWAY_GITHUB_PRIVATE_KEY_PATH` | The file path to the private key for GitHub authentication. |  |
| `GITEWAY_GITHUB_APP_ID` | The ID of the GitHub app. |  |
| `GITEWAY_GITHUB_INSTALLATIONS` | The map of GitHub installations. |  |

## API Documentation

Documentation for the Giteway API can be found [here](https://docs.galactiks.com/giteway/reference/api/).

## License

This project is licensed under the [MIT License](https://github.com/thegalactiks/giteway/LICENSE) @ [Galactiks](https://www.galactiks.com/).
