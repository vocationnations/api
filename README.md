REST API
============

![Go](https://github.com/vocationnations/api/workflows/Go/badge.svg?branch=master)

# Table of contents

1. [Introduction](#1-introduction)
2. [Installation and Deployment](#2-installation-and-deployment)

## 1. Introduction

This repository serves as the back-end API to the current version of VocationNations.

## 2. Installation and Deployment

**Step 1: Clone and enter the repository**

```
$ git clone git@github.com:vocationnations/api.git
$ cd api
```

**Step 2: Install the dependencies**

```
$ go get
```

**Step 3: Create a .env file**

```shell
$ cat << EOF > .env
DB_PORT=5432
DB_NAME="postgres"
DB_USER="postgres"
DB_PASS="your_pass"
DB_HOST="your.hostname.com"

API_PORT=5000
EOF
```

**Step 4: Build and run the program**

```
$ go build
$ ./api
```

**Step 4: Goto http://0.0.0.0:5000 in your browser**
