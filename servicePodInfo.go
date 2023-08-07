package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

const (
	defaultServiceName = "defaultService"
	defaultVersion     = "0.0.0.0000"
	defaultTimeZone    = "Asia/Shanghai"
	defaultPort        = "32262"
)

var (
	startTime   = time.Now()
	serviceName = getEnv("SERVICE_NAME", defaultServiceName)
	serviceVer  = getEnv("SERVICE_VERSION", defaultVersion)
	timeZone    = getEnv("TIME_ZONE", defaultTimeZone)
	port        = getEnv("PORT", defaultPort)
	podName     = getEnv("HOSTNAME", "")
)

type ResponseData struct {
	StartTimeUTC string `json:"start_time_utc"`
	StartTimeCST string `json:"start_time_cst"`
	Duration     string `json:"duration"`
	ServiceName  string `json:"service_name"`
	ServiceVer   string `json:"service_version"`
	PodName      string `json:"pod_name"`
	ResponseCST  string `json:"response_cst"`
	PodIP        string `json:"pod_ip"`
	GoVersion    string `json:"go_version"`
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("环境变量 %s 未设置，使用默认值: %s\n", key, defaultValue)
		return defaultValue
	}
	return value
}

func logRequest(r *http.Request) {
	log.Printf("Received request from %s for route %s", r.RemoteAddr, r.URL.Path)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		duration := time.Since(startTime).String()
		startTimeUTC := startTime.UTC().Format(time.RFC1123Z)
		location, err := time.LoadLocation(timeZone)
		if err != nil {
			log.Printf("无法加载时区：%v，使用默认时区", err)
			location = time.UTC
		}
		startTimeCST := startTime.In(location).Format(time.RFC1123Z)
		responseCST := time.Now().In(location).Format(time.RFC1123Z)
		podIP := os.Getenv("POD_IP")

		respData := ResponseData{
			StartTimeUTC: startTimeUTC,
			StartTimeCST: startTimeCST,
			Duration:     duration,
			ServiceName:  serviceName,
			ServiceVer:   serviceVer,
			PodName:      podName,
			ResponseCST:  responseCST,
			PodIP:        podIP,
			GoVersion:    runtime.Version(),
		}

		jsonResp, err := json.Marshal(respData)
		if err != nil {
			log.Printf("Error while marshaling response data to JSON: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)

		logRequest(r)
	}
}

func main() {
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("无法打开日志文件：%v", err)
	}
	defer logFile.Close()

	logWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(logWriter)

	log.Printf("%s-%s 正常启动，并准备好提供服务", serviceName, serviceVer)

	http.HandleFunc("/runinfo", infoHandler)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("服务器启动错误：%v", err)
	}
}
