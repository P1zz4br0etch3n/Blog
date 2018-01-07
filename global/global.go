/*
    Matrikelnummern: 5836402, 2416160
*/

package global

import "de/vorlesung/projekt/2416160-5836402/models"

var Settings = models.Settings{
	PortNumber:     "4443",
	SessionTimeout: 15,

	PostSuffix:     ".json",

	KeyFile:  "key.pem",
	CertFile: "cert.pem",
}