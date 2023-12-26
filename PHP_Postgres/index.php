<!DOCTYPE html>
<html lang="en">
<head>
          <meta charset="UTF-8">
          <meta name="viewport" content="width=device-width, initial-scale=1.0">
          <title>Document</title>
</head>
<body>
          <h2>
                    employees page
          </h2>
          <?php
          $connection = pg_connect("host=localhost port=5433 dbname=CompanyDB user=postgres password=mflv[DB2023");
          $query = "SELECT * FROM public.employee";
          $result = pg_query($connection, $query);
          ?>

          <table>
                    <thead>
                              <tr>
                                        <th>employeeid</th>
                                        <th>firstname</th>
                                        <th>lastname</th>
                              </tr>
                    </thead>
                    <tbody>
                              <?php
                              while ($row = pg_fetch_row($result)) {
                                        echo "<tr>";
                                        echo "<td>" . $row[0] . "</td>";
                                        echo "<td>" . $row[1] . "</td>";
                                        echo "<td>" . $row[2] . "</td>"; 
                                        echo "</tr>";
                              }
                              ?>
                    </tbody>
          </table>

</body>
</html>