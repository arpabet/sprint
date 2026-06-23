/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintapp

import (
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/framework/pkg/assets"
	"go.arpabet.com/sprint/framework/pkg/assetsgz"
	"go.arpabet.com/sprint/framework/pkg/resources"
	"os"
)

type FileModes map[string]interface{}

var DefaultFileModes = FileModes {
	"log.dir": os.FileMode(0775),
	"log.file": os.FileMode(0664),
	"backup.file": os.FileMode(0664),
	"exe.file": os.FileMode(0775),
	"run.dir": os.FileMode(0775),
	"pid.file": os.FileMode(0666),
	"data.dir": os.FileMode(0770),
	"data.file": os.FileMode(0664),
}

var DefaultResources = &glue.ResourceSource{
	Name: "resources",
	AssetNames: resources.AssetNames(),
	AssetFiles: resources.AssetFile(),
}

var DefaultAssets = &glue.ResourceSource{
	Name: "assets",
	AssetNames: assets.AssetNames(),
	AssetFiles: assets.AssetFile(),
}

var DefaultGzipAssets = &glue.ResourceSource{
	Name: "assets-gzip",
	AssetNames: assetsgz.AssetNames(),
	AssetFiles: assetsgz.AssetFile(),
}

var ApplicationBeans = []interface{} {
	ApplicationFlags(100000), // override any property resolvers
	FlagSetFactory(),
	ResourceService(),
}

