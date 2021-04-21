## upctl server modify

Modifies the configuration of an existing server

```
upctl server modify <UUID/Title/Hostname...> [flags]
```

### Examples

```
upctl server modify 00bb4617-c592-4b32-b869-35a60b323b18 --plan 1xCPU-1GB
upctl server modify 00bb4617-c592-4b32-b869-35a60b323b18 0053a6f5-e6d1-4b0b-b9dc-b90d0894e8d0 --plan 1xCPU-1GB
upctl server modify my_server1 --plan 1xCPU-2GB
upctl server modify myapp --hostname superapp
```

### Options

```
      --boot-order string               The boot device order, disk / cdrom / network or comma separated combination.
      --cores int                       Number of cores. Sets server plan to custom.
      --hostname string                 Hostname.
      --firewall string                 Enables or disables firewall on the server. You can manage firewall rules with the firewall command.
                                        Available: true, false
      --memory int                      Memory amount in MiB. Sets server plan to custom.
      --metadata string                 Enable metadata service.
      --plan string                     Server plan to use.
      --simple-backup string            Simple backup rule. Format (HHMM,{dailies,weeklies,monthlies}).
                                        Example: 2300,dailies
      --title string                    A short, informational description.
      --time-zone string                Time zone to set the RTC to.
      --video-model string              Video interface model of the server.
                                        Available: vga,cirrus
      --remote-access-enabled string    Enables or disables the remote access.
                                        Available: true, false
      --remote-access-type string       The remote access type.
      --remote-access-password string   The remote access password.
  -h, --help                            help for modify
```

### Options inherited from parent commands

```
  -t, --client-timeout duration   CLI timeout when using interactive mode on some commands (default 1m0s)
      --colours                   Use terminal colours (default true)
      --config string             Config file
  -o, --output string             Output format (supported: json, yaml and human) (default "human")
```

### SEE ALSO

* [upctl server](upctl_server.md)	 - Manage servers

###### Auto generated by spf13/cobra on 21-Apr-2021