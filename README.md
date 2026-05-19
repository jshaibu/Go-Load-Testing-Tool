### Go Load Testing Tool
A generic load testing tool for simulating high-concurrency API usage and measuring backend performance. Designed to be reusable and extensible for any API or backend service.

Installation
1. Install Go
2. Download and install Go (example for Linux):


# Download Go (replace with the latest version if needed)wget https://go.dev/dl/go1.26.3.linux-amd64.tar.gz# Extract to /usr/localsudo tar -C /usr/local -xzf ~/Downloads/go1.26.3.linux-amd64.tar.gz# Add Go to your PATH (add these lines to ~/.bashrc or ~/.profile)export PATH=$PATH:/usr/local/go/bin# Reload your shell configurationsource ~/.bashrc# Verify installationgo version

# Download Go (replace with the latest version if needed)
wget https://go.dev/dl/go1.26.3.linux-amd64.tar.gz

# Extract to /usr/local
sudo tar -C /usr/local -xzf ~/Downloads/go1.26.3.linux-amd64.tar.gz

# Add Go to your PATH (add these lines to ~/.bashrc or ~/.profile)
export PATH=$PATH:/usr/local/go/bin

# Reload your shell configuration
source ~/.bashrc

# Verify installation
go version

git clone <your-repo-url>
cd load-test

go run main.go

go build -o load-test
./load-test

# Key Features
1. Concurrent worker simulation
2. Real-time metrics dashboard
3. Configurable API endpoints and job counts
4. Thread-safe metrics collection


# Usage & Configuration
1. Changing API URLs and Credentials
2. The API endpoints and credentials are defined at the top of main.go

```go
const (
    BaseURL      = "http://localhost:5115"
    LoginURL     = BaseURL + "/api/Auth/login"
    ProductsURL  = BaseURL + "/api/Products?page=1&pageSize=200&includeInactive=false"
    OrdersURL    = BaseURL + "/api/Orders"

    Email    = "user@test.com"
    Password = "password"
)
```

# To use this tool with your own API:

1. Open main.go.
2. Change BaseURL to your server’s address.
3. Update the endpoint paths (LoginURL, ProductsURL, OrdersURL) as needed for your API.
4. Set your own Email and Password for authentication.

# Running the Tool
After configuring the URLs and credentials, run:

```go
go run main.go
// or build and run the binary:
go build -o load-test
./load-test
```

# Vision
The goal is to evolve this project into a reusable library for stress testing and benchmarking any API or backend service.

### Contributing
Contributions are welcome! If you have ideas, improvements, or bug fixes, feel free to open an issue or submit a pull request.

### Happy coding!

