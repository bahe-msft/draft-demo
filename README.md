# gator usage experiences

## Goals

- Enforce policies on deployment specs in a mono repo environment
- Detect issues with policies without the need of deploying to a real cluster (shift left)

## Process Overview

Currently we are enforcing via a build job in pull request phrase, which does the following things:

- Build the policies into a single YAML

  ```
  # contains policies for different teams
  $ ls ./policy
  team-a/policy.yaml team-b/policy.yaml kustomization.yaml
  $ kustomize build ./policy > policy.yaml
  ```
- Build the deployment specs into a single YAML

  ```
  $ kustomize build ./app/overlays/production > production-app.yaml
  ```

- Run the tests

  ```
  $ gator test --filename policy.yaml --filename production-app.yaml
  ```