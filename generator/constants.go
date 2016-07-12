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

const htmlStartTag = `<!doctype html>
<html>`

const htmlEndTag = `</html>`

//TODO: Move JS includes at the end of body
const headerTag = `<head>
<meta http-equiv="X-UA-Compatible" content="IE=9; IE=8; IE=7; IE=EDGE"/>
<title>Gauge Test Results</title>
<link rel="shortcut icon" type="image/x-icon" href="images/favicon.ico">
<link rel="stylesheet" type="text/css" href="css/open-sans.css">
<link rel="stylesheet" type="text/css" href="css/font-awesome.css">
<link rel="stylesheet" type="text/css" href="css/normalize.css"/>
<link rel="stylesheet" type="text/css" href="css/angular-hovercard.css"/>
<link rel="stylesheet" type="text/css" href="css/style.css"/>
<script src="js/lightbox.js"></script>
</head>`

const bodyStartTag = `<body>
`

const bodyEndTag = `</body>`

const bodyHeaderTag = `
<header class="top">
  <div class="header">
    <div class="container">
      <div class="logo"><img src="images/logo.png" alt="Report logo"></div>
      <h2 class="project">Project: {{.ProjectName}}</h2>
    </div>
  </div>
</header>`

const mainStartTag = `<main class="main-container">`

const mainEndTag = `</main>`

const containerStartDiv = `<div class="container">`

const reportOverviewTag = `<div class="report-overview">
  <div class="report_chart">
    <div class="chart">
      <nvd3 options="options" data="data"></nvd3>
    </div>
    <div class="total-specs"><span class="value">{{.TotalSpecs}}</span> <span class="txt">Total specs</span></div>
  </div>
  <div class="report_test-results">
    <ul>
      <li class="fail"><span class="value">{{.Failed}}</span> <span class="txt">Failed</span></li>
      <li class="pass"><span class="value">{{.Passed}}</span> <span class="txt">Passed</span></li>
      <li class="skip"><span class="value">{{.Skipped}}</span> <span class="txt">Skipped</span></li>
    </ul>
  </div>
  <div class="report_details">
    <ul>
      <li>
        <label>Environment </label>
        <span>{{.Env}}</span>
      </li>
      {{if .Tags}}
      <li>
        <label>Tags </label>
        <span>{{.Tags}}</span>
      </li>
      {{end}}
      <li>
        <label>Success Rate </label>
        <span>{{.SuccRate}}%</span>
      </li>
      <li>
        <label>Total Time </label>
        <span>{{.ExecTime}}</span>
      </li>
      <li>
        <label>Generated On </label>
        <span>{{.Timestamp}}</span>
      </li>
    </ul>
  </div>
</div>
`

//TODO: 1. Set first spec as selected by default and load it
//      2. Javascript action to load spec on click
//      3. Filtering based on search query
const sidebarDiv = `{{if not .IsPreHookFailure}}
<aside class="sidebar">
  <h3 class="title">Specifications</h3>

  <div class="searchbar">
    <input id="searchSpecifications" placeholder="Type specification or tag name" type="text"/>
    <i class="fa fa-search"></i>
  </div>

  <div id="listOfSpecifications">
    <ul id="scenarios" class="spec-list">
    {{range $index, $specMeta := .Specs}}
      {{if $specMeta.Failed}} <li class='failed spec-name'>
      {{else if $specMeta.Skipped}} <li class='skipped spec-name'>
      {{else}} <li class='passed spec-name'>
      {{end}}
        <span id="scenarioName" class="scenarioname">{{$specMeta.SpecName}}</span>
        <span id="time" class="time">{{$specMeta.ExecTime}}</span>
      </li>
      {{end}}
    </ul>
  </div>
</aside>
{{end}}`

const specsStartDiv = `<div class="specifications">`

//TODO: Hide if pre/post hook failed
const congratsDiv = `{{if not .Failed}}
  <div class="congratulations details">
    <p>Congratulations! You've gone all <span class="green">green</span> and saved the environment!</p>
  </div>{{end}}`

//TODO 1. Change text on toggle collapse
//     2. Check for collapsible
const hookFailureDiv = `<div class="error-container failed">
  <div collapsable class="error-heading">{{.HookName}} Failed: <span class="error-message">{{.ErrMsg}}</span></div>
  <div class="toggleShow" data-toggle="collapse" data-target="#hookFailureDetails">
    <span>[Show details]</span>
  </div>
  <div class="exception-container" id="hookFailureDetails">
      <div class="exception">
        <pre class="stacktrace">{{.Stacktrace}}</pre>
      </div>
      {{if .Screenshot}}<div class="screenshot-container">
        <a href="data:image/png;base64,{{.Screenshot}}" rel="lightbox">
          <img ng-src="data:image/png;base64,{{.Screenshot}}" class="screenshot-thumbnail"/>
        </a>
      </div> {{end}}
  </div>
</div>`

const specHeaderTag = `<header class="curr-spec">
  <h3 class="spec-head">{{.SpecName}}</h3>
  <span class="time">{{.ExecTime}}</span>
  {{if .Tags}}
  <div class="tags scenario_tags contentSection">
    <strong>Tags:</strong>
    {{range .Tags}}
    <span>{{.}}</span>
    {{end}}
  </div>
  {{end}} 
</header>`

//TODO: Hide this if there is a pre hook failure
const specContainerStartDiv = `<div id="specificationContainer" class="details">`

const endDiv = `</div>`