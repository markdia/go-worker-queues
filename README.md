go-worker-queues
================

Putting together a pattern based worker queue system as a way to learn GoLang.  I have some AWS needs, so I'm going to do some intergration there with MitchellH AMZGo fork.

At a point in time where this stops being a pet project and an educational bit, I'll post more accurate notes.

The system is relatively straight forward:
1. There is a collector which exposes an http/https endpoint for AWS SNS to talk to and message.  This is buffered so its not blocking.  This is basically how you ask for work to be done.
2. There is a dispatcher which takes jobs and hands them off to configurable number worker queues
3. There are worker queues which are how 'work' gets done

Nothing too sexy or inventive, but it should certainly be fast and effective.

Questions that I want to answer
1. best practices for logging in GoLang - using the "log" package vs printlns
2. best practices around testing in GoLang - building out tests to run on build/install time
3. speed of using bufferend channels
4. Passing in parameters like worker number and S3 bucket number - maybe AWS credentials at runtime to override defaults
5. Configuration files and management of those
6. Using the HTTP package... parsing URL params, requests, responses

