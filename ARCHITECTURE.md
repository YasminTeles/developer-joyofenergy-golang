# About the project

JOI Energy is a new energy company that uses data to ensure customers are able to be on the best price plan for their energy consumption.

## Endpoints

This simple server provides the following end points:

- `GET /version`
 That returns the version of the server. It's useful for blue-green development.

- `GET /healthcheck`
 That returns the health of the server running. It's useful for check if the server can be able handling requests.

- `GET /readings/read/<smartMeterId>`
  Get Stored Readings.

    **Output**

    ```json
      [
        { "time": "2017-09-07T10:37:52.362Z", "reading": 1.3524882598124337 },
        { "time": "2024-05-02T16:10:20.177916414Z", "reading": 0.037697719480008995 }
      ]
    ```

- `GET /price-plans/recommend/<smartMeterId>[?limit=<limit>]`
  View Recommended Price Plans for Usage.

    `smartMeterId`: A string value, e.g. `smart-meter-0`

    `limit`: Optional limit to display only a number of price plans, e.g. `2`

    **Output**

    ```json
      [
        { "price-plan-0": 15.084324881035297 },
        ...
      ]
    ```

- `GET /price-plans/compare-all/<smartMeterId>`
  View Current Price Plan and Compare Usage Cost Against all Price Plans.

    `smartMeterId`: A string value, e.g. `smart-meter-0`

    **Output**

    ```json
      {
        "pricePlanId": "price-plan-2",
        "pricePlanComparisons": {
            "price-plan-0": 21.78133785680731809,
            ...
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
