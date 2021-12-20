# Contributing

The team welcomes contributions! To make changes, please read the following simple steps how to contribute.

## Always open an issue

Before doing anything always open an issue on GitHub, describing the contribution you would like to make, the bug you found or any other ideas you have. This will help us to get you started on the right foot.

It is recommended to wait for feedback before continuing to next steps. However, if
the issue is clear (e.g. a typo) and the fix is simple, you can continue and fix it.

## Code Contributions

If you have a bugfix or new feature that you would like to contribute, please do the following:
- Double check in open issues if there are any related issues or PRs
- Open an issue, ensure that you have properly described the use-case and possible solutions, link related issues/PRs if any
- Open a PR and link the issue created in previous step with your code changes.

Doing so allows to:
- Share knowledge and document a bug/missing feature
- Get feedback if someone is already working on it or is having a similar issue
- Benefit from the team experience by discussing it first, there are lots of implementation details that might not be
obvious at first sight.

### License Headers

We require license headers on all Golang files. This will be automatically checked when doing a CI build.

### Workflow

All feature development and most bug fixes hit the main branch first.
Pull requests should be reviewed by someone with commit access.
Once approved, the author of the pull request,
or reviewer if the author does not have commit access,
should "Squash and merge".

## Releasing

To make a release, the process is as follows:

For illustration purpose, `v1.2.3` will be the target release version, and the git remote will be `origin`.

1. Edit [CHANGELOG.md](CHANGELOG.md)
2. `git tag -a v1.2.3 -m "Release v1.2.3"`
3. `git push origin main v1.2.3` or `git push origin main --tags`
4. This will trigger the `My_Release` pipeline which will automatically push a new Docker image and create a GitHub release
