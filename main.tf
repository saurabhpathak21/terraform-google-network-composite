

locals {
  project_name = "acceleration-${var.type}"
  network_name = var.type
  region       = "europe-west2"
  rules = [
    for f in var.firewall_rules : {
      name                    = f.name
      direction               = f.direction
      priority                = lookup(f, "priority", null)
      description             = lookup(f, "description", null)
      ranges                  = lookup(f, "ranges", null)
      source_tags             = lookup(f, "source_tags", null)
      source_service_accounts = lookup(f, "source_service_accounts", null)
      target_tags             = lookup(f, "target_tags", null)
      target_service_accounts = lookup(f, "target_service_accounts", null)
      allow                   = lookup(f, "allow", [])
      deny                    = lookup(f, "deny", [])
      log_config              = lookup(f, "log_config", null)
    }
  ]
}


/******************************************
  Hub Project Creation
 *****************************************/
module "project" {
  source  = "terraform-google-modules/project-factory/google"
  version = "14.2.0"

  random_project_id              = true
  name                           = local.project_name
  org_id                         = var.organization_id
  folder_id                      = var.folder_id
  billing_account                = var.billing_account
  enable_shared_vpc_host_project = true
  default_network_tier           = var.default_network_tier

  activate_apis = [
    "compute.googleapis.com",
    "dns.googleapis.com",
    "cloudresourcemanager.googleapis.com"
  ]
}


/******************************************
	IAM
 *****************************************/

resource "google_project_iam_binding" "project" {
  project = module.project.project_id
  role    = "roles/editor"
  members = [
    "user:saurabh.pathaks21@gmail.com"
  ]
}


#Shared VPC

/******************************************
	VPC configuration
 *****************************************/
module "vpc" {
  source                                 = "terraform-google-modules/network/google//modules/vpc"
  version                                = "6.0.1"
  network_name                           = "${local.network_name}-acceleration-xpn-001"
  project_id                             = module.project.project_id
  delete_default_internet_gateway_routes = true

}

/******************************************
	Subnet configuration
 *****************************************/
module "subnets" {
  source  = "terraform-google-modules/network/google//modules/subnets"
  version = "6.0.1"

  project_id       = module.project.project_id
  network_name     = module.vpc.network_name
  subnets          = var.subnets
  secondary_ranges = var.secondary_ranges
}

/******************************************
	Routes
 *****************************************/
module "routes" {
  source  = "terraform-google-modules/network/google//modules/routes"
  version = "6.0.1"

  project_id        = module.project.project_id
  network_name      = module.vpc.network_name
  routes            = var.routes
  module_depends_on = [module.subnets.subnets]
}

/******************************************
	Firewall rules
 *****************************************/

module "firewall_rules" {
  source  = "terraform-google-modules/network/google//modules/firewall-rules"
  version = "6.0.1"

  project_id   = module.project.project_id
  network_name = module.vpc.network_name
  for_each     = { for rule in local.rules : "${rule.fileName}--${rule.id}" => rule }
  rules        = each.value
}

/******************************************
	DNS
 *****************************************/

module "dns-private-zone" {
  source  = "terraform-google-modules/cloud-dns/google"
  version = "4.2.1"

  project_id                         = module.project.project_id
  type                               = "private"
  name                               = "${local.network_name}-acceleration"
  domain                             = "new-acceleration.com."
  private_visibility_config_networks = [module.vpc.network_self_link]

  recordsets = [
    {
      name = "localhost"
      type = "A"
      ttl  = 300
      records = [
        "127.0.0.1",
      ]
    },
  ]
}

//DNS Peering- Enable to peer the DNS


module "dns-peering-zone" {
  source  = "terraform-google-modules/cloud-dns/google"
  version = "3.0.0"

  count                              = var.enable_peering == true ? 1 : 0
  project_id                         = module.project.project_id
  type                               = "peering"
  name                               = "${local.network_name}-acceleration-peering"
  domain                             = "new-acceleration.com."
  private_visibility_config_networks = [module.vpc.network_self_link]
  target_network                     = var.target_network_self_link #Update the target network to which DNS peering is required
}


/******************************************
	Router
 *****************************************/


//Create Router

module "cloud_router" {
  source  = "terraform-google-modules/cloud-router/google"
  version = "5.0.0"

  name    = "${local.network_name}-${var.router_name}"
  project = module.project.project_id
  region  = local.region
  network = module.vpc.network_name
}

/******************************************
	VPN
 *****************************************/

module "vpn" {
  source  = "terraform-google-modules/vpn/google//modules/vpn_ha"
  version = "~> 1.3.0"


  project_id  = module.project.project_id
  region      = local.region
  network     = module.vpc.network_self_link
  router_name = module.cloud_router.router.name
  name        = "${local.network_name}-vpn"

  #update with the details of the on-prem network 

  peer_external_gateway = {
    redundancy_type = "SINGLE_IP_INTERNALLY_REDUNDANT"
    interfaces = [{
      id         = 0
      ip_address = "8.8.8.8" # on-prem router ip address

    }]
  }

  router_asn = 64514
  tunnels = {
    remote-0 = {
      bgp_peer = {
        address = "169.254.1.1"
        asn     = 64513
      }
      bgp_peer_options                = null
      bgp_session_range               = "169.254.1.2/30"
      ike_version                     = 2
      vpn_gateway_interface           = 0
      peer_external_gateway_interface = 0
      shared_secret                   = ""
    }
    remote-1 = {
      bgp_peer = {
        address = "169.254.2.1"
        asn     = 64513
      }
      bgp_peer_options                = null
      bgp_session_range               = "169.254.2.2/30"
      ike_version                     = 2
      vpn_gateway_interface           = 1
      peer_external_gateway_interface = 0
      shared_secret                   = ""
    }
  }
  depends_on = [module.cloud_router]
}