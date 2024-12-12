package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

var helmetConfig = helmet.Config{
	XSSProtection:             "0",
	ContentTypeNosniff:        "nosniff",
	XFrameOptions:             "SAMEORIGIN",
	ReferrerPolicy:            "no-referrer",
	CrossOriginEmbedderPolicy: "require-corp",
	CrossOriginOpenerPolicy:   "same-origin",
	CrossOriginResourcePolicy: "same-origin",
	OriginAgentCluster:        "?1",
	XDNSPrefetchControl:       "off",
	XDownloadOptions:          "noopen",
	XPermittedCrossDomain:     "none",
}

func Helmet() fiber.Handler {
	return helmet.New(helmetConfig)
}
