# Apollo 11 DevSecOps 🚀

Welcome to the **Apollo 11 DevSecOps** repository. This project serves as a comprehensive guide and reference implementation for building secure CI/CD pipelines using GitHub Actions. It demonstrates how to integrate security controls at every stage of the software development lifecycle (SDLC), embodying the "Shift Left" philosophy to ensure security is baked in from the start, not bolted on at the end.

## 🛡️ What is DevSecOps?

DevSecOps is the evolution of DevOps, integrating security practices directly into the DevOps process.

*   **Security by Default**: Security is an integral part of the development process, not a separate phase.
*   **Shift Left**: Moving security checks to the earliest possible stages of development (planning and coding) to find and fix vulnerabilities cheaper and faster.
*   **Automation**: Leveraging tools to automate security scans, policy enforcement, and compliance checks within the CI/CD pipeline.

## 🔄 DevSecOps Stages & Workflows

This repository hosts examples of GitHub Action workflows for every stage of the DevSecOps lifecycle.

### 1. 📝 Plan (Threat Modeling)
Understanding potential threats before writing a single line of code.
*   **Activities**: Design review, Threat Modeling.
*   **Tools**: IriusRisk, Threat Dragon.
*   **Workflows**: Automating threat model updates and tracking security requirements.

### 2. 💻 Code (Secure Coding)
Ensuring the code written is secure and free of secrets before it enters the repository.
*   **Pre-commit Hooks**: Prevent bad code from being committed.
    *   *Tool*: `talisman`, `pre-commit`
*   **Secret Scanning**: Detect hardcoded secrets (API keys, passwords, tokens).
    *   *Tool*: **Gitleaks**, **Talisman**
*   **IDE Plugins**: Real-time security feedback for developers.

### 3. 🏗️ Build (SCA & SAST)
Analyzing the application code and dependencies during the build process.
*   **Software Composition Analysis (SCA)**: identifying vulnerabilities in third-party libraries and dependencies.
    *   *Tool*: **OWASP Dependency-Check**, **Snyk**, **GitHub Dependabot**
*   **Static Application Security Testing (SAST)**: Analyzing source code for security flaws (e.g., SQL injection, XSS) without executing it.
    *   *Tool*: **CodeQL**, **Semgrep**, **SonarQube**, **Veracode**

### 4. 📦 Test (DAST & Container Security)
Testing the running application and the artifacts produced.
*   **Container Security**: Scanning Docker images for OS-level vulnerabilities and misconfigurations.
    *   *Tool*: **Trivy**, **Grype**, **Clair**
*   **Dynamic Application Security Testing (DAST)**: Attacking the running application to find vulnerabilities executable in a live environment.
    *   *Tool*: **OWASP ZAP**, **Burp Suite Enterprise**

### 5. 🚀 Deploy (IaC Security)
Ensuring the infrastructure hosting the application is secure.
*   **Infrastructure as Code (IaC) Scanning**: specialized SAST for Terraform, Kubernetes, CloudFormation, etc.
    *   *Tool*: **Checkov**, **Terrascan**, **TFLint**, **Kics**
*   **Compliance as Code**: Validating infrastructure against policy requirements.
    *   *Tool*: **OPA (Open Policy Agent)**, **Chef InSpec**

### 6. 👁️ Monitor (Runtime Security)
Continuous monitoring of the application and infrastructure in production.
*   **Runtime Application Self-Protection (RASP)**: Protecting the app at runtime.
*   **Vulnerability Management**: Agregating and tracking vulnerabilities.
    *   *Tool*: **DefectDojo**, **ArcherySec**
*   **Security Information and Event Management (SIEM)**: Log analysis and alerting.
    *   *Tool*: **ELK Stack**, **Splunk**

---

## ✅ Implementation Roadmap

Implementation of DevSecOps workflows covered in this repository.

### Phase 1: Foundation & Code Security
- [ ] **Secret Scanning Workflow**: Integrate `gitleaks` to scan for secrets in PRs and push events.
- [ ] **Linting & formatting**: Add standard linters (e.g., `golangci-lint`, `pylint`) to ensure code quality.
- [ ] **Pre-commit Setup**: Documentation and config for local `pre-commit` hooks (guides for users).

### Phase 2: Build & Dependency Security
- [ ] **SCA Workflow**: Implement **OWASP Dependency-Check** or **Trivy** (fs mode) for dependency scanning.
- [ ] **SAST Workflow**: Set up **CodeQL** (GitHub Advanced Security) or **Semgrep** (OSS) for static analysis.
- [ ] **Docker Build & Scan**: Create a workflow to build Docker images and scan them with **Trivy** or **Grype**.

### Phase 3: Infrastructure & Deployment Security
- [ ] **IaC Scanning**: Integrate **Checkov** or **Terrascan** to scan Terraform/Kubernetes files.
- [ ] **Kubernetes Manifest Validation**: Use `kube-linter` or verify manifests against strict policies.

### Phase 4: Dynamic Testing
- [ ] **DAST Workflow**: specialized workflow to deploy a test instance and run **OWASP ZAP** baseline scan.

### Phase 5: Reporting & Unified Dashboard
- [ ] **Security Report Artifacts**: Ensure all scanners upload SARIF reports to GitHub Security tab (or as artifacts).
- [ ] **DefectDojo Integration (Optional)**: Workflow to push findings to a central DefectDojo instance.





