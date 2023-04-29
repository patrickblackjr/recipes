---
title: {{.Title}}
description: {{.Description}}
date: {{.Date}}
draft: false
time: {{.TotalTime}}
tags: []
featured_image: {{.ImageURL}}
---

## Ingredients

{{range .Ingredients}}- {{.}}
{{end}}
## Instructions

{{range .Instructions}}1. {{.}}
{{end}}