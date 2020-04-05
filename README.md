# Github Actions example

- Create Kubernetes cluster
- Create service account for Github Actions workflow

  ```bash
  kubectl create serviceaccount github-actions --namespace default
  SA_SECRET_NAME=$(kubectl get serviceaccount github-actions --namespace default --output go-template='{{ (index .secrets 0).name }}')

  kubectl get secret $SA_SECRET_NAME --namespace default --output go-template='{{ index .data "ca.crt" }}' | base64 --decode > kubeconfig-ca.crt
  KUBECONFIG_NS=$(kubectl get secret $SA_SECRET_NAME --namespace default --output go-template='{{ .data.namespace }}' | base64 --decode)
  KUBECONFIG_TOKEN=$(kubectl get secret $SA_SECRET_NAME --namespace default --output go-template='{{ .data.token }}' | base64 --decode)

  KUBECONFIG_SERVER=$(kubectl config view --minify --output go-template='{{ (index .clusters 0).cluster.server }}')

  kubectl --kubeconfig kubeconfig.yml config set-cluster production --server=$KUBECONFIG_SERVER --certificate-authority kubeconfig-ca.crt --embed-certs=true
  kubectl --kubeconfig kubeconfig.yml config set-credentials github-actions --token $KUBECONFIG_TOKEN
  kubectl --kubeconfig kubeconfig.yml config set-context github-actions-production --cluster production --user github-actions --namespace default
  kubectl --kubeconfig kubeconfig.yml config use-context github-actions-production

  kubectl create rolebinding github-actions --clusterrole edit --serviceaccount default:github-actions

  cat kubeconfig.yml | base64

  rm kubeconfig.yml kubeconfig-ca.crt
  ```

  - Kubeconfig in Github secret
    - Settings > Secrets > Add a new secret
      - Name: KUBECONFIG
        Value: Base64-encoded Kubeconfig file
      - Name: DOCKER_USERNAME
        Value: DockerHub username
      - Name: DOCKER_PASSWORD
        Value: DockerHub password
