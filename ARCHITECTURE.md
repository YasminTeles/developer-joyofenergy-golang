# About the project

## Endpoints

This simple server provides the following end points:

- `GET /version`
 That returns the version of the server. It's useful for blue-green development.

- `GET /healthcheck`
That returns the health of the server running. It's useful for checking if the server can handle requests.

- `POST /price-plans/estimate`
That returns the electricity cost and the lowest tariffs. You need to send usage data.

    **Input**

    ```json
      {
        "smartMeterId": "smart-meter-0",
        "electricityReadings": [{ "time": <UTC-Time>, "reading": <reading> }],
      }
    ```

    `time`: UTC time, e.g. `2024-04-13T19:41:26Z`

    `reading`: kW reading of smart meter at that time, e.g. `0.0503`

    **Output**

    ```json
        {
          "electricityCost": 159.0,
          "recommendations": {
            "recommendations": [
              {
                "key": "price-plan-2",
                "value": 11.346088258014543
              },
              {
                "key": "price-plan-1",
                "value": 22.692176516029086
              },
              {
                "key": "price-plan-0",
                "value": 113.46088258014544
              }
            ]
          }
        }
    ```

## Built with

This project uses the following technologies:

Continuos integration:

- [Editor config](https://editorconfig.org/) - EditorConfig helps maintain consistent coding styles for multiple developers working on the same project across various editors and IDEs.
- [Github Actions](https://docs.github.com/en/actions) - Automate, customize, and execute your software development workflows right in your repository with GitHub Actions.
- [Dependabot](https://github.com/dependabot) - Automated dependency updates built into GitHub.
- [Semantic release](https://github.com/semantic-release) - Fully automated version management and package publishing.
- [Gitleaks](https://github.com/gitleaks/gitleaks) - Gitleaks is an open-source tool to detect and prevent secrets in Git repositories.
- [Golangci-lint](https://github.com/golangci/golangci-lint) - Linter for your Golang code.
- [Trivy](https://trivy.dev/) - The container vulnerability scanner.

Tests:

- [Testify](https://github.com/stretchr/testify) - Provides a set of tools for writing tests in Go.

Continuos delivery:

- [Docker](https://www.docker.com/) - Docker enables you to separate your applications from your infrastructure so you can deliver software quickly.

Troubleshooting features:

- **Structured and centralized logger**: using [Logrus](https://github.com/sirupsen/logrus);
- **Log each request received**;
- **Unique identifier request**;

## Versions

We use [Semantic version](http://semver.org) for versioning. For versions available, see [changelog](CHANGELOG.md).

## License

This project is licensed under the [Creative Commons License](LICENSE).
