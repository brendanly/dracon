package main

import (
	v1 "github.com/thought-machine/dracon/api/proto/v1"
	types "github.com/thought-machine/dracon/producers/npm_audit/types/npmaudit-issue"

	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOut(t *testing.T) {
	var results types.NpmAuditOut
	json.Unmarshal([]byte(exampleOutput), &results)

	issues := parseOut(&results)
	expectedIssue := &v1.Issue{
		Target:      "foo",
		Type:        "Vulnerable Dependency",
		Title:       "Denial of Service",
		Severity:    v1.Severity_SEVERITY_LOW,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_HIGH,
		Description: "Vulnerable Versions: <1.2.3\nRecommendation: Upgrade to version 1.2.3 or later.\nOverview: Versions of 'foo' prior to 1.2.3 are vulnerable to Denial of Service (DoS).  The package fails to remove old values from the cache even after a value passes its 'maxAge' property. This may allow attackers to exhaust the system's memory if they are able to abuse the application logging.\nReferences: - [Snyk Report](https://snyk.io/vuln/npm:foo:20180117)\nNPM Advisory URL: https://npmjs.com/advisories/123\n",
	}

	assert.Equal(t, issues[0].Target, expectedIssue.Target)
	assert.Equal(t, issues[0].Type, expectedIssue.Type)
	assert.Equal(t, issues[0].Title, expectedIssue.Title)
	assert.Equal(t, issues[0].Severity, expectedIssue.Severity)
	assert.Equal(t, issues[0].Cvss, expectedIssue.Cvss)
	assert.Equal(t, issues[0].Confidence, expectedIssue.Confidence)
	assert.Equal(t, issues[0].Description, expectedIssue.Description)
}

var exampleOutput = `{
	"actions": [{
		"isMajor": true,
		"action": "install",
		"resolves": [{
			"id": 1179,
			"path": "auditjs>npm>bin-links>cmd-shim>mkdirp>bar",
			"dev": false,
			"optional": false,
			"bundled": true
		}]

	}],
	"advisories": {
		"1084": {
			"findings": [{
				"version": "1.1.0",
				"paths": [
					"auditjs>npm>libnpx>yargs>os-locale>foo"
				]
			}],
			"id": 1084,
			"created": "2019-07-18T21:30:31.935Z",
			"updated": "2019-11-19T23:31:37.349Z",
			"deleted": null,
			"title": "Denial of Service",
			"found_by": {
				"link": "",
				"name": "asdfasdfa",
				"email": ""
			},
			"reported_by": {
				"link": "",
				"name": "asdfasfa",
				"email": ""
			},
			"module_name": "foo",
			"cves": [],
			"vulnerable_versions": "<1.2.3",
			"patched_versions": ">=1.2.3",
			"overview": "Versions of 'foo' prior to 1.2.3 are vulnerable to Denial of Service (DoS).  The package fails to remove old values from the cache even after a value passes its 'maxAge' property. This may allow attackers to exhaust the system's memory if they are able to abuse the application logging.",
			"recommendation": "Upgrade to version 1.2.3 or later.",
			"references": "- [Snyk Report](https://snyk.io/vuln/npm:foo:20180117)",
			"access": "public",
			"severity": "low",
			"cwe": "CWE-400",
			"metadata": {
				"module_type": "",
				"exploitability": 2,
				"affected_components": ""
			},
			"url": "https://npmjs.com/advisories/123"
		}
	},
	"muted": [],
	"metadata": {
		"vulnerabilities": {
			"info": 0,
			"low": 136,
			"moderate": 0,
			"high": 23,
			"critical": 0
		},
		"dependencies": 748,
		"devDependencies": 0,
		"optionalDependencies": 5,
		"totalDependencies": 753
	},
	"runId": "75a2f929-191f-4b32-a0a1-28bfd230a36d"
}
`
