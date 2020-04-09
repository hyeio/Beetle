// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/clivern/beetle/internal/app/kubernetes"
	"github.com/clivern/beetle/internal/app/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Namespaces controller
func Namespaces(c *gin.Context) {
	cn := c.Param("cn")
	result := []model.Namespace{}

	clusters, err := kubernetes.GetClusters()

	if err != nil {
		log.WithFields(log.Fields{
			"CorrelationId": c.Request.Header.Get("X-Correlation-ID"),
		}).Error(fmt.Sprintf(`Error: %s`, err.Error()))

		c.Status(http.StatusInternalServerError)
		return
	}

	for _, cluster := range clusters {
		if cn != cluster.Name {
			continue
		}

		result, err = cluster.GetNamespaces(context.Background())

		if err != nil {
			log.WithFields(log.Fields{
				"CorrelationId": c.Request.Header.Get("X-Correlation-ID"),
			}).Error(fmt.Sprintf(`Error: %s`, err.Error()))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"namespaces": result,
	})
}
