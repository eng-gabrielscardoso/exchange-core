# Contributing

Thank you for considering contributing to the project! We welcome contributions from anyone who wants to improve or enhance the project. By following these guidelines, you can ensure that your contributions are effective and aligned with the project's goals.

## Installation

1. Firstly, you should have installed the Golang SDK in your machine (*we haven't done the dockerisation yet*) and the IDE of your preference (we strongly recommend the Visual Studio Code). Also, make sure to have the most recent version of Docker and Docker Compose to run the necessary services such as the Broker and Control Center.
2. Install the dependencies using the command: `go mod tidy`
3. Before serve application, up the containers with necessary infrastructure using the command `docker-compose up -d`. Depend from your connection speed this command could take a long (for a fresh installation this command will pull the necessary images in your local), so take a coffee and wait a little
4. Serve the application using the following command: `go run cmd/trade/main.go`
5. Build something incredible 🌟

## How to contribute

1. Fork the repository and clone it locally.
2. Create a new branch from the `develop` branch to work on your changes.
3. Implement your changes, ensuring that your code follows the project's coding style and guidelines.
4. Write tests to validate your changes and ensure that existing functionality remains unaffected.
5. Document any new features, changes, or significant updates in the project's documentation.
6. Commit your changes with a clear and descriptive commit message.
7. Push your branch to your forked repository.
8. Submit a pull request (PR) to the `develop` branch of the original repository. Clearly describe the purpose and scope of your changes in the PR description.
9. Engage in discussions and address any feedback or suggestions provided by the maintainers or reviewers.
10. Once your changes have been reviewed and approved, they will be merged into the `develop` branch.

## Guidelines

- Follow the established coding style and conventions used in the project.
- Ensure that your changes do not introduce new warnings or errors when building or running the project.
- Keep your changes focused and limited to the purpose of the PR. If you have multiple unrelated changes, consider submitting separate PRs.
- Be respectful and professional when interacting with others in the project. Provide constructive feedback and engage in meaningful discussions.
- Prioritize writing clear and concise documentation to help users understand and utilize the project's features effectively.
- Test your changes thoroughly to ensure they work as intended and do not introduce regressions.
- Be responsive to any comments, feedback, or requests for changes from the maintainers or reviewers.
