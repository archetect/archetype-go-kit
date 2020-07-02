<h2>Build:</h2>

cd {{artifact_id}}-service

<h4>normal build:</h4>
go build .
<br \>
go run main.go


<h4>docker build:</h4>
docker build -t {{artifact_id}}-service .
<br \>
docker run -p {{default_port}}:{{default_port}} -it {{artifact_id}}-service

<h2>Endpoints:</h2>
<h5>status: http://localhost:{{default_port}}/status

metrics: http://localhost:{{default_port}}/metrics</h5>