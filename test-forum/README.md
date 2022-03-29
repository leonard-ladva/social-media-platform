Forum task

In order to build the docker image and run server, use test file: bash dockerbuild.sh

Forum will be available via: https://localhost:8000/
Dummy posts, users and comments will be automatically generated.

User - password

admin - admin

moderator - moderator

test1 - test1

test2 - test2

test3 - test3

## AUTHENTICATION

Audit questions: https://github.com/01-edu/public/blob/master/subjects/forum/authentication/audit.md

Test live @ https://onlycans.ee

OR

Register an OAuth app at Facebook, Github and Google. Save the client ID-s and secrets and export them to environment variables before running the server:

export GCLIENT_ID="YOUR GOOGLE CLIENT ID"

export GOOGLE_SECRET="YOUR GOOGLE SECRET"

export FBCLIENT_ID="YOUR FACEBOOK APP ID"

export FB_SECRET="YOUR FACEBOOK SECRET"

export GHCLIENT_ID="YOUR GITHUB APP ID"

export GH_SECRET="YOUR GITHUB SECRET"

Instructions for creating an OAuth APP:

Github - https://docs.github.com/en/developers/apps/building-oauth-apps/creating-an-oauth-app

Google - https://developers.google.com/identity/sign-in/web/sign-in

Facebook - https://developers.facebook.com/docs/development/create-an-app/

Remember to add the OAuth callback uri (localhost:8000/connect) in all of the above OAuth providers



Authors: georgi.suikanen, Rostislav, specest, Tk
