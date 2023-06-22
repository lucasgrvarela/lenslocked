package controllers

import (
	"html/template"
	"net/http"

	"github.com/lucasgrvarela/lenslocked/views"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func FAQ(tmpl views.Template) http.HandlerFunc {
	faq := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We offer a free trial for 30 days on any paid plans.",
		},
		{
			Question: "Are you enjoyining?",
			Answer:   "Yes!",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us at <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`,
		},
		{
			Question: "Where is your office located?",
			Answer:   "We are all remote!",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, faq)
	}
}
