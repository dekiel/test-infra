moved {
  from = module.artifact_registry["modules-internal"].google_artifact_registry_repository.artifact_registry
  to   = module.artifact_registry["modules-internal"].google_artifact_registry_repository.protected_repository[0]
}

# TODO (dekiel): remove after migration to modulectl is done
module "artifact_registry" {
  source = "../../modules/artifact-registry"

  providers = {
    google = google.kyma_project
  }


  for_each                  = var.kyma_project_artifact_registry_collection
  repository_name           = each.value.name
  description               = each.value.description
  type                      = each.value.type
  immutable_tags            = each.value.immutable
  multi_region              = each.value.multi_region
  owner                     = each.value.owner
  repoAdmin_serviceaccounts = each.value.repoAdmin_serviceaccounts
  reader_serviceaccounts    = each.value.reader_serviceaccounts
  public                    = each.value.public
  cleanup_policy_dry_run    = each.value.cleanup_policy_dry_run
  cleanup_policies          = each.value.cleanup_policies
}

moved {
  from = module.prod_docker_repository.google_artifact_registry_repository.artifact_registry
  to   = module.prod_docker_repository.google_artifact_registry_repository.protected_repository[0]
}

module "prod_docker_repository" {
  source = "../../modules/artifact-registry"

  providers = {
    google = google.kyma_project
  }

  repository_name            = var.prod_docker_repository.name
  description                = var.prod_docker_repository.description
  location                   = var.prod_docker_repository.location
  immutable_tags             = var.prod_docker_repository.immutable_tags
  format                     = var.prod_docker_repository.format
  cleanup_policies           = var.prod_docker_repository.cleanup_policies
  cleanup_policy_dry_run     = var.prod_docker_repository.cleanup_policy_dry_run
  repository_prevent_destroy = var.prod_docker_repository.repository_prevent_destroy
}

moved {
  from = module.dev_docker_repository.google_artifact_registry_repository.artifact_registry
  to   = module.dev_docker_repository.google_artifact_registry_repository.protected_repository[0]
}

module "dev_docker_repository" {
  source = "../../modules/artifact-registry"

  providers = {
    google = google.kyma_project
  }

  repository_name            = var.dev_docker_repository.name
  description                = var.dev_docker_repository.description
  location                   = var.dev_docker_repository.location
  immutable_tags             = var.dev_docker_repository.immutable_tags
  format                     = var.dev_docker_repository.format
  cleanup_policies           = var.dev_docker_repository.cleanup_policies
  cleanup_policy_dry_run     = var.dev_docker_repository.cleanup_policy_dry_run
  type                       = var.dev_docker_repository.type
  repository_prevent_destroy = var.dev_docker_repository.repository_prevent_destroy
}
