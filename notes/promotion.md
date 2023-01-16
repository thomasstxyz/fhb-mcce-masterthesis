# promotion

## what

- application version
- configuration settings
- static business settings
- non-static business settings

## how

### environment configuration tool

- kustomize
- helm
- jsonnet

### kinds of promotions

- preview environments (short-lived, dynamic, created from pull request)
- sequential promotion env after env (qa --> staging --> prod)
- multiple envs in same stage (prod-eu, prod-asia)
