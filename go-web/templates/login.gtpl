<html>
    <head>
    <title></title>
    </head>
    <body>
    <form action="/login" method="post">
        username:<input type="text" name="username">
        passowrd:<input type="password" name="password">
        <input type="hidden" name="token" value="{{.}}">
        <input type="submit" value="login">
    </form>
    </body>
</html>
