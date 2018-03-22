Traffic Monitor
===============

Introduction
------------

Traffic Monitor is a Java Tomcat application that monitors caches,
provides health state information to Traffic Router, and collects
statistics for use in tools such as Traffic Ops and Traffic Stats. The
health state provided by Traffic Monitor is used by Traffic Router to
control which caches are available on the CDN.

Software Requirements
---------------------

To work on Traffic Monitor you need a \*nix (MacOS and Linux are most
commonly used) environment that has the following installed:

-   Eclipse &gt;= Kepler SR2 (or another Java IDE)
-   Maven &gt;= 3.3.1
-   JDK &gt;= 6.0

Traffic Monitor Project Tree Overview
-------------------------------------

-   `traffic_control/traffic_monitor/` - base directory for Traffic
    Monitor

    > -   `etc/` - Miscellaneous simulator utilities
    > -   `src/main` - Main source directory for the Traffic Monitor
    >
    >     > -   `bin/` - Configuration tools
    >     > -   `conf/` - Configuration files
    >     > -   `java/` - Java source code for Traffic Monitor
    >     > -   `opt/tomcat/conf` - Contains Tomcat configuration
    >     >     file(s) pulled in during an RPM build
    >     > -   `resources/` - Resources pulled in during an RPM build
    >     > -   `scripts/` - Scripts used by the RPM build process
    >     > -   `webapp/` - Java webapp resources
    >
    > -   `src/test` - Test source directory for Traffic Monitor
    >
    >     > -   `java/` - JUnit based unit tests for Traffic Monitor
    >     > -   `resources/conf` - Configuration files used by unit
    >     >     tests
    >     > -   `resources/db` - Files downloaded by unit tests
    >     > -   `resources/var` - Files generated by unit tests
    >
Java Formatting Conventions
---------------------------

None at this time. The codebase will eventually be formatted per Java
standards.

Installing The Developer Environment
------------------------------------

To install the Traffic Monitor Developer environment:

1.  Clone the traffic\_control repository using Git.
2.  Change directories into `traffic_control/traffic_monitor`.
3.  Edit the following parameters in
    src/test/resources/conf/traffic\_monitor\_config.js:

>   ------------------------------------------------------------------------
>   Parameter        Value
>   ---------------- -------------------------------------------------------
>   `tm.hostname`    FQDN of the Traffic Ops instance (do not include
>                    <http://>).
>
>   `tm.username`    Admin username for Traffic Ops
>
>   `tm.password`    Password for admin user
>
>   `cdnName`        Name of the CDN this Traffic Monitor will monitor
>   ------------------------------------------------------------------------
>
> Note: any change done later in the configuration file requires a mvn
> build in order to be applied.

4.  Import the existing git repo into Eclipse:

    > a.  File -&gt; Import -&gt; Git -&gt; Projects from Git; Next
    > b.  Existing local repository; Next
    > c.  Add -&gt; browse to find `traffic_control`; Add
    > d.  Select `traffic_control`; Next
    > e.  Ensure "Import existing projects" is selected, expand
    >     `traffic_control`, select `traffic_monitor`; Next
    > f.  Ensure `traffic_monitor` is checked; Finish
    > g.  Ensure `traffic_monitor` has been opened by Eclipse after
    >     importing

5.  Run `mvn clean verify` from the `traffic_monitor` directory
6.  Start the embedded Jetty instance from within Eclipse

    > a.  In the package explorer, expand `traffic_monitor`
    > b.  Expand `src/test/java`
    > c.  Expand the package
    >     `com.comcast.cdn.traffic_control.traffic_monitor`
    > d.  Open and run `Start.java`
    >
    >     > <div class="admonition note">
    >     >
    >     > If an error is displayed in the Console, run
    >     > `mvn clean verify` from the `traffic_monitor` directory
    >     >
    >     > </div>
    >
    > e.  With a web browser, navigate to <http://localhost:8080>

Test Cases
----------

Unit tests can be executed using Maven by running `mvn test` at the root
of the `traffic_monitor` project.

API
---

reference-tm-api

> hidden
>
> :   
>
> maxdepth
>
> :   1
>
> traffic\_monitor/traffic\_monitor\_api