# Rose Park

## Getting Started

### Prerequisites
- An API and test Application set up through [Auth0](https://auth0.com)
- Go 1.15
- Postgres 12.4
- GNU Make
- [Air](https://github.com/cosmtrek/air) 
- [Lair](https://github.com/weblair/lair)

### After Cloning
- Create a `.env` file to set the following to their appropriate values
  (these values can be obtained from the Auth0 dashboard):
    - `AUTH0_API_AUDIENCE`
    - `AUTH0_JWK`
- From within the `database` directory, execute the following commands:
    - `lair db create`
    - `lair db migrate`
- To run in development mode, use `make run`

### Running Tests
Before running tests, the following will need to be set in the `.env` file:
- `AUTH0_TEST_AUTH_URL`
- `AUTH0_TEST_CLIENT_ID`
- `AUTH0_TEST_CLIENT_SECRET`

## Contribution Guidelines

### Commit Messages
The rules of [conventional commit messages](https://www.conventionalcommits.org/en/v1.0.0-beta.2/) should be observed.
Observe to keep the first line of the commit message down to 50 characters and insert hard line-breaks at 72 characters
for the rest of the message body.

When working on `feature` or `hotfix` branches, the rules can be relaxed a bit. PRs should only be opened from your 
`develop` branch, and when wrapping up your `feature` branches, you should squash your commits.

#### Tags
  - fix&mdash;for bugfixes
  - feat&mdash;for any new functionality
  - BREAKING CHANGE&mdash;annotation in the commit message body for any changes that will affect backwards-compatability.
  - refactor&mdash;for reworked code that ends up being functionally the same
  - docs&mdash;for changes to docstrings, CHANGELOG.md, this README, etc.
  - chore&mdash;for changes to the repo that don't affect functional code or
    docs (i.e. Makefiles, Dockerfiles, etc.)
