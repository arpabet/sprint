/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package constants

import (
	"encoding/base64"
	"time"
)

var ExecutableName = "template-server"

var ExecutablePID = ExecutableName + ".pid"

var ExecutableLog = ExecutableName + ".log"

var ApplicationName = "TemplateServer"

var Copyright = "Copyright (C) 2020-present Arpabet, Inc. All rights reserved."

var KeySize = 32  // 256-bit AES key

var Encoding = base64.RawURLEncoding

// ones a week change a key
const KeyRotationDuration = time.Hour * 24 * 7

var statusEndpoint = "https://%s/api/status"

