# Example App Go

![Profile](https://img.shields.io/badge/profile-go--service-00add8)
![Lifecycle](https://img.shields.io/badge/lifecycle-pilot-f59e0b)
[![Platform Workflow](https://img.shields.io/badge/platform--workflow-v0.4.0-0969da)](https://github.com/OXY7O/platform-workflow/releases/tag/v0.4.0)
![Canonical](https://img.shields.io/badge/Go-1.26.7-00add8)
![Tanpa deployment](https://img.shields.io/badge/deployment-tidak%20tersedia-6b7280)

Contoh teknis minimal untuk mengadopsi profile `go-service`. Aplikasi ini
menggunakan standard library Go, reusable workflow immutable, satu canonical
binary artifact, dan compatibility lane tanpa artifact.

Repository ini adalah materi pembelajaran dan pilot, bukan production starter
atau klaim operational compliance.

## Hubungan repository

```text
platform-governance@v1.2.0
  policy, lifecycle, control, evidence requirement
        |
platform-workflow@v0.4.0
  reusable Go Service CI dan binary artifact contract
        |
example-app-go
  thin caller dan implementasi consumer yang dapat diuji
```

- [Platform Governance](https://github.com/OXY7O/platform-governance)
- [Go Service profile](https://github.com/OXY7O/platform-workflow/blob/v0.4.0/docs/profiles/go-service/README.md)
- [Traceability](docs/TRACEABILITY.md)

## Quick start

Prasyarat: Go 1.26.7 atau Go 1.27.0.

```bash
go mod verify
go test -race ./...
go run ./cmd/api
```

Service mendengarkan `:8080` secara default. Untuk pengujian lokal, buka
`http://localhost:8080/health`.

## Endpoint

| Method | Path | Status | Response |
|---|---|---|---|
| GET | `/health` | 200 | `{"status":"ok"}` |
| GET | `/examples/42` | 200 | `{"id":42,"name":"example-42"}` |
| GET | `/examples/not-a-number` | 400 | controlled JSON error |
| selain GET | endpoint di atas | 405 | controlled JSON error |

Output menggunakan struct dan JSON encoder agar field serta nilainya
deterministik dan mudah diuji.

## Version matrix

| Lane | Runtime | Tujuan | Blocking | Artifact |
|---|---|---|---|---|
| Canonical | Go 1.26.7 | CI utama dan packaging | Ya | Satu binary package |
| Compatibility | Go 1.27.0 | Validasi runtime berikutnya | Ya | Tidak |

Compatibility tetap `not-validated` sampai evidence pilot aktual direview.

## Pemeriksaan CI

Reusable workflow menjalankan module verification, frozen graph check, format,
vet, unit test, race test, coverage minimum 80%, `govulncheck`, main package
validation, dan deterministic Linux amd64 build.

Caller hanya memiliki `contents: read`. Tidak ada secret, OIDC, environment,
arbitrary command, atau deployment.

## Artifact

Canonical lane menghasilkan tepat satu arsip
`example-api_0.1.0_linux_amd64.tar.gz`. Arsip memuat:

- executable `example-api` untuk Linux amd64;
- manifest JSON dengan source/workflow SHA, contract/module digest, binary
  digest, archive digest, runtime, dan target.

Status `ci-qualified` berarti artifact lolos CI. Status tersebut bukan izin
promotion, release production, atau deployment. Compatibility lane tidak
mengunggah artifact.

## Cara mengadopsi

1. Pastikan aplikasi sesuai profile Go Service/API.
2. Salin pola thin caller dari `.github/workflows/ci.yml`.
3. Ganti module path, binary name, main package, version, dan threshold dalam
   batas kontrak.
4. Pin workflow ke full commit SHA yang sudah disetujui.
5. Commit `go.mod`, `go.sum`, source, test, dan compatibility catalogue.
6. Buka pull request dan tunggu canonical serta compatibility checks.
7. Verifikasi hanya canonical lane menghasilkan artifact.
8. Simpan Safe evidence metadata sesuai klasifikasi storage organisasi.

## Onboarding checklist

- [ ] Owner teknis dan work item onboarding sudah ditetapkan.
- [ ] Runtime dan target sesuai Version matrix.
- [ ] `go.mod` serta `go.sum` konsisten.
- [ ] Main package dapat dibangun dengan `CGO_ENABLED=0`.
- [ ] Format, vet, unit, race, coverage, dan vulnerability check lulus.
- [ ] Caller hanya read-only dan dipin ke immutable SHA.
- [ ] Tidak ada credential atau nilai environment dalam source/evidence.
- [ ] Canonical artifact dan manifest digest sudah diperiksa.
- [ ] Compatibility lane terbukti tidak menghasilkan artifact.
- [ ] Deployment ditangani oleh kontrak dan approval terpisah.

## Troubleshooting

1. Buka check yang gagal dan baca failure category.
2. Pastikan toolchain tepat, lalu jalankan `go mod verify`.
3. Jalankan `gofmt`, `go vet`, unit/race test, coverage, dan `govulncheck`.
4. Periksa `./cmd/api`, target Linux amd64, dan kebutuhan CGO.
5. Jangan melemahkan required check atau compatibility lane.
6. Untuk platform failure, kirim run URL dan Safe evidence metadata tanpa
   menyalin data sensitif.

## Tanpa deployment

Repository ini berhenti pada CI dan artifact handoff. Ia tidak memiliki
credential, environment approval, promotion, release production, rollback, atau
workflow deployment. Implementasi deployment akan mengikuti kontrak terpisah
dari Platform Governance dan Platform Workflow.
