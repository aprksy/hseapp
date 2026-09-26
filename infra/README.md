# Infrastructure

Infrastructure as Code (IaC) and deployment configurations for HSE App.

## Directory Structure

```
infra/
├── terraform/          # Terraform modules for cloud resources
│   ├── modules/        # Reusable modules (database, storage, compute)
│   ├── environments/   # Environment-specific configs (dev, staging, prod)
│   └── state/          # Remote state configuration
├── k8s/                # Kubernetes manifests
│   ├── base/           # Base configurations
│   ├── overlays/       # Environment-specific overlays (kustomize)
│   └── helm/           # Helm charts (if needed)
└── scripts/            # Deployment and maintenance scripts
    ├── deploy.sh       # Deployment automation
    ├── backup.sh       # Database backup scripts
    └── migrate.sh      # Schema migration scripts
```

## Technologies

- **Cloud Provider**: AWS / GCP / Azure (configurable)
- **Database**: Supabase (PostgreSQL) or Firebase Firestore
- **Container Orchestration**: Kubernetes (EKS/GKE/AKS)
- **CI/CD**: GitHub Actions
- **Monitoring**: Prometheus + Grafana, OpenTelemetry

## Quick Start

### Prerequisites
- Terraform >= 1.5.0
- kubectl >= 1.27.0
- Helm >= 3.12.0 (optional)

### Deploy to Development Environment

```bash
cd infra/terraform/environments/dev
terraform init
terraform plan -var-file=dev.tfvars
terraform apply -var-file=dev.tfvars
```

### Deploy to Kubernetes

```bash
cd infra/k8s
kubectl apply -k overlays/dev
```

## Configuration

Environment variables should be managed through:
- Terraform variables for cloud resources
- Kubernetes ConfigMaps/Secrets for application config
- External secrets manager (AWS Secrets Manager, GCP Secret Manager)

## Related Documentation

- [Architecture](../docs/02-ARCHITECTURE.md)
- [Development Setup](../docs/03-DEVELOPMENT.md)
- [Makefile Commands](../Makefile)
