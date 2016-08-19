# drone-rollbar

Deployment notifications for rollbar

## Configuring

```yaml
notify:
  rollbar:
    image: rschmukler/drone-rollbar
    access_token: $$ROLLBAR_ACCESS_TOKEN
    environment: $$BRANCH
```
