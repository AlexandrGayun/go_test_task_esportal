<h1> How to run </h1>
<ol>
<li> provide credentials to .env</li>
<li> <code>$ docker-compose build</code> </li>
<li> <code>$ docker-compose up</code> </li>
<li> use commands from Makefile for migrate/seed/test/benchmark</li>
</ol>
<h1> Makefile </h1>
<ul>
<li>make migrateup/migratedown - run migrations up/down in the API container</li>
<li>make migrateuptest/migratedowntest - do the same but for test container</li>
<li>make populateseeds/populateseedstest - populate seeds for API/test container</li>
<li>make test - run all tests. You need to previously run migrations and populate seeds
or tests will fail</li>
<li>make benchmark - run benchmarks for Util package</li>
<li>make autotest - this will run migration up, population seeds, tests, migration down consequently for test container</li>
</ul>