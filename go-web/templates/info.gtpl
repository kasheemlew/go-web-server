<html>
  <body>
  <div>
    <form action="/info" method="post">

        <h3>Choose Fruit</h3></br>
            <select name="fruit">
                <option value="apple">apple</option>
                <option value="pear">pear</option>
                <option value="banane">banane</option>
            </select></br>

        <h3>Choose Gender</h3></br>
            <input type="radio" name="gender" value="1">Male</input>
            <input type="radio" name="gender" value="2">Female</input>
            </br>

        <h3>Choose Interest</h3></br>
            <input type="checkbox" name="interest" value="football">Football
            <input type="checkbox" name="interest" value="basketball">Basketball
            <input type="checkbox" name="interest" value="tennis">Tennis
        </br><input type="submit" value="submit">
    </form>
  </div>
  </body>
</html>
