/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"flag"
	"log"
	"github.com/pkg/errors"
	"de/vorlesung/projekt/2416160-5836402/global"
)

var settingsLoaded = false

func LoadSettings() error {
	if settingsLoaded {
		log.Println("Settings already loaded.")
		return nil
	}

	e := readSettingsFile()
	if e != nil {
		log.Println(e)
		return e
	}
	parseCommandLineArgs()

	settingsLoaded = true
	return nil
}

func ForceLoadSettingsFile() error {
	return readSettingsFile()
}

func SaveSettings() error {
	err := WriteJsonFile(SettingsDir,"settings", global.Settings)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func readSettingsFile() error {
	err := ReadJsonFile(SettingsDir, "settings", &global.Settings)
	if err != nil {
		log.Println("Could not read Settings-File. Creating new file with default values...")
		e := SaveSettings()
		if e != nil {
			return errors.New("Could not create default Settings-File.")
		}
	}
	log.Println("Reading Settings-File done.")
	return nil
}

func parseCommandLineArgs() {
	var clPortNumber = flag.String("port", global.Settings.PortNumber, "")
	var clSessionTimeOut = flag.Int("timeout", int(global.Settings.SessionTimeout), "")
	var clPostSuffix = flag.String("postSufx", global.Settings.PostSuffix, "")

	var clKeyFile = flag.String("keyFile", global.Settings.KeyFile, "")
	var clCertificateFile = flag.String("certFile", global.Settings.CertFile, "")

	flag.Parse()

	if *clPortNumber != global.Settings.PortNumber {
		global.Settings.PortNumber = *clPortNumber
	}
	if *clSessionTimeOut != int(global.Settings.SessionTimeout) {
		global.Settings.SessionTimeout = uint(*clSessionTimeOut)
	}
	if *clPostSuffix != global.Settings.PostSuffix {
		global.Settings.PostSuffix = *clPostSuffix
	}
	if *clKeyFile != global.Settings.KeyFile {
		global.Settings.KeyFile = *clKeyFile
	}
	if *clCertificateFile != global.Settings.CertFile {
		global.Settings.CertFile = *clCertificateFile
	}
}
