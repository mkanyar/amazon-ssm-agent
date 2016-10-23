// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// +build windows

// Package appconfig manages the configuration of the agent.

package appconfig

import (
	"os"
	"path/filepath"
)

const (
	// SSM folder path under local app data.
	SSMFolder = "Amazon\\SSM"

	// SSM plugins folder path under local app data.
	SSMPluginFolder = "Amazon\\SSM\\Plugins\\"

	// Exit Code that would trigger a Soft Reboot
	RebootExitCode = 3010

	// List all plugin names, unfortunately golang doesn't support const arrays of strings

	// PluginNameAwsRunScript is the name of the run script plugin
	PluginNameAwsRunScript = "aws:runPowerShellScript"

	// PluginNameAwsPowerShellModule is the name of the PowerShell Module
	PluginNameAwsPowerShellModule = "aws:psModule"

	// PluginNameAwsApplications is the name of the Applications plugin
	PluginNameAwsApplications = "aws:applications"
)

// Program Folder
var DefaultProgramFolder string

// AppConfig Path
var AppConfigPath string

// DefaultDataStorePath represents the directory for storing system data
var DefaultDataStorePath string

// PackageRoot specifies the directory under which packages will be downloaded and installed
var PackageRoot string

// DefaultPluginPath represents the directory for storing plugins in SSM
var DefaultPluginPath string

// DownloadRoot specifies the directory under which files will be downloaded
var DownloadRoot string

// UpdaterArtifactsRoot represents the directory for storing update related information
var UpdaterArtifactsRoot string

// SSMData specifies the directory we used to store SSM data.
var SSMDataPath string

// Windows environment variable %ProgramFiles%
var EnvProgramFiles string

// Windows environment variable %WINDIR%
var EnvWinDir string

// Plugin folder path
var PluginFolder string

func init() {
	/*
		System environment variable "AllUsersProfile" maps to following locations in different locations:

		WindowsServer 2003  -> C:\Documents and Settings\All Users\Application Data
		WindowsServer 2008+ -> C:\ProgramData
	*/

	SSMDataPath = os.Getenv("ProgramData")
	if SSMDataPath == "" {
		SSMDataPath = filepath.Join(os.Getenv("AllUsersProfile"), "Application Data")
	}
	SSMDataPath = filepath.Join(SSMDataPath, SSMFolder)

	EnvProgramFiles = os.Getenv("ProgramFiles")
	EnvWinDir = os.Getenv("WINDIR")
	temp := os.Getenv("TEMP")

	DefaultProgramFolder = filepath.Join(EnvProgramFiles, SSMFolder)
	DefaultPluginPath = filepath.Join(EnvProgramFiles, SSMPluginFolder)
	AppConfigPath = filepath.Join(DefaultProgramFolder, AppConfigFileName)
	DefaultDataStorePath = filepath.Join(SSMDataPath, "InstanceData")
	PackageRoot = filepath.Join(SSMDataPath, "Packages")
	DownloadRoot = filepath.Join(temp, SSMFolder, "Download")
	UpdaterArtifactsRoot = filepath.Join(temp, SSMFolder, "Update")
}
