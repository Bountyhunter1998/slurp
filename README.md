# Slurp

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

### Frontend
- HTMX
- Tailwind CSS
### Backend
- Go

## Management Tools
- Github: For version control and collaboration
- Discord: For communication and coordination

## How To Install
1. [Install Go](https://go.dev/dl/)
2. [Install nvm-windows](https://github.com/coreybutler/nvm-windows?tab=readme-ov-file#install-nvm-windows)
- [Install nvm-setup.exe](https://github.com/coreybutler/nvm-windows/releases)
- ```console
# install node.js
nvm install 22.5.1
# after install, reinstall global utilities. 
nvm use 22.5.1
npm --version
npm install -g yarn
```
3. Initialize your go project by running the following in your bash terminal:
- ```console
# initialize your go project
go mod init myproject

```
4. Install Tailwindcss by running the following in your bash terminal:
- `npm install -D tailwindcss`

5. Set up Windows Defender Firewall to allow connections to localhost port 8080. [Follow this stack overflow answer.](https://stackoverflow.com/a/65393403)
5. Run your go website by running the following in your bash terminal:
- `go run main.go`

6. Visit your webpage!
- Open a web browser and go to [http://localhost:8080](http://localhost:8080)

## Contributors
- Kelley Danger
- Sergio Sartini