# Slurp
Run the website locally by running `go run main.go` and visiting [localhost:8080](localhost:8080).

[Check it out live!](http://ec2-13-57-214-164.us-west-1.compute.amazonaws.com:8080/)

# Table of Contents
- [To Do](#todo)
- [How to Install](#how-to-install)
- [Tech Stack](#tech-stack)
- [Server](#server)
- [Management Tools](#management-tools)
- [Contributors](#contributors)

## TODO
- purchase domain name, point to ec2 instance
- create terraform for server setup
- set up https
- set up redirect :80 to :8080

## How To Install
1. [Install Go](https://go.dev/dl/)
2. [Install nvm-windows](https://github.com/coreybutler/nvm-windows?tab=readme-ov-file#install-nvm-windows)
3. [Install nvm-setup.exe](https://github.com/coreybutler/nvm-windows/releases)
4. Install other dependencies 
```console
# install node.js
nvm install 22.5.1
# after install, reinstall global utilities. 
nvm use 22.5.1
npm --version
npm install -g yarn
```
5. Initialize your go project by running the following in your bash terminal:
```console
# initialize your go project
go mod init myproject
```
6. Install Tailwindcss by running the following in your bash terminal:
```console 
npm install -D tailwindcss
```

7. Set up Windows Defender Firewall to allow connections to localhost port 8080. [Follow this stack overflow answer.](https://stackoverflow.com/a/65393403)
8. Run your go website by running the following in your bash terminal:
```console
go run main.go
```

9. Visit your webpage!
Open a web browser and go to [http://localhost:8080](http://localhost:8080).


## Tech Stack
### [Go](https://go.dev/doc/)
What: Backend language to build web servers and APIs.
Why: It's known for its performance, simplicity, and strong concurrency support.



### [Temple](https://github.com/docwhat/temple)
What: Templating engine for Go.
Why: It simplifies the process of rendering HTML pages dynamically.

### [Tailwind CSS](https://tailwindcss.com/docs/installation)
What: Utility-first CSS framework for styling.
Why: It allows rapid UI development with a consistent design language.

### [Htmx](https://htmx.org/docs/#introduction)
What: Frontend library for making HTML hyperactive.
Why: It allows you to add AJAX, CSS transitions, WebSockets, and more, directly in HTML attributes.


## Server
The website is hosted on an Amazon Web Service (AWS) EC2 instance. 

ip: 13.57.214.164

dns: ec2-13-57-214-164.us-west-1.compute.amazonaws.com

### Setup
- Connect to the ec2 instance
- Install go
```console
sudo yum install golang -y
``` 
- By default, EC2 instances block incoming traffic on non-standard ports. You need to open port 8080 in your instance’s Security Group:
    - Go to the AWS EC2 Console.
    - Navigate to Instances, select your instance.
    - Under Security, find the Security Group associated with your instance.
    - Click Edit inbound rules.
    - Add a new rule:
        - Type: Custom TCP
        - Protocol: TCP
        - Port Range: 8080
        - Source: Choose 0.0.0.0/0 (public access) or restrict it to your IP if needed.
    - Save the changes.

## Management Tools
- Github: For version control and collaboration
- Discord: For communication and coordination

## Contributors
- Kelley Danger
- Sergio Sartini