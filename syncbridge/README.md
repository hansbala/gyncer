# Sync Bridge

This service lives on the server. It is responsible for running periodically scheduled cron jobs.

## Architecture Overview

The idea behind this service is it is always running and will execute sync jobs every `x` minutes - this is configurable via a config. By default, this would be something like 10 mins for me, personally.

Every `x` minutes it will poll SQL to see if any syncs need to be run. If so, it will run them and update the database saying "hey this job las ran at this time. The next time this sync will run will be `x1`.

Then, at time `now + x` it will poll again. Let's say `x1 < (now + x)` it will run sync job and update database again. The cycle carries on for all jobs.