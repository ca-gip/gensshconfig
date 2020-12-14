# gen-ssh-config

Generate a *config* file based on `Ansible` inventories

## Getting Started

### Intallation

```bash
curl -ssf -L https://raw.githubusercontent.com/ca-gip/gensshconfig/master/install.sh | sudo sh
```

### (Alternative) Build the binary

You will need go v1.15

```bash
go build -o gensshconfig main.go
```

## Usage

### Ansible Inventory Layout

Generation works for project that use the following inventory layout
```bash
inventories
├── common_vars
│   ├── all.yml
│   └── all_vault.yml
└── os
    ├── client_one_hp
    │   ├── group_vars
    │   └── hosts.ini
    ├── client_one_prd
    │   ├── group_vars
    │   └── hosts.ini
    ├── client_three_hp
    │   ├── group_vars
    │   └── hosts.ini
    ├── client_three_prd
    │   ├── group_vars
    │   └── hosts.ini
    ├── client_two_hp
    │   ├── group_vars
    │   └── hosts.ini
    └── client_two_prd
        ├── group_vars
        └── hosts.ini
```

### Example

#### You can check what the config will look like with

```bash
gensshconfig --user $USER_BASTION --bastion $IP_BASTION --inventory ~/Project/ansible-kube/inventories/os
```

#### Alternatively save it 

```bash
gensshconfig --user $USER_BASTION --bastion $IP_BASTION --inventory ~/Project/ansible-kube/inventories/os > ~/.ssh/k8s_config
```

If you do not want to overwrite the default config like in the above exmaple, do not forget to include the following line in your main `config` file
```config
Include k8s_config
```

### Help

To see the available parameters
```
gensshconfig -h
```

### Available options

| Name             | Description                                      | Mandatory    | Type      |
| :----------------| :---------------------------------------------- | :---------:  | :-------: |
|  **bastion**     |  *IP Address of the bastion*                     | `yes`        | `String`  |
|  **inventory**   |  *Inventory folder to extract clusters host*     | `yes`        | `String`  |
|  **user**        |  *User for the bastion*                          | `yes`        | `String`  |


