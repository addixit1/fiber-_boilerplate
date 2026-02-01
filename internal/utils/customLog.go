package utils

import (
	"fmt"
	"time"
)

// ANSI color codes
const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
	ColorGray    = "\033[90m"

	// Bold colors
	ColorBoldRed     = "\033[1;31m"
	ColorBoldGreen   = "\033[1;32m"
	ColorBoldYellow  = "\033[1;33m"
	ColorBoldBlue    = "\033[1;34m"
	ColorBoldMagenta = "\033[1;35m"
	ColorBoldCyan    = "\033[1;36m"
	ColorBoldWhite   = "\033[1;37m"

	// Background colors
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
)

// LogSuccess logs a success message in green
func LogSuccess(message string) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s[SUCCESS]%s %s%s %s✅ %s%s\n",
		ColorBoldGreen, ColorReset,
		ColorGray, timestamp, ColorReset,
		ColorGreen, message)
}

// LogError logs an error message in red
func LogError(message string) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s[ERROR]%s %s%s %s❌ %s%s\n",
		ColorBoldRed, ColorReset,
		ColorGray, timestamp, ColorReset,
		ColorRed, message)
}

// LogWarning logs a warning message in yellow
func LogWarning(message string) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s[WARNING]%s %s%s %s⚠️  %s%s\n",
		ColorBoldYellow, ColorReset,
		ColorGray, timestamp, ColorReset,
		ColorYellow, message)
}

// LogInfo logs an info message in blue
func LogInfo(message string) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s[INFO]%s %s%s %sℹ️  %s%s\n",
		ColorBoldBlue, ColorReset,
		ColorGray, timestamp, ColorReset,
		ColorBlue, message)
}

// LogDebug logs a debug message in cyan
func LogDebug(message string) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s[DEBUG]%s %s%s %s🔍 %s%s\n",
		ColorBoldCyan, ColorReset,
		ColorGray, timestamp, ColorReset,
		ColorCyan, message)
}

// LogDatabase logs a database message in magenta
func LogDatabase(message string) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s[DATABASE]%s %s%s %s🗄️  %s%s\n",
		ColorBoldMagenta, ColorReset,
		ColorGray, timestamp, ColorReset,
		ColorMagenta, message)
}

// LogServer logs a server message in bold blue
func LogServer(message string) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s[SERVER]%s %s%s %s🚀 %s%s\n",
		ColorBoldBlue, ColorReset,
		ColorGray, timestamp, ColorReset,
		ColorCyan, message)
}

// LogRequest logs an HTTP request in white
func LogRequest(method, path, status string) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")

	// Color status code based on value
	statusColor := ColorGreen
	if status[0] == '4' {
		statusColor = ColorYellow
	} else if status[0] == '5' {
		statusColor = ColorRed
	}

	fmt.Printf("%s[REQUEST]%s %s%s %s%s %s%s%s %s→%s %s%s\n",
		ColorBoldWhite, ColorReset,
		ColorGray, timestamp, ColorReset,
		ColorBoldCyan, method, ColorReset,
		statusColor, status, ColorReset,
		ColorWhite, path)
}

// ColorText returns text with the specified color
func ColorText(text, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, ColorReset)
}

// Custom formatted log with emoji
func LogWithEmoji(emoji, color, label, message string) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s[%s]%s %s%s %s%s %s%s\n",
		color, label, ColorReset,
		ColorGray, timestamp, ColorReset,
		emoji, ColorReset, message)
}

// PrintBanner prints a colored banner
func PrintBanner(text string) {
	banner := fmt.Sprintf(`
%s╔══════════════════════════════════════════════════╗
║  %s%-46s%s  ║
╚══════════════════════════════════════════════════╝%s`,
		ColorBoldCyan, ColorBoldWhite, text, ColorBoldCyan, ColorReset)
	fmt.Println(banner)
}

// LogStartup logs application startup message
func LogStartup(appName, version, port string) {
	fmt.Printf("\n%s", ColorBoldCyan)
	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Printf("║  🚀 %s%-41s%s    ║\n", ColorBoldWhite, appName, ColorBoldCyan)
	fmt.Printf("║  📦 Version: %-36s║\n", version)
	fmt.Printf("║  🌐 Port: %-39s║\n", port)
	fmt.Printf("║  ⏰ Started: %-36s║\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("╚══════════════════════════════════════════════════╝")
	fmt.Printf("%s\n", ColorReset)
}
