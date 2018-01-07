/*
    Matrikelnummern: 5836402, 2416160

	Da die Settings das einzige globale Objekt sind, wurden sie in dieses Package gelegt, um den Aufruf intuitiver zu gegstalten
*/

package global

import "de/vorlesung/projekt/2416160-5836402/models"

var Settings = models.Settings{
	PortNumber:     "4443",
	SessionTimeout: 15,  // in Minuten

	PostSuffix: ".json", // Dateiendung für BlogPost-Dateien

	KeyFile:  "key.pem", // Dateiname von Schlüssel und Zertifikat für TLS Betrieb
	CertFile: "cert.pem",// Diese müssen im gleichen Verzeichnis wie blog.go liegen
}
