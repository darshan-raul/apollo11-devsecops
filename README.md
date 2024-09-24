# Agenda

This project will be a guide and will serve examples for go and python code with regards to creating a devsecops pipelines

This repo is a sub project of a much larger [Apollo11](https://github.com/darshan-raul/Apollo11) project which covers not just the deployment aspect but overall covers all the components of creating a full fledged cloud native architecture leveraging all the possible tooling in the space.

> Note: The examples here are in github actions but the methodology mentioned can be implemented in any other CI tool

## What is DecSecOps?

DevSecOps acts like a cover or a shield for the existing DevOps practices. 

**Effort to be "secure by default"**

- ‘Shift left’ is a DevSecOps mantra ie. Start it at a early stage
- It encourages software engineers to move security from the right (end) to the left (beginning) of the DevOps (delivery) process. 
- In a DevSecOps environment, security is an integral part of the development process from the beginning. 
- An organization that uses DevSecOps brings in its cybersecurity architects and engineers as part of the development team. 
- Their job is to ensure every component, and every configuration item in the stack is patched, configured securely, and documented. 


## How do we achieve it? 

- integrate security via tools
- create security as code culture
- promote cross skilling

Security in Devops can be categorised as:
- Source code review
- dependency checking
- compliance review
- automated security scans
- penetration testing
- security incidents and response

### Planning stage

- Threat modelling 

### Developer stage

- pre-commit hooks - talisman

### Source code stage

- secrets management - vault

### CI stage

- dependency checks - dependency checker
- SAST (static analysis security testing ) (source code review) - semgrep/veracode/checkmark
- find sec bugs

### Before deploying

- Container scanning - Trivy

### After code is deployed

- DAST (Dynamic analysis security testing) - OWASP zap 
- Vulnerability assessment tool - OpenVAS 
- Compliance as code - chef Inspec
- RASP ( Runtime Application Security) 
    - Watches a company's application at runtime , analyzing its behaviour as well as context in which the behaviour occurs.
- CENTRAL vulnerabity management tool - archerysec

== Security events

- security incidents and response - EFK (elastalert)
- WAF - modsecurity





