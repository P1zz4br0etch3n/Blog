/*
    Matrikelnummern: 5836402, 2416160
*/

package pages

import "de/vorlesung/projekt/2416160-5836402/models"

type BlogPage struct {
	Data *models.Blog
}

type IndexPage struct {
	Title string
	UserLoggedIn bool
	UserName string
}