// generated code; DO NOT EDIT

package config

var docStrings = map[string]string{
	"registration":                    "The plugins used when publishing Registration data, when this is unset or empty sending registration data is disabled",
	"registration_collective":         "The Sub Collective to publish registration data to",
	"registerinterval":                "How often to publish registration data",
	"registration_splay":              "When true delays initial registration publish by a random period up to registerinterval following registration publishes will be at registerinterval without further splay",
	"collectives":                     "The list of known Sub Collectives this node will join or communicate with, Servers will subscribe the node and each agent to each sub collective and Clients will publish to a chosen sub collective. Defaults to the build settin build.DefaultCollectives",
	"main_collective":                 "The Sub Collective where a Client will publish to when no specific Sub Collective is configured",
	"logfile":                         "The file to write logs to, when set to 'discard' logging will be disabled. Also supports 'stdout' and 'stderr' as special log destinations.",
	"loglevel":                        "The lowest level log to add to the logfile",
	"libdir":                          "The directory where Agents, DDLs and other plugins are found",
	"identity":                        "The identity this machine is known as, when empty it's derived based on the operating system hostname or by calling facter fqdn",
	"direct_addressing":               "Enables the direct-to-node communications pattern, unused in the Go clients",
	"color":                           "Disables or enable CLI color",
	"securityprovider":                "Used to select the security provider in Ruby clients, only sensible value is \"choria\"",
	"connector":                       "Configures the network connector to use, only sensible value is \"nats\", unused in Go based code",
	"classesfile":                     "Path to a file listing configuration classes applied to a node, used in matches using Class filters",
	"discovery_timeout":               "How long to wait for responses while doing broadcast discovery",
	"publish_timeout":                 "Ruby clients use this to determine how long they will allow when publishing requests",
	"connection_timeout":              "Ruby clients use this to determine how long they will try to connect, fails after timeout",
	"rpcaudit":                        "When enabled uses rpcauditprovider to audit RPC requests processed by the server",
	"rpcauditprovider":                "The audit provider to use, unused at present as there is only a \"choria\" one",
	"rpcauthorization":                "When enables authorization is performed on every RPC request based on rpcauthprovider",
	"rpcauthprovider":                 "The Authorization system to use",
	"rpclimitmethod":                  "When limiting nodes to a subset of discovered nodes this is the method to use, random is influenced by",
	"logger_type":                     "The type of logging to use, unused in Go based programs",
	"threaded":                        "Enables multi threaded mode in the Ruby client, generally a bad idea",
	"ttl":                             "How long published messages are allowed to linger on the network, lower numbers have a higher reliance on clocks being in sync",
	"default_discovery_method":        "The default discovery plugin to use. The default \"mc\" uses a network broadcast, \"choria\" uses PuppetDB, external calls external commands",
	"plugin.yaml":                     "Where to look for YAML or JSON based facts",
	"default_discovery_options":       "Default options to pass to the discovery plugin",
	"plugin.choria.puppetserver_host": "The hostname where your Puppet Server can be found",
	"plugin.choria.puppetserver_port": "The port your Puppet Server listens on",
	"plugin.choria.puppetca_host":     "The hostname where your Puppet Certificate Authority can be found",
	"plugin.choria.puppetca_port":     "The port your Puppet Certificate Authority listens on",
	"plugin.choria.puppetdb_host":     "The host hosting your PuppetDB, used by the \"choria\" discovery plugin",
	"plugin.choria.puppetdb_port":     "The port your PuppetDB listens on",
	"plugin.choria.use_srv":           "If SRV record lookups should be attempted to find Puppet, PuppetDB, Brokers etc",
	"plugin.choria.srv_domain":        "The domain to use for SRV records, defaults to the domain the server FQDN is in",
	"plugin.choria.server.provision":  "Specifically enable or disable provisioning",
	"plugin.choria.discovery.external.command":                 "The command to use for external discovery",
	"plugin.choria.discovery.inventory.source":                 "The file to read for inventory discovery",
	"plugin.choria.discovery.broadcast.windowed_timeout":       "Enables the experimental dynamic timeout for choria/mc discovery",
	"plugin.choria.federation.collectives":                     "List of known remote collectives accessible via Federation Brokers",
	"plugin.choria.federation_middleware_hosts":                "Middleware brokers used by the Federation Broker, if unset uses SRV",
	"plugin.choria.federation.cluster":                         "The cluster name a Federation Broker serves",
	"plugin.choria.stats_address":                              "The address to listen on for statistics",
	"plugin.choria.stats_port":                                 "The port to listen on for HTTP requests for statistics, setting to 0 disables it",
	"plugin.choria.legacy_lifecycle_format":                    "When enabled will publish lifecycle events in the legacy format, else Cloud Events format is used",
	"plugin.nats.user":                                         "The user to connect to the NATS server as. When unset no username is used.",
	"plugin.nats.pass":                                         "The password to use when connecting to the NATS server",
	"plugin.nats.credentials":                                  "The NATS 2.0 credentials to use, required for accessing NGS",
	"plugin.nats.ngs":                                          "Uses NATS NGS global managed network as middleware, overrides broker names to \"connect.ngs.global\"",
	"plugin.choria.middleware_hosts":                           "Set specific middleware hosts in the format host:port, if unset uses SRV",
	"plugin.choria.randomize_middleware_hosts":                 "Shuffle middleware hosts before connecting to spread traffic of initial connections",
	"plugin.choria.network.listen_address":                     "Address the Network Broker will listen on",
	"plugin.choria.network.websocket_port":                     "Port to listen on for websocket connections",
	"plugin.choria.network.websocket_advertise":                "The URL to advertise for websocket connections",
	"plugin.choria.network.client_port":                        "Port the Network Broker will accept client connections on",
	"plugin.choria.network.client_tls_force_required":          "Force requiring/not requiring TLS for all clients",
	"plugin.choria.network.client_anon_tls":                    "Use anonymous TLS for client connections (disables verification)",
	"plugin.choria.network.peer_port":                          "Port used to communicate with other local cluster peers",
	"plugin.choria.network.peer_user":                          "Username to use when connecting to cluster peers",
	"plugin.choria.network.peer_password":                      "Password to use when connecting to cluster peers",
	"plugin.choria.network.peers":                              "List of cluster peers in host:port format",
	"plugin.choria.network.leafnode_port":                      "Port to listen on for Leafnode connections, disabled with 0",
	"plugin.choria.network.leafnode_remotes":                   "Remote networks to connect to as a Leafnode",
	"plugin.choria.network.gateway_port":                       "Port to listen on for Super Cluster connections",
	"plugin.choria.network.gateway_name":                       "Name for the Super Cluster",
	"plugin.choria.network.gateway_remotes":                    "List of remote Super Clusters to connect to",
	"plugin.choria.network.write_deadline":                     "How long to allow clients to process traffic before treating them as slow, increase this on large networks or slow networks",
	"plugin.choria.network.client_hosts":                       "CIDRs to limit client connections from, appropriate ACLs are added based on this",
	"plugin.choria.network.deny_server_connections":            "Set ACLs denying server connections to this broker",
	"plugin.choria.network.tls_timeout":                        "Time to allow for TLS connections to establish, increase on slow or very large networks",
	"plugin.choria.network.public_url":                         "Name:Port to advertise to clients, useful when fronted by a proxy",
	"plugin.choria.network.stream.store":                       "Enables Streaming data persistence stored in this path",
	"plugin.choria.network.stream.event_retention":             "When not zero enables retaining Lifecycle events in the Stream Store",
	"plugin.choria.network.stream.event_replicas":              "When configuring LifeCycle events ensure data is replicated in the cluster over this many servers",
	"plugin.choria.network.stream.machine_retention":           "When not zero enables retaining Autonomous Agent events in the Stream Store",
	"plugin.choria.network.stream.machine_replicas":            "When configuring Autonomous Agent event storage ensure data is replicated in the cluster over this many servers",
	"plugin.choria.network.stream.advisory_retention":          "When not zero enables retaining Stream advisories in the Stream Store",
	"plugin.choria.network.stream.advisory_replicas":           "When configuring Stream advisories storage ensure data is replicated in the cluster over this many servers",
	"plugin.choria.network.system.user":                        "Username used to access the Choria system account",
	"plugin.choria.network.system.password":                    "Password used to access the Choria system account",
	"plugin.choria.network.pprof_port":                         "The port the network broker will listen on for pprof requests",
	"plugin.choria.network.provisioning.signer_cert":           "Path to the public cert that signs provisioning tokens, enables accepting provisioning connections into the provisioning account",
	"plugin.choria.network.provisioning.client_password":       "Password the provisioned clients should use to connect",
	"plugin.choria.broker_network":                             "Enables the Network Broker",
	"plugin.choria.broker_federation":                          "Enables the Federation Broker",
	"plugin.choria.registration.file_content.data":             "YAML or JSON file to use as data source for registration",
	"plugin.choria.registration.file_content.target":           "NATS Subject to publish registration data to",
	"plugin.choria.registration.file_content.compression":      "Enables gzip compression of registration data",
	"plugin.choria.registration.inventory_content.compression": "Enables gzip compression of registration data",
	"plugin.choria.registration.inventory_content.target":      "NATS Subject to publish registration data to",
	"plugin.choria.agent_provider.mcorpc.agent_shim":           "Path to the helper used to call MCollective Ruby agents",
	"plugin.choria.agent_provider.mcorpc.config":               "Path to the MCollective configuration file used when running MCollective Ruby agents",
	"plugin.choria.agent_provider.mcorpc.libdir":               "Path to the libdir MCollective Ruby agents should have",
	"plugin.choria.ssldir":                                     "The SSL directory, auto detected via Puppet, when specifically set Puppet will not be consulted",
	"plugin.choria.security.privileged_users":                  "Patterns of certificate names that would be considered privileged and able to set custom callers",
	"plugin.choria.security.certname_whitelist":                "Patterns of certificate names that are allowed to be clients",
	"plugin.security.provider":                                 "The Security Provider to use",
	"plugin.security.always_overwrite_cache":                   "Always store new Public Keys to the cache overwriting existing ones",
	"plugin.security.support_legacy_certificates":              "Allow certificates without SANs to be used",
	"plugin.choria.security.request_signer.token_file":         "Path to the token used to access a Central Authenticator",
	"plugin.choria.security.request_signer.token_environment":  "Environment variable to store Central Authenticator tokens",
	"plugin.choria.security.request_signing_certificate":       "Path to the public certificate of the key used to sign the JWTs in the Signing Service",
	"plugin.choria.security.request_signer.url":                "URL to the Signing Service",
	"plugin.choria.security.request_signer.service":            "Enables signing requests via Choria RPC requests",
	"plugin.security.client_anon_tls":                          "Use anonymous TLS to the Choria brokers from a client, also disables security provider verification - only when a remote signer is set",
	"plugin.security.file.certificate":                         "When using file security provider, the path to the public certificate",
	"plugin.security.file.key":                                 "When using file security provider, the path to the private key",
	"plugin.security.file.ca":                                  "When using file security provider, the path to the Certificate Authority public certificate",
	"plugin.security.file.cache":                               "When using file security provider, the path to the client cache",
	"plugin.security.certmanager.namespace":                    "When using Cert Manager security provider, the namespace the issuer is in",
	"plugin.security.certmanager.issuer":                       "When using Cert Manager security provider, the name of the issuer",
	"plugin.security.certmanager.replace":                      "when using Cert Manager security provider, replace existing CSRs with new ones",
	"plugin.security.certmanager.alt_names":                    "when using Cert Manager security provider, add these additional names to the CSR",
	"plugin.security.cipher_suites":                            "List of allowed cipher suites",
	"plugin.security.ecc_curves":                               "List of allowed ECC curves",
	"plugin.security.pkcs11.driver_file":                       "When using the pkcs11 security provider, the path to the PCS11 driver file",
	"plugin.security.pkcs11.slot":                              "When using the pkcs11 security provider, the slot to use in the device",
	"plugin.choria.adapters":                                   "The list of Data Adapters to activate",
	"plugin.choria.status_file_path":                           "Path to a JSON file to write server health information to regularly",
	"plugin.choria.status_update_interval":                     "How frequently to write to the status_file_path",
	"plugin.choria.machine.store":                              "Directory where Autonomous Agents are stored",
	"plugin.choria.prometheus_textfile_directory":              "Directory where Prometheus Node Exporter textfile collector reads data",
	"plugin.scout.overrides":                                   "Path to a file holding overrides for Scout checks",
	"plugin.scout.tags":                                        "Path to a file holding tags for a Scout entity",
	"plugin.scout.agent_disabled":                              "Disables the scout agent",
	"plugin.choria.require_client_filter":                      "If a client filter should always be required, only used in Go clients",
	"plugin.choria.services.registry.store":                    "Directory where the Registry service finds DDLs to read",
	"plugin.choria.services.registry.cache":                    "Directory where the Registry client stores DDLs found in the registry",
	"plugin.choria.submission.spool":                           "Path to a directory holding messages to submit to the middleware",
	"plugin.choria.submission.max_spool_size":                  "Maximum amount of messages allowed into each priority",
}
