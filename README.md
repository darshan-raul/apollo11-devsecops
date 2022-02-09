# devsecops-primer

- WAF

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
