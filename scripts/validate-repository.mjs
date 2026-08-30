import assert from "node:assert/strict";
import fs from "node:fs";

const required = [
  ".github/workflows/ci.yml",
  "README.md",
  "cmd/api/main.go",
  "compatibility/go-service.json",
  "docs/TRACEABILITY.md",
  "docs/evidence/example-safe-evidence.json",
  "go.mod",
  "go.sum",
  "internal/httpapi/handler.go",
  "repository-policy.json",
];

for (const path of required) {
  assert.ok(fs.existsSync(path), `required file missing: ${path}`);
}

const workflow = fs.readFileSync(".github/workflows/ci.yml", "utf8");
assert.match(workflow, /@[a-f0-9]{40}/, "workflow must use an immutable pin");
assert.doesNotMatch(workflow, /secrets:|id-token:|environment:|deploy/i);

const policy = JSON.parse(fs.readFileSync("repository-policy.json", "utf8"));
assert.equal(policy.deploymentEnabled, false);
assert.equal(policy.secretsAllowed, false);
assert.match(policy.workflowSha, /^[a-f0-9]{40}$/);

const catalogue = JSON.parse(fs.readFileSync("compatibility/go-service.json", "utf8"));
assert.equal(catalogue.lanes.length, 1);
assert.equal(catalogue.lanes[0].artifact, false);
assert.equal(catalogue.lanes[0].blocking, true);

console.log(`Validated ${required.length} required files and repository controls.`);
