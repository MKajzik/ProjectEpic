apiVersion: v1
kind: Secret
metadata:
  name: app-tier-secret
  namespace: free-epic
stringData:
  slack_url: ${slackURL}
  epic_url: ${epicURL}