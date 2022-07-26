{{define "base"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.0/css/bootstrap.min.css">
    <meta charset="UTF-8">
    <title>{{template "title" .}}</title>
</head>
<body>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/socket.io/2.3.0/socket.io.js"></script>
    <script src="https://code.jquery.com/jquery-3.4.1.min.js"></script>
    <script src="/static/js/main.js"></script>

    <div class="container">
        {{template "main" .}}
    </div>
</body>
</html>
{{end}}