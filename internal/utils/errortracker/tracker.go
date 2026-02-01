package errortracker

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// Layer represents different application layers
type Layer string

const (
	LayerRoute      Layer = "ROUTE"
	LayerController Layer = "CONTROLLER"
	LayerService    Layer = "SERVICE"
	LayerRepository Layer = "REPOSITORY"
	LayerDTO        Layer = "DTO"
	LayerMiddleware Layer = "MIDDLEWARE"
	LayerDatabase   Layer = "DATABASE"
	LayerExternal   Layer = "EXTERNAL"
)

// ErrorContext holds error tracking information
type ErrorContext struct {
	Layer     Layer
	Function  string
	File      string
	Line      int
	Timestamp time.Time
	Message   string
	Original  error
}

// Track logs an error with context information
func Track(layer Layer, message string, err error) error {
	// Get caller information
	pc, file, line, ok := runtime.Caller(1)
	funcName := "unknown"
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			funcName = fn.Name()
			// Extract just the function name, not full path
			parts := strings.Split(funcName, "/")
			if len(parts) > 0 {
				funcName = parts[len(parts)-1]
			}
		}
		// Extract just the file name
		fileParts := strings.Split(file, "/")
		if len(fileParts) > 0 {
			file = fileParts[len(fileParts)-1]
		}
	}

	ctx := ErrorContext{
		Layer:     layer,
		Function:  funcName,
		File:      file,
		Line:      line,
		Timestamp: time.Now(),
		Message:   message,
		Original:  err,
	}

	// Print formatted error log
	printError(ctx)

	return err
}

// printError formats and prints error information
func printError(ctx ErrorContext) {
	color := getLayerColor(ctx.Layer)
	reset := "\033[0m"

	fmt.Printf("\n%s╔════════════════════════════════════════════════════════════════╗%s\n", color, reset)
	fmt.Printf("%s║ ❌ ERROR TRACKER                                               ║%s\n", color, reset)
	fmt.Printf("%s╠════════════════════════════════════════════════════════════════╣%s\n", color, reset)
	fmt.Printf("%s║ Layer:    %-52s ║%s\n", color, ctx.Layer, reset)
	fmt.Printf("%s║ Function: %-52s ║%s\n", color, ctx.Function, reset)
	fmt.Printf("%s║ Location: %-52s ║%s\n", color, fmt.Sprintf("%s:%d", ctx.File, ctx.Line), reset)
	fmt.Printf("%s║ Time:     %-52s ║%s\n", color, ctx.Timestamp.Format("2006-01-02 15:04:05"), reset)
	fmt.Printf("%s╠════════════════════════════════════════════════════════════════╣%s\n", color, reset)

	// Message (word wrap if needed)
	if ctx.Message != "" {
		fmt.Printf("%s║ Message:                                                       ║%s\n", color, reset)
		fmt.Printf("%s║ %s%-62s%s ║%s\n", color, "\033[1m", truncate(ctx.Message, 62), "\033[0m"+color, reset)
	}

	// Original error
	if ctx.Original != nil {
		fmt.Printf("%s║ Error:                                                         ║%s\n", color, reset)
		fmt.Printf("%s║ %-62s ║%s\n", color, truncate(ctx.Original.Error(), 62), reset)
	}

	fmt.Printf("%s╚════════════════════════════════════════════════════════════════╝%s\n\n", color, reset)
}

// getLayerColor returns ANSI color code for each layer
func getLayerColor(layer Layer) string {
	switch layer {
	case LayerRoute:
		return "\033[35m" // Magenta
	case LayerController:
		return "\033[36m" // Cyan
	case LayerService:
		return "\033[33m" // Yellow
	case LayerRepository:
		return "\033[34m" // Blue
	case LayerDTO:
		return "\033[32m" // Green
	case LayerMiddleware:
		return "\033[95m" // Light Magenta
	case LayerDatabase:
		return "\033[31m" // Red
	case LayerExternal:
		return "\033[37m" // White
	default:
		return "\033[0m" // Reset
	}
}

// truncate truncates a string to maxLen characters
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// TrackWithDetails logs error with additional details
func TrackWithDetails(layer Layer, message string, err error, details map[string]interface{}) error {
	Track(layer, message, err)

	if len(details) > 0 {
		fmt.Println("Additional Details:")
		for k, v := range details {
			fmt.Printf("  • %s: %v\n", k, v)
		}
		fmt.Println()
	}

	return err
}
