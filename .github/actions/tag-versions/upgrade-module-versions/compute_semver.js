const semver = require('semver')

const JSON_INPUT = process.argv[2]
const versionObj = JSON.parse(JSON_INPUT)

for (let [mod, version] of Object.entries(versionObj)) {
    const currVer = semver.valid(version) || "0.0.0"
    const newVer = semver.inc(currVer, 'patch')
    versionObj[mod] = `v${newVer}`
}

console.log(JSON.stringify(versionObj))
