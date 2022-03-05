# devsecops-primer


Security in Devops can be categorised as:
- Source code review
- dependency checking
- compliance review
- automated security scans
- penetration testing
- security incidents and response

- Shift the security left ie. Start it at a early stage


Devsecops : Effort to be "secure by default"
- integrate security via tools
- create security as code culture
- promote cross skilling

== Planning stage

- Threat modelling 

== Developer stage

- pre-commit hooks - talisman

== Source code stage

- secrets management - vault

== CI stage

- dependency checks - dependency checker
- SAST (static analysis security testing ) (source code review) - semgrep/veracode/checkmark
- find sec bugs

== Before deployin

- Container scanning - Trivy

== After code is deployed

- DAST (Dynamic analysis security testing) - OWASP zap 
- Vulnerability assessment tool - OpenVAS 
- Compliance as code - chef Inspec
- RASP ( Runtime Application Security) 
    - Watches a company's application at runtime , analyzing its behaviour as well as context in which the behaviour occurs.
- CENTRAL vulnerabity management tool - archerysec

== Security events

- security incidents and response - EFK (elastalert)
- WAF - modsecurity



====================================

- Containers
    - Container image scanning


- unsecured s3 buckets

- insecure configuration
    - CLoud sec configuration - static/event scan - cloudcheckr
    - CLoud sec configuration - Audit scan:  - tenable
    - Sandbox escaping:
        - jumping from containers to host - aqua,twistlock,trend micro

Security for Devops technologies:
- Adapt existing security tools to new tech
- Address new security risks new tech introduced


Methodologies:

- automated app sec testing
    - Static
        - Static testing (SAST)
            - Scan code to find vulneratbilites
            - Scan takes days , builds take minutes
            - Adaption: incremental scans, long scans weekly, delta scans during builds

    - dynamic
        - Dynamoic testing (DAST)
            - tests deployed instance like hacker
            - scans take too long
            - Adaption: IAST (instrument app, run unit-tests,deduce sec issues)
            
- SCA in CICD
    - flag use of libraries with known vulnerabilities

- Microservices:

    - track data flows accross apps - approetto
    - embedded install in deploy flow - signal sciences

- when a container misbehaves, you just kill it 


## References:
- https://aws.amazon.com/blogs/devops/building-end-to-end-aws-devsecops-ci-cd-pipeline-with-open-source-sca-sast-and-dast-tools/
- https://www.infoq.com/presentations/devsecops-2019/
- https://www.youtube.com/watch?v=-v3fs-i2O8w DevSecOps - Automating Security in DevOps | Jovin Lobo | Nullcon Webinar 2021
