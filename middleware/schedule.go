package middleware

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/HMTCITS/hmtc-backend-2025/config"
	"github.com/gin-gonic/gin"
)

// ScheduleResponse represents the expected response from schedule API
type ScheduleResponse struct {
	Timezone   string    `json:"timezone"`
	Mode       string    `json:"mode"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	Path       string    `json:"path"`
	Now        time.Time `json:"now"`
	Active     bool      `json:"active"`
	NextChange time.Time `json:"nextChange"`
	Source     string    `json:"source"`
}

var (
	mu              sync.RWMutex
	currentSchedule *ScheduleResponse
	pollInterval    = 30 * time.Second
	client          = &http.Client{Timeout: 5 * time.Second}
)

// StartSchedulePoller starts a background goroutine that periodically fetches schedule
func StartSchedulePoller() {
	base := config.AppConfig.BaseURLSchedule
	if base == "" {
		// no schedule configured
		return
	}

	// initial fetch
	fetchSchedule(base)

	go func() {
		ticker := time.NewTicker(pollInterval)
		defer ticker.Stop()
		for range ticker.C {
			fetchSchedule(base)
		}
	}()
}

func fetchSchedule(base string) {
	url := base + "/api/schedule"
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	var s ScheduleResponse
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return
	}

	mu.Lock()
	currentSchedule = &s
	mu.Unlock()
}

func boolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// RequireScheduleUpload is a Gin middleware that checks current schedule and blocks upload route when inactive
func RequireScheduleUpload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Only check for POST /upload (but middleware may be attached only to that route)
		mu.RLock()
		s := currentSchedule
		mu.RUnlock()

		// If schedule not available, default to disabled (safe default)
		if s == nil {
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"message": "Schedule unavailable"})
			return
		}

		// if path is set and does not match, allow
		if s.Path != "" && s.Path != ctx.FullPath() {
			ctx.Next()
			return
		}

		if s.Active {
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Upload disabled by schedule"})
	}
}
