// Copyright Mia srl
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type metricsResponse struct {
	TotalProcesses int     `json:"total_processes"`
	TotalCPU       float64 `json:"total_cpu"`
	TotalMemory    uint64  `json:"total_memory"`
}

type statusResponse struct {
	Status  string `json:"status"`
	Name    string `json:"name"`
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}

var (
	agentUptime  = ""
	agentVersion = "v1.0.0"
	mutex        sync.Mutex
	metricsData  metricsResponse
)

func init() {
	agentUptime = "just started"
}

// statusRoutes add status routes to router.
func statusRoutes(app *fiber.App, serviceName, serviceVersion string) {
	app.Get("/metrics", MetricsHandler)
	app.Get("/status", StatusHandler)


	app.Get("/-/healthz", func(c *fiber.Ctx) error {
		status := statusResponse{
			Status:  "OK",
			Name:    serviceName,
			Version: serviceVersion,
		}
		return c.JSON(status)
	})

	app.Get("/-/ready", func(c *fiber.Ctx) error {
		status := statusResponse{
			Status:  "OK",
			Name:    serviceName,
			Version: serviceVersion,
		}
		return c.JSON(status)
	})

	app.Get("/-/check-up", func(c *fiber.Ctx) error {
		status := statusResponse{
			Status:  "OK",
			Name:    serviceName,
			Version: serviceVersion,
		}
		return c.JSON(status)
	})
}

func UpdateMetrics(totalProcesses int, totalCPU float64, totalMemory uint64) {
	mutex.Lock()
	defer mutex.Unlock()
	metricsData = metricsResponse{
		TotalProcesses: totalProcesses,
		TotalCPU:       totalCPU,
		TotalMemory:    totalMemory,
	}
}

func StatusHandler(c *fiber.Ctx) error {
	status := statusResponse{
		Status:  "running",
		Uptime:  agentUptime,
		Version: agentVersion,
	}
	return c.JSON(status)
}

func MetricsHandler(c *fiber.Ctx) error {
	mutex.Lock()
	defer mutex.Unlock()
	return c.JSON(metricsData)
}
