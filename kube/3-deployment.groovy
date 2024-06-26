apiVersion: apps/v1
kind: Deployment
metadata:
  name: epic-deployment
  labels:
    app: epic
spec:
  replicas: 1
  selector:
    matchLabels:
      app: epic
  template:
    metadata:
      labels:
        app: epic
    spec:
      containers:
      - name: epic
        image: registry.togi.cloud/epic:${buildNumber}
        imagePullPolicy: Always
        env:
          - name: SLACK_URL
            valueFrom:
              secretKeyRef:
                name: app-tier-secret
                key: slack_url
          - name: EPIC_URL
            valueFrom:
              secretKeyRef:
                name: app-tier-secret
                key: epic_url
        ports:
        - containerPort: 5000
