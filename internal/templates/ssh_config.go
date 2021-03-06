package templates

const SshConfig = `# Auto generated by gensshconfig
{{ $bastion := .Bastion.Hostname }}
# Bastion
Host {{ $bastion }}
  ForwardAgent yes
  Hostname  {{ .Bastion.Addr }}
  User {{ .Username }}
  ControlPath ~/.ssh/cm-%r@%h:%p
  ControlMaster auto
  ControlPersist 10m
{{ range .Clusters }}
# {{ .Name }}
{{ range .Hosts }}
Host {{ .Hostname }}
  ProxyJump {{ $bastion }}
  Hostname {{ .Addr }}
  User root
{{ end }}
{{ end }}`
