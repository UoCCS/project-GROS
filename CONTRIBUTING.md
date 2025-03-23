# How to contribute to GROSlang

## Did you find a bug?

The GROSlang project uses GitHub as a bug tracker.  To report a bug, sign in to
your GitHub account, navigate to [GitHub issues](https://github.com/UoCCS/project-GROS/issues)
and click on **New issue** .

To be assigned to an issue, add a comment "take" to that issue.

Before you create a new bug entry, we recommend you first search among existing
GROSlang issues in [GitHub](https://github.com/UoCCS/project-GROS/issues).

We conventionally prefix the issue title with the componentname in brackets, 
such as "[rust_gc] Borrow-checker dereferences a null pointer", so as to make lists more easy to navigate, and
we'd be grateful if you did the same.

## Did you write a patch that fixes a bug or brings an improvement?

First create a GitHub issue as described above, selecting **Bug Report** or
**Enhancement Request**. Then, submit your changes as a GitHub Pull Request.
We'll ask you to prefix the pull request title with the GitHub issue number
and the component name in brackets. (for example: "GH-14736: [rust_gc] Automate the borrow-checker"). 
Respecting this convention makes it easier for us to process the backlog of submitted Pull Requests.

### Minor Fixes

Any functionality change should have a GitHub issue opened. For minor changes that
affect documentation, you do not need to open up a GitHub issue. Instead you can
prefix the title of your PR with "MINOR: " if meets one of the following:

*  Grammar, usage and spelling fixes that affect no more than 2 files
*  Documentation updates affecting no more than 2 files and not more
   than 500 words.

## Do you want to propose a significant new feature or an important refactoring?

We ask that all discussions about major changes in the codebase happen
publicly [GitHub](https://github.com/UoCCS/project-GROS/issues) issues.

## Do you have questions about the source code, the build procedure or the development process?

You can also ask in the project Github issues, see above.