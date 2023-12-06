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

## Problems

### Gator only supports GVK objects

Take the minimum kustomization.yaml as an example, which doesn't require to have GVK set. In this case, gator would throw error like this:

> auditing objects: adding data of GVK "/, Kind=": admission.k8s.gatekeeper.sh: invalid request object: resource has no version

Sample failing build:

https://github.com/bahe-msft/draft-demo/actions/runs/7107991414/job/19350445670?pr=3

**Why we want this?**:

1. Some policies might more suitable for common level configurations, for example, common labels/annotations.
   It's easier to verify on shared configuration like kustomization.yaml (`kustomize`) or values.yaml (`helm`)
2. In pull request phrase, we want to focus more on the actual changes instead of the merged changes.
   Errors from the merged files are difficult to trace back to the source file changes.


### Most of the gatekeeper-library policies are pod oriented

Gatekeeper checks in pod level, therefore most of the rules from the gatekeeper-library are expecting the target is a pod.
However, when running in the same rule in pull request phrase, we cannot detect any issues.

Sample build:

- PR: https://github.com/bahe-msft/draft-demo/pull/2
- build: https://github.com/bahe-msft/draft-demo/actions/runs/7107939397/job/19350307526?pr=2


**Why we want this?**:

1. In shift-left practice, we want to verify changes before actually deploying into a cluster. If the policy is controller aware,
   then we can make the policies more reusable.

### Helm support

Should user build the chart first, or can we let gator build the chart instead?