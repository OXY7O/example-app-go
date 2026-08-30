import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";

const read = (path) => fs.readFileSync(path, "utf8");

test("repository policy declares a pilot without deployment", () => {
  const policy = JSON.parse(read("repository-policy.json"));
  assert.equal(policy.profileKey, "go-service");
  assert.equal(policy.governanceBaseline, "platform-governance@v1.2.0");
  assert.equal(policy.workflowRelease, "platform-workflow@v0.4.0");
  assert.equal(policy.lifecycle, "pilot");
  assert.equal(policy.compatibility, "not-validated");
  assert.equal(policy.deploymentEnabled, false);
  assert.equal(policy.secretsAllowed, false);
});

test("thin caller uses the released immutable workflow without secrets", () => {
  const workflow = read(".github/workflows/ci.yml");
  assert.match(workflow, /ci-profile-go-service\.yml@010dee6edcdf21f813e842ef3f193f4a2c83593e/);
  assert.match(workflow, /permissions:\s*\n\s*contents: read/);
  assert.match(workflow, /\"goVersion\":\"1\.26\.7\"/);
  assert.doesNotMatch(workflow, /secrets:|id-token:|environment:|deploy/i);
});

test("compatibility catalogue contains only the governed artifactless lane", () => {
  const catalogue = JSON.parse(read("compatibility/go-service.json"));
  assert.equal(catalogue.profileKey, "go-service");
  assert.equal(catalogue.lanes.length, 1);
  assert.deepEqual(catalogue.lanes[0], {
    laneId: "go-1.27.0-linux-amd64",
    goVersion: "1.27.0",
    targetOs: "linux",
    targetArch: "amd64",
    workingDirectory: ".",
    executionMode: "compatibility-only",
    lifecycle: "active",
    blocking: true,
    eligible: true,
    artifact: false,
  });
});

test("README supports a developer onboarding journey", () => {
  const readme = read("README.md");
  for (const heading of [
    "Hubungan repository", "Quick start", "Endpoint", "Version matrix",
    "Artifact", "Cara mengadopsi", "Onboarding checklist", "Troubleshooting",
    "Tanpa deployment",
  ]) {
    assert.match(readme, new RegExp(heading));
  }
  assert.ok((readme.split("\n").slice(0, 12).join("\n").match(/!\[/g) ?? []).length <= 5);
});

test("safe evidence example contains no sensitive field", () => {
  const evidence = JSON.parse(read("docs/evidence/example-safe-evidence.json"));
  assert.equal(evidence.schemaVersion, "1.0");
  assert.equal(evidence.profileKey, "go-service");
  assert.equal(evidence.actual, false);
  assert.equal(evidence.deploymentAuthorized, false);
  assert.doesNotMatch(JSON.stringify(evidence), /secret|token|password|private.?key/i);
});
