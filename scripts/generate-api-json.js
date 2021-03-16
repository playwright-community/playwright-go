#!/usr/bin/env node
const fs = require("fs")
const path = require("path")
const { getAPIDocs } = require("./helpers")

const api = getAPIDocs()

fs.writeFileSync(path.join(__dirname, "..", "api.json"), JSON.stringify(api, null, 2))