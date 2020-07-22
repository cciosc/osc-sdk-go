# v0.0.1 to v0.0.9

- first (beta) versions of Outscale API

# v0.15 to v0.17

- same API features
- ... but breaking change with previous version
- SDK generated with openapi-generator
- add examples

# v1.0

- first stable version of Outscale API !
- breaking change if you come from v0.0.9 version
- more examples
- does not use a custom version of openapi-generator
- major new API calls support:
  - AccessKey actions
  - CheckAuthentication
  - Account actions
  - ReadConsumptionAccount
  - ListenerRule actions
  - ReadRegions
  - SecretAccessKey actions
  - ServerCertificate actions

# v1.1

- New region: cn-southeast-1

# 1.2

- SnapshotSize: fixed yaml type to avoid size limitation (as it is expressed in bytes)
- CreateAccessKey: has it's own object description
- New filters on ressources:
  - InternetService: LinkNetIds, LinkNetIds
  - Net: IsDefault
  - RouteTable: LinkRouteTableIds, LinkRouteTableMain, RouteDestinationServiceIds
  - SecurityGroup: NetIds
  - FlexibleGpuCatalog: Generations
- Documentation fixes