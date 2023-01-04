# Build process

- Execute the build.sh script
- Then run: docker build -t test .
- Then run: docker run -di -p 3000:3000 -p 5000:5000 --env-file api/.env --name test1 test
