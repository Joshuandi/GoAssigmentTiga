package web

const Web = `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="refresh" content="15">
    <title>web</title>
</head>
<body>
    <div>
        <ul>
            Status: {{value.Status}}
            <li>
                water : {{$value.Water}}
                wind : {{$value.Wind}}
            </li>
        </ul>
    </div>
</body>
</html>
`
