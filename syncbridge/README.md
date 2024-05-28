# Sync Bridge

This service lives on the server. It is responsible for running periodically scheduled cron jobs.

## Architecture Overview

The idea behind this service is it is always running and will execute sync jobs every `x` minutes - this is configurable via a config. By default, this would be something like 10 mins for me, personally.

Every `x` minutes it will poll SQL to see if any syncs need to be run. If so, it will run them and update the database saying "hey this job las ran at this time. The next time this sync will run will be `x1`.

Then, at time `now + x` it will poll again. Let's say `x1 < (now + x)` it will run sync job and update database again. The cycle carries on for all jobs.

## TODO

The more I think about this, the more I feel like this doesn't need to be it's own service. Since the service is so simple, it should live within the API but the periodic fetching should be on it's own thread in the API service. We'll just have to make sure to allocate at least 2 cores to the machine running the API which in today's day and age should be pretty okay.

In the future, if the API load is way too high, then the entire syncing process should live in it's own microservice. I don't expect too many users so YOLO :))