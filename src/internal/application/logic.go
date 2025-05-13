package application

import (
	"bytes"
	"iter"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/chmikata/incident-notification/internal/incident"
)

type Code string

const (
	Detection     Code = "検出"
	Investigation Code = "調査中"
	Correction    Code = "修正中"
	Resolved      Code = "解決"
	Completed     Code = "完了"
)

type IncidentNotificatin struct {
	collection   *incident.Collection
	notification *incident.Notificaton
}

func NewIncidentNotificatin(apiKey string, webhookUrl string) *IncidentNotificatin {
	return &IncidentNotificatin{
		collection:   incident.NewCollection(apiKey),
		notification: incident.NewNotificator(webhookUrl),
	}
}

func (i *IncidentNotificatin) Do(tpl string) error {

	incidents, err := i.collection.GetIncidents()
	if err != nil {
		panic(err)
	}

	content, err := i.convert(tpl, incidents)
	if err != nil {
		panic(err)
	}

	err = i.notification.SendNotification(content)
	if err != nil {
		panic(err)
	}
	return nil
}

func (i *IncidentNotificatin) convert(tpl string, incidents []incident.Incident) (string, error) {
	output, err := template.New("incident").Funcs(sprig.FuncMap()).Parse(tpl)
	if err != nil {
		return "", err
	}

	var targets []incident.Incident
	for inc := range selectIncident(incidents) {
		targets = append(targets, inc)
	}

	buff := new(bytes.Buffer)
	if err := output.ExecuteTemplate(buff, "incident", targets); err != nil {
		return "", err
	}
	return buff.String(), nil
}

func selectIncident(incidents []incident.Incident) iter.Seq[incident.Incident] {
	return func(yeild func(incident.Incident) bool) {
		for _, inc := range incidents {
			if inc.Status != string(Completed) && !inc.Experimental {
				yeild(inc)
			}
		}
	}
}
