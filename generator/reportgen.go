// Copyright 2015 ThoughtWorks, Inc.

// This file is part of getgauge/html-report.

// getgauge/html-report is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// getgauge/html-report is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with getgauge/html-report.  If not, see <http://www.gnu.org/licenses/>.
package generator

import (
	"io"
	"log"
	"text/template"

	"github.com/getgauge/html-report/gauge_messages"
)

type overview struct {
	ProjectName string
	Env         string
	Tags        string
	SuccRate    float32
	ExecTime    string
	Timestamp   string
	TotalSpecs  int
	Failed      int
	Passed      int
	Skipped     int
}

type specsMeta struct {
	SpecName string
	ExecTime string
	Failed   bool
	Skipped  bool
	Tags     []string
}

type sidebar struct {
	IsPreHookFailure bool
	Specs            []*specsMeta
}

type hookFailure struct {
	HookName   string
	ErrMsg     string
	Screenshot string
	Stacktrace string
}

type specHeader struct {
	SpecName string
	ExecTime string
	FileName string
	Tags     []string
}

type scenario struct {
	Heading  string
	ExecTime string
	Tags     []string
	Res      result
}

type spec struct {
	CommentsBeforeTable []string
	Table               *table
	CommentsAfterTable  []string
	Scenarios           []*scenario
}

type table struct {
	Headers []string
	Rows    []*row
}

type row struct {
	Cells []string
	Res   result
}

type result int

const (
	PASS result = iota
	FAIL
	SKIP
)

func gen(tmplName string, w io.Writer, data interface{}) {
	tmpl, err := template.New("Reports").Parse(tmplName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func generate(suiteRes *gauge_messages.ProtoSuiteResult, w io.Writer) {
	overview := toOverview(suiteRes)
	sidebar := toSidebar(suiteRes)
	specHeader := toSpecHeader(suiteRes.GetSpecResults()[0])
	spec := toSpec(suiteRes.GetSpecResults()[0])

	gen(htmlStartTag, w, nil)
	gen(pageHeaderTag, w, nil)
	gen(bodyStartTag, w, nil)
	gen(bodyHeaderTag, w, overview)
	gen(mainStartTag, w, nil)
	gen(containerStartDiv, w, nil)
	gen(reportOverviewTag, w, overview)
	gen(specsStartDiv, w, nil)
	gen(sidebarDiv, w, sidebar)
	gen(specContainerStartDiv, w, nil)
	gen(specHeaderStartTag, w, specHeader)
	gen(tagsDiv, w, specHeader)
	gen(headerEndTag, w, nil)
	gen(specsItemsContainerDiv, w, nil)
	gen(specCommentsAndTableTag, w, spec)
	gen(scenarioContainerStartDiv, w, spec.Scenarios[0])
	gen(scenarioHeaderStartDiv, w, spec.Scenarios[0])
	gen(endDiv, w, nil)
	gen(endDiv, w, nil)
	gen(endDiv, w, nil)
	gen(endDiv, w, nil)
	gen(endDiv, w, nil)
	gen(endDiv, w, nil)
	gen(endDiv, w, nil)
	gen(mainEndTag, w, nil)
	gen(bodyFooterTag, w, nil)
	gen(bodyEndTag, w, nil)
	gen(htmlEndTag, w, nil)
}
