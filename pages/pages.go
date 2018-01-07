/*
    Matrikelnummern: 5836402, 2416160
*/

package pages

import (
	"de/vorlesung/projekt/2416160-5836402/models"
)

type IndexPage struct {
	UserLoggedIn    bool
	ShowArchiveLink bool
	UserName        string
	Posts           []models.BlogPost
}