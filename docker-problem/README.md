# Docker Debugging Challenge

Let's find out how long it takes you to solve a real-world, production issue with Docker.

This challenge is based on a real-world bug that caused intermittent 502 errors in production. It drove developers crazy for weeks. Sometimes it worked. Sometimes it didn't. The cause? Subtle. The fix? Simple once you see it.

## The Rules

- Take a look at the files and folders in the example.
- If you can, don't peek inside the `simulated_error.sh` file. Just glance over it to confirm it does not do anything suspicious on your system.
- First, use only Docker commands and Google. No AI. Just your brain and plain searches.
- Then solve the same problem using any AI assistant and compare the results and times.

Log your time for both approaches and see which one wins.

## Step-by-Step Instructions

1. Download the repository.
2. Navigate to the `docker-problem` folder inside the project.
3. Run `docker compose up`.
4. Open a browser and confirm that both `http://localhost/app1` and `http://localhost/app2` are working correctly.
5. Run the script `simulated_error.sh`.
6. Confirm that `/app1` no longer works.
7. Make the necessary changes in the `proxy/nginx.conf` file to solve the error.
8. Ensure that, after your changes, no matter how many times you run `simulated_error.sh`, both endpoints continue to work perfectly.

## Purpose

This is not just about fixing a bug. It is about understanding Docker, container networking and the limitations of tools like Nginx.

Download the materials, track your time and share your results in the comments.
