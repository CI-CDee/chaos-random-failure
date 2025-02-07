
## 📂 Project Structure  

```plaintext
chaos-engineering-go/
│── main.go              # Entry point, starts the server
│── chaos/              
│   ├── middleware.go    # Chaos middleware (latency, errors, crashes)
│── config/             
│   ├── config.go        # Reads failure rates from env vars
│── handlers/           
│   ├── handler.go       # Simple HTTP handler
│── go.mod              # Go module file
│── README.md           # Project documentation
│── .env                # Environment variables for failure rates
```

## ⚡ Features  

✅ **Random Latency** - Adds artificial delay to requests.  
✅ **Error Injection** - Randomly returns `500 Internal Server Error`.  
✅ **Process Crash Simulation** - Randomly kills the server.  
✅ **Configurable via `.env`** - Customize chaos settings.  

## 🛠️ Setup & Run  

### Clone the Repository  
```sh
git clone https://github.com/yourusername/chaos-engineering-go.git
cd chaos-engineering-go
```

### Install Dependencies  
```sh
go mod tidy
```

### Configure Chaos Variables  
Create a `.env` file (or use system environment variables):  
```plaintext
CHAOS_LATENCY_MS=200      # Delay in milliseconds
CHAOS_ERROR_RATE=0.3      # 30% chance of HTTP 500 errors
CHAOS_CRASH_RATE=0.1      # 10% chance of process crash
PORT=8080                 # Server port
```

### Run the Server  
```sh
go run main.go
```

### Test the Chaos  
```sh
curl -v http://localhost:8080
```
Expect **random failures**, **delayed responses**, or even the **server crashing**!

---

💡 **Inspired by Chaos Monkey!** Use this to **test system resilience** under failure conditions.  
Made with ❤️ in Go. 🚀  

