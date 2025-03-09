# 🚀 Golang CLI with Cobra

This is a simple CLI application built with [Cobra](https://github.com/spf13/cobra).  
It includes three main commands:

1. **Password Generator CLI**  
   - Generate secure passwords with options for digits and special characters.
   - Save generated passwords for future use.

2. **Weather API CLI**  
   - Fetch weather data for Indonesian cities using [Open-Meteo](https://open-meteo.com/).

---

## 📌 Features
✅ Generate random passwords with customizable settings  
✅ Store and retrieve passwords  
✅ Fetch real-time weather data  
✅ Interactive CLI with loading states  

---

## 🛠 Installation

### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/yourusername/your-cli-project.git
cd your-cli-project
```

### **2️⃣ Install Dependencies**
```sh
go mod tidy
```

### **3️⃣ Build the CLI**
```sh
go build -o cli-app
```

### **4️⃣ Run the CLI**
```sh
./cli-app
```

---

## 🔥 Usage

### **Generate a Password**
```sh
./cli-app password generate -l 10 -d -c
```
**Options:**
- `-l, --length`: Set password length (max 12, default 6).
- `-d, --digits`: Include digits.
- `-c, --special`: Include special characters.

Example Output:
```
Generated Password: 8a@T3!bXz
```

### **Save a Password**
```sh
./cli-app password save
```

---

### **Fetch Weather Data**
```sh
./cli-app weather fetch
```
Example Output:
```
Jakarta: 29°C, Light Rain
```

### **Show Weather for a City**
```sh
./cli-app weather show -c "Jakarta"
```

---

## 🏗 Project Structure
```
📂 your-cli-project
 ┣ 📂 cmd/        # CLI commands
 ┃ ┣ 📜 root.go   # Main CLI entry
 ┃ ┣ 📜 password.go   # Password generator commands
 ┃ ┣ 📜 weather.go    # Weather API commands
 ┣ 📂 pkg/
 ┃ ┣ 📜 cities.go   # City data & methods
 ┣ 📜 main.go       # Entry point
 ┣ 📜 go.mod        # Go module file
 ┣ 📜 README.md     # Documentation
```

---

## 🤝 Contributing
Contributions are welcome!  
1. **Fork** the project  
2. **Create a branch** (`feature/new-feature`)  
3. **Commit your changes**  
4. **Push** to GitHub  
5. **Open a PR** 🎉  

---

## ⚖️ License
This project is licensed under the **MIT License**.  
Feel free to use and modify it!

---

🚀 **Happy coding!**

