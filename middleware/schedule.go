package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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
	pollInterval    = 5 * time.Second
	client          = &http.Client{Timeout: 5 * time.Second}
)

// StartSchedulePoller starts a background goroutine that periodically fetches schedule
func StartSchedulePoller() {
	base := config.AppConfig.BaseURLSchedule
	if base == "" {
		log.Println("schedule: BASE_URL_SCHEDULE not configured; poller disabled")
		return
	}

	log.Println("schedule: starting poller, base=", base)
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
	log.Println("schedule: fetching", url)
	resp, err := client.Get(url)
	if err != nil {
		log.Println("schedule: fetch error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("schedule: non-200 response:", resp.Status)
		return
	}

	var s ScheduleResponse
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		log.Println("schedule: decode error:", err)
		return
	}

	mu.Lock()
	currentSchedule = &s
	mu.Unlock()
	log.Println("schedule: updated: active=", s.Active, "path=", s.Path, "mode=", s.Mode, "start=", s.Start, "end=", s.End, "now=", s.Now)
}

// RequireScheduleUpload is a Gin middleware that checks current schedule and blocks upload route when inactive
// RequireScheduleUpload accepts optional path(s). If provided, middleware will
// compare request path against provided path(s); otherwise it will use schedule.Path.
func RequireScheduleUpload(paths ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Only check for POST /upload (but middleware may be attached only to that route)
		mu.RLock()
		s := currentSchedule
		mu.RUnlock()

		// If schedule not available, default to disabled (safe default)

		if s == nil {
			log.Println("schedule: no schedule loaded; rejecting request - Schedule unavailable")
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"message": "Schedule unavailable"})
			return
		}

		reqPath := ctx.Request.URL.Path
		log.Println("schedule: request path=", reqPath, "schedule path=", s.Path, "active=", s.Active)

		// determine which paths we should match against: provided paths override schedule.Path
		var matchPaths []string
		if len(paths) > 0 {
			matchPaths = paths
		} else if s.Path != "" {
			matchPaths = []string{s.Path}
		}

		if len(matchPaths) > 0 {
			matched := false
			for _, p := range matchPaths {
				if p == reqPath || p == ctx.FullPath() || strings.HasSuffix(reqPath, p) || strings.Contains(reqPath, p) {
					matched = true
					break
				}
			}
			if !matched {
				log.Println("schedule: path not matched; allowing request")
				ctx.Next()
				return
			}
		}

		if s.Active {
			log.Println("schedule: active=true; allowing request")
			ctx.Next()
			return
		}

		log.Println("schedule: active=false; blocking request")
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Upload disabled by schedule"})
	}
}
